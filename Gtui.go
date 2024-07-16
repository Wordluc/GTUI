package GTUI

import (
	C "GTUI/Core"
	"GTUI/Core/Color"
	iTerminal "GTUI/Terminal"
	implTerm "GTUI/Terminal/impl"
	"strings"
)

type GTui struct {
	GlobalColor Color.Color
	term iTerminal.ITerminal
	buff map[string]C.IEntity
}

func NewGtui() (*GTui,error) {
	term:=&implTerm.Terminal{}
	e:=term.Start()
	if e!=nil {
		return nil,e
	}
	return &GTui{term: term,buff: make(map[string]C.IEntity)},nil
}

func (c *GTui) Close() {
	c.IResetGlobalColor()
	c.term.Clear()
	c.term.Stop()
}

func (c *GTui) InsertEntity(entity C.IEntity) {
	c.buff[entity.GetName()]=entity
}

func (c *GTui) IRefreshAll() {
	var str strings.Builder
	str.WriteString(c.GlobalColor.GetAnsiColor())
	for _, b := range c.buff {
		str.WriteString (b.GetAnsiCode())
		str.WriteString (c.GlobalColor.GetAnsiColor())
	}
	c.term.PrintStr(str.String())
}

func (c *GTui) IRefresh(name string) {
	b := c.buff[name]
	str := b.GetAnsiCode()
	c.term.PrintStr(str)
}
