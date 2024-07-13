package GTUI

import (
	"GTUI/internal/Color"
	"GTUI/internal/Drawing"
	"strings"
)
type MetaData struct {
	Color Color.Color
}


func (c *Core) InsertStrs(strs ...string) {
	var strBuilder  strings.Builder
	for _, str := range strs {
		strBuilder.WriteString(str)
	}
	c.IInsertToBuff(strBuilder.String())
}

func (c *Core) InsertStrsf( opt MetaData ,strs ...string) {
	var strBuilder  strings.Builder
	strBuilder.WriteString(Color.GetAnsiColor(opt.Color))
	for _, str := range strs {
		strBuilder.WriteString(str)
	}
	strBuilder.WriteString(Color.GetAnsiColor(c.GlobalColor))

	c.IInsertToBuff(strBuilder.String())
}

func (c *Core) InsertColor(color Color.Color,str string) {
	c.IInsertToBuff(Color.GetAnsiColor(color)+str)
	c.GlobalColor = color
}
func (c *Core) ClearPortion(x,y,xpos, ypos int) {
	c.IInsertToBuff(Drawing.GetAnsiClear(x, y, xpos, ypos))
}
