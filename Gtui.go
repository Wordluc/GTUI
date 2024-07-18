package GTUI

import (
	C "GTUI/Core"
	"GTUI/Core/Color"
	iTerminal "GTUI/Terminal"
	implTerm "GTUI/Terminal/impl"
	"strings"
)

type GTui struct {
	globalColor Color.Color
	term iTerminal.ITerminal
	buff []C.IEntity
}

func NewGtui() (*GTui,error) {
	term:=&implTerm.Terminal{}
	e:=term.Start()
	if e!=nil {
		return nil,e
	}
	return &GTui{term: term,buff: make([]C.IEntity,0)},nil
}

func (c *GTui) Close() {
	c.IResetGlobalColor()
	c.term.Clear()
	c.term.Stop()
}

func (c *GTui) InsertEntity(entity C.IEntity) {
	c.buff=append(c.buff,entity)
}

func (c *GTui) IRefreshAll() {
	var str strings.Builder
	str.WriteString(c.globalColor.GetAnsiColor())
	for _, b := range c.buff {
		str.WriteString (b.GetAnsiCode(c.globalColor))
		str.WriteString(c.globalColor.GetAnsiColor())
	}
	c.term.PrintStr(str.String())
}
