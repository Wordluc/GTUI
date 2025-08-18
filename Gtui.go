package GTUI

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Component"
	"github.com/Wordluc/GTUI/Core/EventManager"
	"github.com/Wordluc/GTUI/Core/Utils"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
	"github.com/Wordluc/GTUI/Keyboard"
	"github.com/Wordluc/GTUI/Terminal"
)

type Loop func(Keyboard.IKeyBoard, *Gtui) bool
type Gtui struct {
	globalColor          Color.Color
	term                 Terminal.ITerminal
	keyb                 Keyboard.IKeyBoard
	drawingTree          *Core.MatrixHandler
	componentTree        *Core.MatrixHandler
	xCursor              int
	yCursor              int
	xSize                int
	ySize                int
	cursorVisibility     bool
	loop                 Loop
	preComponentsHovered []Core.IComponent
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
		drawingTree:      Core.CreateMatrixHandler(10, 10),
		componentTree:    Core.CreateMatrixHandler(10, 10),
		xCursor:          0,
		yCursor:          0,
		xSize:            xSize,
		ySize:            ySize,
	}, nil
}

func (c *Gtui) initEventManager() {
	EventManager.Subscribe(EventManager.CursorAlign, 100, func(comp []any) {
		c.alignCursor()
	})
	EventManager.Subscribe(EventManager.Refresh, 100, func(comp []any) {
		c.SetCur(c.xCursor, c.yCursor)
		c.refresh(true)
	})
	EventManager.Subscribe(EventManager.ForceRefresh, 100, func(comp []any) {
		c.IClear()
		c.SetCur(c.xCursor, c.yCursor)
		c.refresh(false)
	})

	EventManager.Subscribe(EventManager.ReorganizeElements, 100, func(comp []any) {
		c.IClear()
		group := sync.WaitGroup{}
		group.Add(2)
		go func() {
			c.drawingTree.Refresh()
			group.Done()
		}()
		go func() {
			c.componentTree.Refresh()
			group.Done()
		}()
		c.SetCur(c.xCursor, c.yCursor)
		group.Wait()
		c.refresh(false)
	})
}
func (c *Gtui) getHigherLayerElementsNoDisabled(x, y int) (res []Core.IComponent) {
	compInLayers := c.componentTree.SearchInAllLayers(x, y)
	for i := len(compInLayers) - 1; i >= 0; i-- {
		if compInLayers[i] == nil {
			continue
		}
		var res []Core.IComponent
		for _, comp := range compInLayers[i] {
			comp := comp.(Core.IComponent)
			if comp.GetActive() {
				res = append(res, comp)
			}
		}
		if len(res) != 0 {
			return res
		}
	}
	return []Core.IComponent{}
}
func (c *Gtui) SetCur(x, y int) error {
	if x < 0 || y < 0 || x >= c.xSize || y >= c.ySize {
		return errors.New("cursor out of range")
	}
	if !c.cursorVisibility {
		return nil
	}
	compsPostSet := c.getHigherLayerElementsNoDisabled(x, y)

	inPreButNotInPost := Utils.GetDiff(compsPostSet, c.preComponentsHovered)
	inPostButNotInPre := Utils.GetDiff(c.preComponentsHovered, compsPostSet)

	for i, comp := range inPostButNotInPre {
		if ci, ok := inPostButNotInPre[i].(Core.IWritableComponent); ok {
			if ci.IsTyping() {
				continue
			}
		}
		comp.OnHover()
	}

	for i, comp := range inPreButNotInPost {
		if ci, ok := inPreButNotInPost[i].(Core.IWritableComponent); ok {
			if ci.IsTyping() {
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
			} else {
				comp.OnHover()
			}
			break
		}
	}

	c.yCursor = y
	c.xCursor = x
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
	c.preComponentsHovered = compsPostSet
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
	c.initEventManager()
	c.term.Start()
	c.term.Clear()
	c.lazyCheck()
	c.refresh(true)
	c.keyb.Start(c.innerLoop)
	c.term.Stop()
}
func (c *Gtui) lazyCheck() {
	time.AfterFunc(time.Second*2, func() {
		if c.term.Resized() {
			c.xSize, c.ySize = c.term.Size()
			EventManager.Call(EventManager.ForceRefresh, nil)
		}
		c.lazyCheck()
	})
}
func (c *Gtui) AddDrawing(entitiesToAdd ...Core.IDrawing) {
	for _, draw := range entitiesToAdd {
		c.drawingTree.AddElement(draw)
	}
}

