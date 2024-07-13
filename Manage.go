package GTUI

import (
	"GTUI/internal/Color"
	"strings"
)


func (c *Core) ISetGlobalColor(color Color.Color) {
	c.term.PrintStr(Color.GetAnsiColor(color))
	c.GlobalColor = color

}
func (c *Core) IResetGlobalColor() {
	c.term.PrintStr(Color.GetResetColor())
	c.GlobalColor = Color.Color{Foreground: Color.ResetF, Background: Color.ResetB}
}

func (c *Core) IClearAll() {
	c.term.Clear()
	c.buff = []string{}
}

func (c *Core) ILen() (int, int) {
	return c.term.Len()
}

func (c *Core) IGetCursor() (int, int) {
	x,y:=c.term.GetCursor()
	x++
	y++
	return x, y
}

func (c *Core) ISetCursor(x, y int) {
	c.term.SetCursor(x+1, y+1)
}

func (c *Core) IInsertToBuff(str string) {
	(*c).buff = append(c.buff, str)
}

func (c *Core) IRefresh() {
 	var str strings.Builder
	for _, s := range c.buff {
		str.WriteString(s)
	}
	(*c).buff	= []string{}
	c.term.PrintStr(str.String())
}

