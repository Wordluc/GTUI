package Drawing

import "strconv"

func getAnsiMoveTo(y, x int) string {
	return "\033[" + strconv.Itoa(x+1) + ";" + strconv.Itoa(y+1) + "H"
}
