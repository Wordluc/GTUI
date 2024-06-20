package Color

import (
	"strconv"
)

func SetColor(c Color) string {
	return ("\x1b[;" + strconv.Itoa(int(c.Foreground)) + ";" + strconv.Itoa(int(c.Background)) + "m")
}

func ResetColor() string {
	return ("\033[;0;0m")
}
