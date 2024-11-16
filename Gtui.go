package GTUI

import (
	"errors"
	"strings"
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Component"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/EventManager"
	"github.com/Wordluc/GTUI/Core/Utils"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
	"github.com/Wordluc/GTUI/Keyboard"
	"github.com/Wordluc/GTUI/Terminal"
)
type Loop func (Keyboard.IKeyBoard,*Gtui) bool
type Gtui struct {
	globalColor      Color.Color
	term             Terminal.ITerminal
	keyb             Keyboard.IKeyBoard
	entityTree *Core.TreeManager[Core.IEntity]
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
		entityTree: Core.CreateTreeManager[Core.IEntity](),
		xCursor:          0,
		yCursor:          0,
		xSize:            xSize,
		ySize:            ySize,
	}, nil
}

func (c *Gtui) initializeEventManager() {
	EventManager.Subscribe(EventManager.Refresh,100, func(comp []any) {
		c.refresh(true)
	})

	EventManager.Subscribe(EventManager.ReorganizeElements,100, func(comp []any) {
		c.entityTree.Refresh()
		c.IClear()
		c.refresh(false)
	})
}

func (c *Gtui) SetCur(x, y int) error {
	if x < 0 || y < 0 || x >= c.xSize || y >= c.ySize {
		return errors.New("cursor out of range")
	}
	if !c.cursorVisibility{
		c.xCursor = x
		c.yCursor = y
		return nil
	}
	compPreSet, _ := c.entityTree.Search(c.xCursor, c.yCursor)
	compsPostSet, _ := c.entityTree.Search(x, y)
	inPreButNotInPost := Utils.GetDiff(compsPostSet, compPreSet)
	inPostButNotInPre := Utils.GetDiff(compPreSet, compsPostSet)
	var comp Core.IComponent
	var ok bool
	for i, e := range inPostButNotInPre {
		if comp, ok = e.(Core.IComponent); !ok {
			continue
		}
		if ci, ok := inPostButNotInPre[i].(Core.IWritableComponent); ok {
			if ci.IsTyping() {
				continue
			}
		}
		comp.OnHover()
	}
	lenInPreButNotInPost := len(inPreButNotInPost)
	for i := 0; i < lenInPreButNotInPost; i++ {
		if comp, ok = inPreButNotInPost[i].(Core.IComponent); !ok {
			continue
		}
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

	c.yCursor = y
	c.xCursor = x

	for _, comp := range compsPostSet {
		if ci, ok := comp.(Core.IWritableComponent); ok {
			if ci.IsTyping() {
				c.xCursor, c.yCursor = ci.SetCurrentPosCursor(x, y)
			}
			break
		}
	}

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
	c.refresh(false)
	c.keyb.Start(c.innerLoop)
	c.term.Stop()
}

func (c *Gtui) InsertEntity(entityToAdd Core.IDrawing) {
	if container, ok := entityToAdd.(*Drawing.Container); ok {
		for _, entity := range container.GetChildren() {
			c.InsertEntity(entity)
		}
		return
	}
	c.entityTree.AddElement(entityToAdd)
}

func (c *Gtui) InsertComponent(componentToAdd Core.IComponent) error {
	if container, ok := componentToAdd.(*Component.Container); ok {
		for _, component := range container.GetComponents() {
			c.entityTree.AddElement(component)
			component.OnLeave()
		}
		c.InsertEntity(componentToAdd.GetGraphics())
	} else {
		componentToAdd.OnLeave()
		c.entityTree.AddElement(componentToAdd)
		c.InsertEntity(componentToAdd.GetGraphics())
	}
	return nil
}

func (c *Gtui) reorganizeTree() {
	c.entityTree.Refresh()
}

func (c *Gtui) CallEventOn(x, y int, event func(Core.IComponent)) error {
	resultArray, e := c.entityTree.Search(x, y)
	if e != nil {
		return e
	}
	for i := range resultArray {
		if comp,ok:=resultArray[i].(Core.IComponent);ok{
			event(comp)
		}
	}
	return nil
}

func (c *Gtui) Click(x, y int) error {
	resultArray, e := c.entityTree.Search(x, y)
	if e != nil {
		return e
	}
	for i := range resultArray {
		if comp,ok:=resultArray[i].(Core.IComponent);ok{
			comp.OnClick()
		}
	}
	return nil
}

func (c *Gtui) AllineCursor() {
	x, y := c.GetCur()
	comps, _ := c.entityTree.Search(x, y)
	for _, comp := range comps {
		if ci, ok := comp.(Core.IWritableComponent); ok {
			if ci.IsTyping() {
				deltax, deltay := ci.DiffCurrentToXY(x, y)
				c.yCursor = y + deltay
				c.xCursor = x + deltax
				break
			}
		}
	}
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
}

func (c *Gtui) refresh(onlyTouched bool) {
	var str strings.Builder
	var drawing Core.IDrawing
	var ok bool
	var elementToRefresh map[Core.IDrawing]struct{} = make(map[Core.IDrawing]struct{})
	cond := func(node *Core.TreeNode[Core.IEntity]) bool {
		var el Core.IDrawing
		if drawing, ok = node.GetElement().(Core.IDrawing); !ok {
			return true
		}
		if el = drawing; !el.IsTouched() && onlyTouched {
			return true
		}
		x, y := el.GetPos()
		width, height := el.GetSize()
		str.WriteString(c.ClearZone(x, y, width, height))
		str.WriteString(el.GetAnsiCode(c.globalColor))
		str.WriteString(c.globalColor.GetAnsiColor())
		for _, child := range c.entityTree.GetCollidingElement(node) {
			if drawing, ok = child.(Core.IDrawing); ok {
				elementToRefresh[drawing] = struct{}{}
			}
		}
		return true
	}
	c.entityTree.Execute(cond)
	for drawing := range elementToRefresh {
		str.WriteString(drawing.GetAnsiCode(c.globalColor))
		str.WriteString(c.globalColor.GetAnsiColor())
	}
	c.term.PrintStr(str.String())
	if c.cursorVisibility {
		str.WriteString(c.term.ShowCursor())
	}else{
		str.WriteString(c.term.HideCursor())
	}
	//DO NOT CHANGE THE ORDER
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
	return
}

func (c *Gtui) innerLoop(keyb Keyboard.IKeyBoard) bool {
	//Keyboard
	//	c.IRefreshAll()
	c.AllineCursor()
	if !c.loop(c.keyb,c) {
		return false
	}
	c.refresh(true)
	return true
}
