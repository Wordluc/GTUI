package Color

import (
	"strconv"
)

type Color struct {
	foreground ColorValue
	Background ColorValue
}

func (c *Color) SetForeground(color ColorValue)  {
	c.foreground=color
}

func (c *Color) SetBackground(color ColorValue)  {
	c.Background=gBg(color)
}

func (c Color) IsEqual(other Color) bool {
	return c.foreground == other.foreground && c.Background == other.Background
}

func Get(foreGround ColorValue, backGround ColorValue) Color {
	return Color{foreground: foreGround, Background: gBg(backGround)}
}

func gBg(i ColorValue) ColorValue {
	return i + 10
}

func GetNoneColor() Color {
	return Color{foreground: None, Background:gBg(None)}
}

func GetDefaultColor() Color {
	return Color{foreground: Reset, Background: gBg(Reset)}
}
//Get the color that is the combination of the default color and the current color,if the a part of default color is not set, the default color will be used
func (c Color) GetMixedColor(defaultColor Color) Color {
	if GetNoneColor().IsEqual(c) {
		return GetNoneColor()
	}
	color:=c
	if c.Background == gBg(None) {
		color.Background = defaultColor.Background
	}
	if c.foreground == None {
		color.foreground = defaultColor.foreground
	}
	return color
}

func (c Color) GetAnsiColor() string {
	if GetNoneColor().IsEqual(c) {
		return ""
	}
	bg := "0"
	if c.Background != gBg(None) {
		bg = strconv.Itoa(int(c.Background))
	}
	fg := "0"
	if c.foreground != None {
		fg = strconv.Itoa(int(c.foreground))
	}
	return ("\x1b[;" + fg + ";" + bg + "m")
}

func GetResetColor() string {
	return ("\033[;0;0m")
}
