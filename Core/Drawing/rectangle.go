package Drawing

import (
	"GTUI/Core/Color"
	U "GTUI/Core/Utils"
	"errors"
	"strings"
)

type Rectangle struct {
	isChanged bool
	ansiCode  string
	color     Color.Color
	XPos      int
	YPos      int
	Width     int
	Height    int
}

func CreateRectangle(x, y, width, height int) *Rectangle {
	return &Rectangle{
		ansiCode:  "",
		XPos:      x,
		YPos:      y,
		Width:     width,
		isChanged: true,
		Height:    height,
		color:     Color.GetNoneColor(),
	}
}

func (r *Rectangle) Touch() {
	r.isChanged = true
}

func (r *Rectangle) GetAnsiCode(defaultColor Color.Color) string {
	if r.isChanged {
		r.ansiCode = r.getAnsiRectangle(defaultColor)
		r.isChanged = false
	}
	return r.ansiCode
}

func (l *Rectangle) SetPos(x, y int) {
	l.XPos = x
	l.YPos = y
	l.Touch()
}
func (l *Rectangle) SetSize(x, y int) error {
	if x < 0 || y < 0 {
		return errors.New("invalid size")
	}
	l.Width = x
	l.Height = y
	l.Touch()
	return nil
}

func (c *Rectangle) SetColor(color Color.Color) {
	c.color = color
	c.Touch()
}
func (l *Rectangle) GetPos() (int, int) {
	return l.XPos, l.YPos
}
func (s *Rectangle) getAnsiRectangle(defaultColor Color.Color) string {
	var horizontal = strings.Repeat(U.HorizontalLine, s.Width-2)
	var color = s.color.GetMixedColor(defaultColor).GetAnsiColor()
	var builStr strings.Builder
	builStr.WriteString(color)
	builStr.WriteString(U.SaveCursor)
	builStr.WriteString(U.GetAnsiMoveTo(s.XPos, s.YPos))
	builStr.WriteString(U.TopLeftCorner)
	builStr.WriteString(horizontal)
	for i := 1; i < s.Height-1; i++ {
		builStr.WriteString(U.GetAnsiMoveTo(s.XPos, s.YPos+i))
		builStr.WriteString(U.VerticalLine)
	}
	builStr.WriteString(U.GetAnsiMoveTo(s.XPos, s.YPos+s.Height-1))
	builStr.WriteString(U.BottomLeftCorner)
	builStr.WriteString(horizontal)
	builStr.WriteString(U.BottomRightCorner)
	for i := 1; i < s.Height-1; i++ {
		builStr.WriteString(U.GetAnsiMoveTo(s.XPos+s.Width-1, s.YPos+i))
		builStr.WriteString(U.VerticalLine)
	}
	builStr.WriteString(U.GetAnsiMoveTo(s.XPos+s.Width-1, s.YPos))
	builStr.WriteString(U.TopRightCorner)
	builStr.WriteString(U.RestoreCursor)
	return builStr.String()
}
