package GTUI


import (
	"GTUI/Core/Utils/Color"
)

func (c *Gtui) ISetGlobalColor(color Color.Color) {
	c.globalColor = color
	c.IRefreshAll()
	c.IClear()
	for _, b := range c.buff {
		b.Touch()
	}
}
func (c *Gtui) IResetGlobalColor() {
	c.ISetGlobalColor(Color.GetDefaultColor())
}

func (c *Gtui) ILen() (int, int) {
	return c.term.Size()
}

func (c *Gtui) ISetCursor(x, y int) {
	c.term.SetCursor(x+1, y+1)
}

func (c *Gtui) IClear() {
	c.term.Clear()
}
