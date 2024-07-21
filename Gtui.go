package GTUI

import (
	C "GTUI/Core"
	"GTUI/Core/Utils/Color"
	iTerminal "GTUI/Terminal"
	implTerm "GTUI/Terminal/impl"
	"strings"
)

type Gtui struct {
	globalColor Color.Color
	term iTerminal.ITerminal
	buff []C.IEntity
}

func NewGtui() (*Gtui,error) {
	term:=&implTerm.Terminal{}
	e:=term.Start()
	if e!=nil {
		return nil,e
	}
	return &Gtui{term: term,buff: make([]C.IEntity,0)},nil
}

func (c *Gtui) Close() {
	c.IResetGlobalColor()
	c.term.Clear()
	c.term.Stop()
}

func (c *Gtui) InsertEntity(entity C.IEntity) {
	c.buff=append(c.buff,entity)
}

func (c *Gtui) IRefreshAll() {
	var str strings.Builder
	str.WriteString(c.globalColor.GetAnsiColor())
	for _, b := range c.buff {
		str.WriteString (b.GetAnsiCode(c.globalColor))
		str.WriteString(c.globalColor.GetAnsiColor())
	}
	c.term.PrintStr(str.String())
}
