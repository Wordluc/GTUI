package GTUI

import (
	"GTUI/Core/Color"
)

func (c *GTui) ISetGlobalColor(color Color.Color) {
	c.globalColor = color
	c.IRefreshAll()
	c.IClear()
	for _, b := range c.buff {
		b.Touch()
	}
}
func (c *GTui) IResetGlobalColor() {
	c.ISetGlobalColor(Color.GetDefaultColor())
}

func (c *GTui) ILen() (int, int) {
	return c.term.Size()
}

func (c *GTui) IGetCursor() (int, int) {
	x, y := c.term.GetCursor()
	x++
	y++
	return x, y
}

func (c *GTui) ISetCursor(x, y int) {
	c.term.SetCursor(x+1, y+1)
}

func (c *GTui) IClear() {
	c.term.Clear()
}
