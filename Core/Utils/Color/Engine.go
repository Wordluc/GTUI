package Color

import (
	"strconv"
)

type Color struct {
	foreground IColor
	Background IColor
}

func (c *Color) SetForeground(color IColor)  {
	c.foreground=color
}

func (c *Color) SetBackground(color IColor)  {
	c.Background=gBg(color)
}

func (c Color) isEqual(other Color) bool {
	return c.foreground == other.foreground && c.Background == other.Background
}

func Get(foreGround IColor, backGround IColor) Color {
	return Color{foreground: foreGround, Background: gBg(backGround)}
}

func gBg(i IColor) IColor {
	return i + 10
}

func GetNoneColor() Color {
	return Color{foreground: None, Background:gBg(None)}
}

func GetDefaultColor() Color {
	return Color{foreground: Reset, Background: gBg(Reset)}
}

func (c Color) GetMixedColor(defaultColor Color) Color {
	if GetNoneColor().isEqual(c) {
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
	if GetNoneColor().isEqual(c) {
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
