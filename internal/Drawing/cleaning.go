package Drawing

import "strings"

func GetAnsiClear(x, y,xpos, ypos int)string {
	var str strings.Builder
	str.WriteString(saveCursor)
	for ix := 0; ix < x; ix++ {
		for iy := 0; iy < y; iy++ {
      str.WriteString(getAnsiMoveTo(ix+xpos, iy+ypos))
			str.WriteString(" ")
		}
	}
	str.WriteString(restoreCursor)
	return str.String()
}
