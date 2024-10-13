package GTUI

import (
	"strings"

	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/Utils"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

func (c *Gtui) ISetGlobalColor(color Color.Color) {
	c.globalColor = color
	c.IRefreshAll()
	c.IClear()
	c.drawingManager.Execute(func(node *Drawing.TreeNode){
		for _, child := range node.GetElements() {
			child.Touch()
		}
	})
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

func (c *Gtui) ClearZone(x, y, xSize, ySize int) {
	whiteLine:=strings.Repeat(" ",xSize)
	r:=strings.Builder{}
	for i:=0;i<ySize;i++{
		r.WriteString(whiteLine)
		r.WriteString(Utils.GetAnsiMoveTo(x,y+i))
	}
	c.term.PrintStr(r.String())
}
