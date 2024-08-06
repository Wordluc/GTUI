package Utils

import (
	"strconv"
	"strings"
)

func GetAnsiMoveTo(y, x int) string {
	return "\033[" + strconv.Itoa(x+1) + ";" + strconv.Itoa(y+1) + "H"
}
func GetAnsiClear(x, y,xpos, ypos int)string {
	var str strings.Builder
	str.WriteString(SaveCursor)
	for ix := 0; ix < x; ix++ {
		for iy := 0; iy < y; iy++ {
      str.WriteString(GetAnsiMoveTo(ix+xpos, iy+ypos))
			str.WriteString(" ")
		}
	}
	str.WriteString(RestoreCursor)
	return str.String()
}
