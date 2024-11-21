package GTUI

import (
	"errors"
	"strings"
	"sync"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Component"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/EventManager"
	"github.com/Wordluc/GTUI/Core/Utils"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
	"github.com/Wordluc/GTUI/Keyboard"
	"github.com/Wordluc/GTUI/Terminal"
)

type Loop func(Keyboard.IKeyBoard, *Gtui) bool
type Gtui struct {
	globalColor      Color.Color
	term             Terminal.ITerminal
	keyb             Keyboard.IKeyBoard
	drawingTree      *Core.TreeManager[Core.IDrawing]
	componentTree    *Core.TreeManager[Core.IComponent]
	xCursor          int
	yCursor          int
	xSize            int
	ySize            int
	cursorVisibility bool
	loop             Loop
}

func NewGtui(loop Loop, keyb Keyboard.IKeyBoard, term Terminal.ITerminal) (*Gtui, error) {
	xSize, ySize := term.Size()
	EventManager.Setup()
	return &Gtui{
		globalColor:      Color.GetDefaultColor(),
		cursorVisibility: true,
		loop:             loop,
		term:             term,
		keyb:             keyb,
		drawingTree:      Core.CreateTreeManager[Core.IDrawing](),
		componentTree:    Core.CreateTreeManager[Core.IComponent](),
		xCursor:          0,
		yCursor:          0,
		xSize:            xSize,
		ySize:            ySize,
	}, nil
}

func (c *Gtui) initializeEventManager() {
	EventManager.Subscribe(EventManager.Refresh, 100, func(comp []any) {
		c.refresh(true)
	})

	EventManager.Subscribe(EventManager.ReorganizeElements, 100, func(comp []any) {
		c.IClear()
		group:=sync.WaitGroup{}
		group.Add(2)
		go func(){
			c.drawingTree.Refresh()
			group.Done()
		}()
		go func(){
			c.componentTree.Refresh()
			group.Done()
		}()
		group.Wait()
		c.refresh(false)
	})
}

//must be sorted
func getHighterElements(comps []Core.IComponent) []Core.IComponent {
	if len(comps) == 0 {
		return []Core.IComponent{}
	}
	var result []Core.IComponent=[]Core.IComponent{comps[0]}
	for _, comp := range comps[1:] {
		if comp.GetLayer() < comps[0].GetLayer() {
			break
		}
		result = append(result, comp)
	}
	return result
}

func (c *Gtui) SetCur(x, y int) error {
	if x < 0 || y < 0 || x >= c.xSize || y >= c.ySize {
		return errors.New("cursor out of range")
	}
	if !c.cursorVisibility {
		c.xCursor = x
		c.yCursor = y
		return nil
	}

	compPreSet := c.componentTree.GetHighterLayerElements(c.xCursor, c.yCursor)
	compsPostSet := c.componentTree.GetHighterLayerElements(x, y)

	inPreButNotInPost := Utils.GetDiff(compsPostSet, compPreSet)
	inPostButNotInPre := Utils.GetDiff(compPreSet, compsPostSet)

	for i, comp := range inPostButNotInPre {
		if ci, ok := inPostButNotInPre[i].(Core.IWritableComponent); ok {
			if ci.IsTyping() {
				continue
			}
		}
		comp.OnHover()
	}

	for i,comp := range inPreButNotInPost{
		if ci, ok := inPreButNotInPost[i].(Core.IWritableComponent); ok {
			if ci.IsTyping() {
				ci.SetCurrentPosCursor(x, y)
				return nil
			} else {
				comp.OnLeave()
			}
		} else {
			comp.OnLeave()
		}
	}

	for _, comp := range compsPostSet {
		if ci, ok := comp.(Core.IWritableComponent); ok {
			if ci.IsTyping() {
				c.xCursor, c.yCursor = ci.SetCurrentPosCursor(x, y)
				return nil
			}
			break
		}
	}

	c.yCursor = y
	c.xCursor = x
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
	return nil
}

func (c *Gtui) GetCur() (int, int) {
	return c.xCursor, c.yCursor
}

func (c *Gtui) CreateStreamingCharacter() Component.StreamCharacter {
	stream := Component.StreamCharacter{}
	stream.Get = func() chan string {
		i := c.keyb.NewChannel()
		stream.IChannel = i
		return c.keyb.GetChannels()[i]
	}
	stream.Delete = func() {
		c.keyb.DeleteChannel(stream.IChannel)
	}
	return stream
}

func (c *Gtui) Start() {
	c.initializeEventManager()
	c.term.Start()
	c.term.Clear()
	c.refresh(true)
	c.keyb.Start(c.innerLoop)
	c.term.Stop()
}

func (c *Gtui) AddDrawing(entityToAdd Core.IDrawing) {
	if container, ok := entityToAdd.(*Drawing.Container); ok {
		for _, entity := range container.GetChildren() {
			c.AddDrawing(entity)
		}
		return
	}
	c.drawingTree.AddElement(entityToAdd)
}

