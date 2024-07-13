package Drawing

import (
	"errors"
	"strings"
)

type Line struct {
	XPos int
	YPos int
	Len  int
	Angle int16
}

func GetAnsiLine(l Line) (string,error) {
	if !(l.YPos > 0 && l.XPos > 0 && l.Len > 0) {
		return "",errors.New("XPos and YPos must be greater than 0 and Len must be greater than 0")
	}
	var str *strings.Builder = new(strings.Builder)
	str.WriteString(saveCursor)
	switch l.Angle {
	case 0:
		ansiLineAngle0(l,str)
	case 45:
		ansiLineAngle45(l,str)
	case 90:
		ansiLineAngle90(l,str)
	case 270:
		ansiLineAngle270(l,str)
	default:
		return "",errors.New("Angle must be 0, 45, 90 or 270")
	}
	str.WriteString(restoreCursor)
	return str.String(),nil
}

func ansiLineAngle0(l Line,str *strings.Builder) {
	str.WriteString( strings.Repeat(" ", l.Len))
}

func ansiLineAngle270(l Line,str *strings.Builder){
		for i := 1; i < l.Len; i++ {
			str.WriteString(getAnsiMoveTo(l.XPos+i, l.YPos+i))
			str.WriteString(topLine)
		}
}

func ansiLineAngle90(l Line,str *strings.Builder){
	for i := 0; i < l.Len; i++ {
		str.WriteString(getAnsiMoveTo(l.XPos, l.YPos+i))
		str.WriteString(verticalLine)
	}
}

func ansiLineAngle45(l Line,str *strings.Builder) {
		for i := 0; i < l.Len; i++ {
			str.WriteString(getAnsiMoveTo(l.XPos+i, l.YPos-i))
			str.WriteString(topLine)
	}
}
