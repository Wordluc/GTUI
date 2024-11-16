package GTUI

import (
	"strings"
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Utils"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

func (c *Gtui) ISetGlobalColor(color Color.Color) {
	c.globalColor = color
	c.IRefreshAll()
	c.entityTree.Execute(func(node *Core.TreeNode[Core.IEntity])bool {
		for _, child := range node.GetElements() {
			if drawing,ok:=child.(Core.IDrawing);ok{
				drawing.Touch()
			}
		}
		return true
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
}

func (c *Gtui) IClear() {
	c.term.Clear()
	if c.cursorVisibility {
		c.term.PrintStr(c.term.ShowCursor())
	}else{
		c.term.PrintStr(c.term.HideCursor())
	}
}

func (c *Gtui) ClearZone(x, y, xSize, ySize int) string {
	whiteLine:=strings.Repeat(" ",xSize)
	r:=strings.Builder{}
	for i:=0;i<ySize;i++{
		r.WriteString(Utils.GetAnsiMoveTo(x,y+i))
		r.WriteString(whiteLine)
	}
	return r.String()
}