func (c *Gtui) AddComponent(componentToAdd Core.IComponent) error {
	if container, ok := componentToAdd.(*Component.Container); ok {
		for _, component := range container.GetComponents() {
			c.AddComponent(component)
		}
		c.AddDrawing(componentToAdd.GetGraphics())
		return nil
	}
	c.AddDrawing(componentToAdd.GetGraphics())
	c.componentTree.AddElement(componentToAdd)
	componentToAdd.OnLeave()
	return nil
}

func (c *Gtui) CallEventOn(x, y int, event func(Core.IComponent)) error {
	for i := c.componentTree.GetLayerN(); i >= 0; i-- {
		if isDone, e := c.callEventOnLayer(x, y, Core.Layer(i), event); isDone {
			return e
		}
	}
	return nil
}

func (c *Gtui) callEventOnLayer(x, y int, layer Core.Layer, event func(Core.IComponent)) (bool, error) {
	resultArray, e := c.componentTree.SearchOnLayer(layer, x, y)
	if e != nil {
		return false, e
	}
	if len(resultArray) == 0 {
		return false, nil
	}
	for i := range resultArray {
		event(resultArray[i])
	}
	return true, nil
}

func (c *Gtui) Click(x, y int) error {
	for i := c.componentTree.GetLayerN() ; i >= 0; i-- {
		if isDone, e := c.clickOnLayer(x, y, Core.Layer(i)); isDone {
			return e
		}
	}
	return nil
}
func (c *Gtui) clickOnLayer(x, y int, layer Core.Layer) (bool, error) {
	resultArray, e := c.componentTree.SearchOnLayer(layer, x, y)
	if e != nil {
		return false, e
	}
	if len(resultArray) == 0 {
		return false, nil
	}
	for i := range resultArray {
		resultArray[i].OnClick()
	}
	return true, nil
}

func (c *Gtui) AllineCursor() {
	for i := c.componentTree.GetLayerN(); i >= 0; i-- {
		if isDone := c._allineCursor(Core.Layer(i)); isDone {
			break
		}
	}
}
func (c *Gtui) _allineCursor(layer Core.Layer) bool {
	x, y := c.GetCur()
	comps, _ := c.componentTree.SearchOnLayer(layer, x, y)
	for _, comp := range comps {
		if ci, ok := comp.(Core.IWritableComponent); ok {
			if ci.IsTyping() {
				deltax, deltay := ci.DiffCurrentToXY(x, y)
				c.yCursor = y + deltay
				c.xCursor = x + deltax
				return true
			}
		}
	}
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
	return false
}

func (c *Gtui) refresh(onlyTouched bool) error { //TODO: optimize
	var str strings.Builder
	var s strings.Builder
	var drew bool
	for i := 0; i < c.componentTree.GetLayerN(); i++ {
		s, drew = c.refreshLayer(Core.Layer(i), onlyTouched)
		if drew {
			onlyTouched = false
		}
		str.WriteString(s.String())
	}
	if c.cursorVisibility {
		str.WriteString(c.term.ShowCursor())
	} else {
		str.WriteString(c.term.HideCursor())
	}
	//DO NOT CHANGE THE ORDER
	c.term.PrintStr(str.String())
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
	return nil
}

func (c *Gtui) refreshLayer(layer Core.Layer, onlyTouched bool) (strings.Builder, bool) {
	var str strings.Builder
	var drawing Core.IDrawing
	var drew bool=false
	var elementToRefresh map[Core.IDrawing]struct{} = make(map[Core.IDrawing]struct{})
	cond := func(node *Core.TreeNode[Core.IDrawing]) bool {
		drawing = node.GetElement()
		if onlyTouched && !drawing.IsTouched() {
			return true
		}
		x, y := drawing.GetPos()
		width, height := drawing.GetSize()
		str.WriteString(c.ClearZone(x, y, width, height))
		str.WriteString(drawing.GetAnsiCode(c.globalColor))
		str.WriteString(c.globalColor.GetAnsiColor())
		drew = true
		for _, child := range c.drawingTree.GetCollidingElement(layer, node) {
			elementToRefresh[child] = struct{}{}
		}
		return true
	}
	c.drawingTree.ExecuteOnLayerForAll(layer, cond)
	for drawing := range elementToRefresh {
		str.WriteString(drawing.GetAnsiCode(c.globalColor))
		str.WriteString(c.globalColor.GetAnsiColor())
	}
	return str, drew
}

func (c *Gtui) innerLoop(keyb Keyboard.IKeyBoard) bool {
	//Keyboard
	//	c.IRefreshAll()
	c.AllineCursor()
	if !c.loop(c.keyb, c) {
		return false
	}
	c.refresh(true)
	return true
}
