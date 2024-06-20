package GTUI

import (
	"GTUI/internal/Color"
	"strconv"
)

func (c *Core) SetColor(color Color.Color) {
	c.term.PrintStr(Color.SetColor(color))
	c.GlobalColor = color
}

func (c *Core) ResetColor() {
	c.term.PrintStr(Color.ResetColor())
	c.GlobalColor = Color.Color{Foreground: Color.ResetF, Background: Color.ResetB}
}

func (c *Core) Clear() {
	c.term.Clear()
}
func (c *Core) Len() (int, int) {
	return c.term.Len()
}

func (c *Core) GetCursor() (int, int) {
	return c.term.GetCursor()
}

func (c *Core) SetCursor(x, y int) {
	c.term.PrintStr("\033[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H")
}
