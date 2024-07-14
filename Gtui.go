package GTUI

import (
	"GTUI/Core/Color"
	iTerminal "GTUI/Terminal"
	implTerm "GTUI/Terminal/impl"
	C "GTUI/Core"
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
	return &GTui{term: term},nil
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
	for _, b := range c.buff {
		str := b.GetAnsiCode()
		c.term.PrintStr(str)
	}
}

func (c *GTui) IRefresh(name string) {
	b := c.buff[name]
	str := b.GetAnsiCode()
	c.term.PrintStr(str)
}
