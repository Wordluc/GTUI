package Drawing

import (
	"errors"
	"strings"
)

type Rectangle struct {
	XPos int
	YPos int
	Width int
	Height int
}
func GetAnsiRectangle(s Rectangle) (string,error) {
	if !(s.YPos > 0 && s.XPos > 0 && s.Width > 0 && s.Height > 0) {
		return "",errors.New("XPos and YPos must be greater than 0 and Width and Height must be greater than 0")
	}
	var horizontal = strings.Repeat(topLine, s.Width-2)
	var builStr strings.Builder
	builStr.WriteString(saveCursor)
	builStr.WriteString(getAnsiMoveTo(s.XPos, s.YPos))
	builStr.WriteString(topLeftCorner)
	builStr.WriteString(horizontal)
	for i := 1; i < s.Height-1; i++ {
		builStr.WriteString(getAnsiMoveTo(s.XPos, s.YPos+i))
		builStr.WriteString(verticalLine)
	}
	builStr.WriteString(getAnsiMoveTo(s.XPos, s.YPos+s.Height-1))
	builStr.WriteString(bottomLeftCorner)
	builStr.WriteString(horizontal)
	builStr.WriteString(bottomRightCorner)
	for i := 1; i < s.Height-1; i++ {
		builStr.WriteString(getAnsiMoveTo(s.XPos+s.Width-1, s.YPos+i))
		builStr.WriteString(verticalLine)
	}
	builStr.WriteString(getAnsiMoveTo(s.XPos+s.Width-1, s.YPos))
	builStr.WriteString(topRightCorner)
	builStr.WriteString(restoreCursor)
	return builStr.String(),nil
}
