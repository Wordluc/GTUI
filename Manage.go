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

func (c *Gtui) Size() (int, int) {
	return c.term.Size()
}
func (c *Gtui) SetVisibilityCursor(visibility bool) {
	c.cursorVisibility = visibility
	if c.cursorVisibility {
		c.term.ShowCursor()
	}else{
		c.term.HideCursor()
	}
}
func (c *Gtui) IClear() {
	c.term.Clear()
	if c.cursorVisibility {
		c.term.ShowCursor()
	}else{
		c.term.HideCursor()
	}
}
