package Drawing

import (
	"GTUI/Core/Color"
	U "GTUI/Core/Utils"
	"errors"
	"strings"
)

type Line struct {
	isChanged bool
	ansiCode  string
	name      string
	Color     Color.Color
	XPos      int
	YPos      int
	Len       int
	angle     int16
}

func CreateLine(name string, x, y, len int, angle int16) *Line {
	return &Line{
		ansiCode:  "",
		XPos:      x,
		YPos:      y,
		Len:       len,
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

func (r *Line) SetAngle(angle int16) {
	switch angle {
	case 0, 45, 90, 270:
		r.angle = angle
	default:
		return
	}
}

func (l *Line) SetPos(x, y int) {
	l.XPos = x
	l.YPos = y
}
func (l *Line) SetLen(len int)error {
	if len < 0 {
		return errors.New("len < 0")
	}
	l.Len = len
	return nil
}
func (l *Line) GetPos() (int,int) {
	return l.XPos, l.YPos
}
func (l *Line) getAnsiLine() string {
	var str *strings.Builder = new(strings.Builder)
	str.WriteString(U.SaveCursor)
	str.WriteString(l.Color.GetAnsiColor())
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

func ansiLineAngle0(l *Line, str *strings.Builder) {
	str.WriteString(U.GetAnsiMoveTo(l.XPos, l.YPos))
	str.WriteString(strings.Repeat(U.HorizontalLine, l.Len))
}

func ansiLineAngle270(l *Line, str *strings.Builder) {
	for i := 1; i < l.Len; i++ {
		str.WriteString(U.GetAnsiMoveTo(l.XPos+i, l.YPos+i))
		str.WriteString(U.HorizontalLine)
	}
}

func ansiLineAngle90(l *Line, str *strings.Builder) {
	for i := 0; i < l.Len; i++ {
		str.WriteString(U.GetAnsiMoveTo(l.XPos, l.YPos+i))
		str.WriteString(U.VerticalLine)
	}
}

func ansiLineAngle45(l *Line, str *strings.Builder) {
	for i := 0; i < l.Len; i++ {
		str.WriteString(U.GetAnsiMoveTo(l.XPos+i, l.YPos-i))
		str.WriteString(U.HorizontalLine)
	}
}
