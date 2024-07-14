package GTUI


import (
	"GTUI/Core/Color"
)


func (c *GTui) ISetGlobalColor(color Color.Color) {
	c.term.PrintStr(color.GetAnsiColor())
	c.GlobalColor = color

}
func (c *GTui) IResetGlobalColor() {
	c.term.PrintStr(Color.GetResetColor())
	c.GlobalColor = Color.Color{Foreground: Color.ResetF, Background: Color.ResetB}
}

func (c *GTui) ILen() (int, int) {
	return c.term.Len()
}

func (c *GTui) IGetCursor() (int, int) {
	x,y:=c.term.GetCursor()
	x++
	y++
	return x, y
}

func (c *GTui) ISetCursor(x, y int) {
	c.term.SetCursor(x+1, y+1)
}

