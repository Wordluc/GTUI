package Drawing

import (
	"GTUI/Core/Color"
	U "GTUI/Core/Utils"
	"strings"
)

type Line struct {
	isChanged bool
	ansiCode  string
	name      string
	color     Color.Color
	xPos      int
	yPos      int
	len       int
	angle     int16
}

func CreateLine(name string, x, y, len int, angle int16) *Line {
	return &Line{
		ansiCode:  "",
		xPos:      x,
		yPos:      y,
		len:       len,
		angle:     angle,
		isChanged: true,
	}
}

func (r *Line) Touch() {
	r.isChanged = true
}

func (r *Line) GetAnsiCode() string {
	if r.isChanged {
		r.ansiCode = r.getAnsiLine()
		r.isChanged = false
	}
	return r.ansiCode
}

func (r *Line) GetName() string {
	return r.name
}

func (s *Line) GetColor() Color.Color {
	return s.color
}

func (s *Line) SetColor(c Color.Color) {
	s.color = c
}

func (l Line) getAnsiLine() string {
	var str *strings.Builder = new(strings.Builder)
	str.WriteString(U.SaveCursor)
	switch l.angle {
	case 0:
		ansiLineAngle0(l, str)
	case 45:
		ansiLineAngle45(l, str)
	case 90:
		ansiLineAngle90(l, str)
	case 270:
		ansiLineAngle270(l, str)
	}
	str.WriteString(U.RestoreCursor)
	return str.String()
}

func ansiLineAngle0(l Line, str *strings.Builder) {
	str.WriteString(strings.Repeat(U.HorizontalLine, l.len))
}

func ansiLineAngle270(l Line, str *strings.Builder) {
	for i := 1; i < l.len; i++ {
		str.WriteString(U.GetAnsiMoveTo(l.xPos+i, l.yPos+i))
		str.WriteString(U.HorizontalLine)
	}
}

func ansiLineAngle90(l Line, str *strings.Builder) {
	for i := 0; i < l.len; i++ {
		str.WriteString(U.GetAnsiMoveTo(l.xPos, l.yPos+i))
		str.WriteString(U.VerticalLine)
	}
}

func ansiLineAngle45(l Line, str *strings.Builder) {
	for i := 0; i < l.len; i++ {
		str.WriteString(U.GetAnsiMoveTo(l.xPos+i, l.yPos-i))
		str.WriteString(U.HorizontalLine)
	}
}