func (c *Gtui) AddComponent(componentsToAdd ...Core.IComponent) error {
	for _, componentToAdd := range componentsToAdd {
		c.AddDrawing(componentToAdd.GetGraphics()...)
		c.componentTree.AddElement(componentToAdd)
		componentToAdd.OnLeave()
	}
	return nil
}

func (c *Gtui) AddContainer(container Core.IContainer) error {
	drawings := container.GetGraphics()
	componets := container.GetComponents()
	c.AddDrawing(drawings...)
	c.AddComponent(componets...)
	return nil
}

func (c *Gtui) AddComplexElement(complEle Core.IComplexElement) error {
	componets := complEle.GetComponents()
	drawings := complEle.GetGraphics()
	for _, comp := range componets {
		c.componentTree.AddElement(comp)
		comp.OnLeave()
	}
	for _, draw := range drawings {
		c.drawingTree.AddElement(draw)
	}
	return nil
}
func (c *Gtui) CallEventOn(x, y int, event func(Core.IComponent)) error {
	resultArray := c.getHigherLayerElementsNoDisabled(x, y)
	if len(resultArray) == 0 {
		return nil
	}
	for i := range resultArray {
		event(resultArray[i])
	}
	return nil
}

func (c *Gtui) Click(x, y int) error {
	resultArray := c.getHigherLayerElementsNoDisabled(x, y)

	if len(resultArray) == 0 {
		return nil
	}
	for i := range resultArray {
		resultArray[i].OnClick()
	}
	return nil
}

func (c *Gtui) alignCursor() {
	x, y := c.GetCur()
	oldX, oldY := x, y
	comps := c.getHigherLayerElementsNoDisabled(x, y)
	for _, comp := range comps {
		if ci, ok := comp.(Core.IWritableComponent); ok {
			if ci.IsTyping() {
				deltax, deltay := ci.DiffCurrentToXY(x, y)
				c.yCursor = y + deltay
				c.xCursor = x + deltax
			}
		}
	}

	x, y = c.GetCur()
	c.preComponentsHovered = comps
	comps = c.getHigherLayerElementsNoDisabled(x, y)

	inPreButNotInPost := Utils.GetDiff(comps, c.preComponentsHovered)
	for i, comp := range inPreButNotInPost {
		if ci, ok := inPreButNotInPost[i].(Core.IWritableComponent); ok {
			if ci.IsTyping() {
				ci.(*Component.TextBox).DeleteLastCharacter()
				c.yCursor = oldY
				c.xCursor = oldX
				c.term.SetCursor(c.xCursor+1, c.yCursor+1)
				return
			} else {
				comp.OnLeave()
			}
		} else {
			comp.OnLeave()
		}
	}

	c.preComponentsHovered = comps
	c.term.SetCursor(c.xCursor+1, c.yCursor+1)
}

func (c *Gtui) refresh(onlyTouched bool) error {
	var str strings.Builder
	var s strings.Builder
	var drew bool
	for i := 0; i < c.drawingTree.GetLayerN(); i++ {
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
	var drew bool = false
	cond := func(ele Core.ElementTree) bool {
		drawing = ele.(Core.IDrawing)
		if !drawing.GetVisibility() {
			return true
		}
		//	if onlyTouched && !drawing.IsTouched() {
		//		return true
		//	}
		x, y := drawing.GetPos()
		width, height := drawing.GetSize()
		str.WriteString(c.ClearZone(x, y, width, height))
		str.WriteString(drawing.GetAnsiCode(c.globalColor))
		str.WriteString(c.globalColor.GetAnsiColor())
		drew = true
		return true
	}
	c.drawingTree.ExecuteOnLayer(int(layer), cond)
	return str, drew
}

func (c *Gtui) innerLoop(keyb Keyboard.IKeyBoard) bool {
	c.alignCursor()
	if !c.loop(c.keyb, c) {
		return false
	}
	c.refresh(true)
	return true
}
