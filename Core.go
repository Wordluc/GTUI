package GTUI

import iTerminal "GTUI/Terminal"
import implTerm "GTUI/Terminal/impl"
import "GTUI/internal/Color"

type Core struct {
	GlobalColor Color.Color
	term iTerminal.ITerminal
	buff []string
}

func NewCore() (*Core,error) {
	term:=&implTerm.Terminal{}
	e:=term.Start()
	if e!=nil {
		return nil,e
	}
	return &Core{term: term},nil
}

func (c *Core) Close() {
	c.IResetGlobalColor()
	c.term.Clear()
	c.term.Stop()
}
