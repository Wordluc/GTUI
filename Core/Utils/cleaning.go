package Utils

import (
	"strconv"
	"strings"
)

func GetAnsiMoveTo(x, y int) string {
	return "\033[" + strconv.Itoa(y+1) + ";" + strconv.Itoa(x+1) + "H"
}
func GetAnsiClear(x,y,sizeX,sizeY int)string {
	var str strings.Builder
	str.WriteString(SaveCursor)
	for iy :=range sizeY {
		for ix := range sizeX {
      str.WriteString(GetAnsiMoveTo(ix+x, iy+y))
			str.WriteString(" ")
		}
	}
	str.WriteString(RestoreCursor)
	return str.String()
}
