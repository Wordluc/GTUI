package Color

import (
	"strconv"
)

type Color struct {
	Foreground iColor
	Background iColor
}   
func GetDefaultColor() Color {
	return Color{Foreground: ResetF, Background: ResetB}
}

func (c Color) GetAnsiColor() string{
	return ("\x1b[;" + strconv.Itoa(int(c.Foreground)) + ";" + strconv.Itoa(int(c.Background)) + "m")
}

func GetResetColor() string {
	return ("\033[;0;0m")
}
