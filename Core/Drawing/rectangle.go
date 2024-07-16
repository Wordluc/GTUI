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
	name      string
	XPos      int
	YPos      int
	Width     int
	Height    int
}

func CreateRectangle(name string,x, y, width, height int) *Rectangle {
	return &Rectangle{
		ansiCode: "",
		name:     name,
		XPos:     x,
		YPos:     y,
		Width:    width,
		isChanged:    true,
		Height:   height,
	}
}

func (r *Rectangle) Touch() {
	r.isChanged = true
}

func (r *Rectangle) GetAnsiCode() string {
	if r.isChanged {
		r.ansiCode = r.getAnsiRectangle()
		r.isChanged = false
	}
	return r.ansiCode
}

func (l *Rectangle) SetPos(x, y int) {
	l.XPos = x
	l.YPos = y
	l.Touch()
}
func (l *Rectangle) SetSize(x, y int) error{
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
func (l *Rectangle) GetPos() (int,int) {
	return l.XPos, l.YPos
}
func (s *Rectangle) GetName() string {
	return s.name
}
func (s *Rectangle) getAnsiRectangle() string {
	var horizontal = strings.Repeat(U.HorizontalLine, s.Width-2)
	var builStr strings.Builder
	builStr.WriteString(U.SaveCursor)
	builStr.WriteString(U.GetAnsiMoveTo(s.XPos, s.YPos))
	builStr.WriteString(s.color.GetAnsiColor())
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
