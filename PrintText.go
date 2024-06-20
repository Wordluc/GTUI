package GTUI

import (
	"GTUI/internal/Color"
	"strings"
)
type MetaData struct {
	Color Color.Color
}

func (c *Core) Print(str string) {
	c.term.PrintStr(str)
}

func (c *Core) Prints(strs ...string) {
	var strBuilder  strings.Builder
	for _, str := range strs {
		strBuilder.WriteString(str)
	}
	c.term.PrintStr(strBuilder.String())
}

func (c *Core) Printf(str string, opt MetaData) {
	var strBuilder  strings.Builder
	strBuilder.WriteString(Color.SetColor(opt.Color))
	strBuilder.WriteString(str)
	strBuilder.WriteString(Color.SetColor(c.GlobalColor))

	c.term.PrintStr(strBuilder.String())
}

func (c *Core) Printfs( opt MetaData ,strs ...string) {
	var strBuilder  strings.Builder
	strBuilder.WriteString(Color.SetColor(opt.Color))
	for _, str := range strs {
		strBuilder.WriteString(str)
	}
	strBuilder.WriteString(Color.SetColor(c.GlobalColor))

	c.term.PrintStr(strBuilder.String())
}


