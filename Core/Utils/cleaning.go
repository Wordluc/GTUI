package Utils

import "strings"

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
