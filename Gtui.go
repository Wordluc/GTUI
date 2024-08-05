package GTUI

import (
	C "GTUI/Core"
	"GTUI/Core/Component"
	"GTUI/Core/Utils/Color"
	iTerminal "GTUI/Terminal"
	implTerm "GTUI/Terminal/impl"
	"strings"
)

type Gtui struct {
	globalColor      Color.Color
	term             iTerminal.ITerminal
	buff             []C.IEntity
	componentManager *Component.ComponentM
}

func NewGtui() (*Gtui, error) {
	term := &implTerm.Terminal{}
	e := term.Start()
	if e != nil {
		return nil, e
	}
	xSize, ySize := term.Size()
	componentManager := Component.Create(xSize, ySize, 5)
	return &Gtui{term: term, buff: make([]C.IEntity, 0), componentManager: componentManager}, nil
}

func (c *Gtui) Close() {
	c.IResetGlobalColor()
	c.term.Clear()
	c.term.Stop()
}

func (c *Gtui) InsertEntity(entity C.IEntity) {
	c.buff = append(c.buff, entity)
}

func (c *Gtui) InsertComponent(component Component.IComponent) {
	c.buff = append(c.buff, component.GetGraphics())
	c.componentManager.Add(component)
}

type InteractiveType int8
const (
   Click InteractiveType = iota
)
func (c *Gtui) Interact(x, y int, t InteractiveType) error {
	resultArray, e := c.componentManager.Search(x, y)
	if e != nil {
		return e
	}
	for i:= range resultArray {
		switch t {
		case Click:
			resultArray[i].OnClick()
		}
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
}

