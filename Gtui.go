package GTUI

import (
	"GTUI/Core"
	"GTUI/Core/Component"
	"GTUI/Core/Utils"
	"GTUI/Core/Utils/Color"
	"GTUI/Terminal"
	"strings"
	"time"
)

type Gtui struct {
	globalColor      Color.Color
	term             Terminal.ITerminal
	buff             []Core.IEntity
	componentManager *Component.ComponentM
	xCursor          int
	yCursor          int
}

func (c *Gtui) SetCur(x, y int) {
	compPreSet, _ := c.componentManager.Search(c.xCursor-1, c.yCursor-1)
	compPostSet, _ := c.componentManager.Search(x, y)
	inPreButNotInPost := Utils.Diff(compPostSet, compPreSet)
	inPostButNotInPre := Utils.Diff(compPreSet, compPostSet)
	for _, e := range inPreButNotInPost {
		e.OnLeave()
	}
	for _, e := range inPostButNotInPre {
		e.OnHover()
	}
	c.xCursor = x + 1
	c.yCursor = y + 1
}

func NewGtui() (*Gtui, error) {
	term := &Terminal.Terminal{}
	e := term.Start()
	if e != nil {
		return nil, e
	}
	xSize, ySize := term.Size()
	componentManager := Component.Create(xSize, ySize, 5)
	return &Gtui{term: term, buff: make([]Core.IEntity, 0), componentManager: componentManager}, nil
}

func (c *Gtui) Close() {
	c.IResetGlobalColor()
	c.term.Clear()
	c.term.Stop()
}

func (c *Gtui) InsertEntity(entity Core.IEntity) {
	c.buff = append(c.buff, entity)
}

func (c *Gtui) InsertComponent(component Component.IComponent) {
	c.buff = append(c.buff, component.GetGraphics())
	c.componentManager.Add(component)
}

func (c *Gtui) Click(x, y int) error {
	resultArray, e := c.componentManager.Search(x, y)
	if e != nil {
		return e
	}
	for i := range resultArray {
		resultArray[i].OnClick()//dare la possibilita' di scegliere il timer
		time.AfterFunc(time.Millisecond*1000, func() {
			resultArray[i].OnRelease()
			c.IRefreshAll()
		})
	}
	return nil
}

func (c *Gtui) IRefreshAll() {
	var str strings.Builder
	for _, b := range c.buff {
		str.WriteString(b.GetAnsiCode(c.globalColor))
		str.WriteString(c.globalColor.GetAnsiColor())
	}
	c.term.PrintStr(str.String())
	c.term.SetCursor(c.xCursor, c.yCursor)
}
