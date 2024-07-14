package Drawing

import (
	"GTUI/Core/Color"
	U "GTUI/Core/Utils"
	"strings"
)

type Rectangle struct {
	isChanged bool
	ansiCode  string
	color     Color.Color
	name      string
	xPos      int
	yPos      int
	width     int
	height    int
}

func CreateRectangle(name string,x, y, width, height int) *Rectangle {
	return &Rectangle{
		ansiCode: "",
		name:     name,
		xPos:     x,
		yPos:     y,
		width:    width,
		isChanged:    true,
		height:   height,
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

func (s *Rectangle) GetName() string {
	return s.name
}
func (s *Rectangle) SetColor(c Color.Color) {
	s.color = c
}
func (s *Rectangle) GetColor() Color.Color {
	return s.color
}
func (s *Rectangle) getAnsiRectangle() string {
	var horizontal = strings.Repeat(U.HorizontalLine, s.width-2)
	var builStr strings.Builder
	builStr.WriteString(U.SaveCursor)
	builStr.WriteString(U.GetAnsiMoveTo(s.xPos, s.yPos))
	builStr.WriteString(U.TopLeftCorner)
	builStr.WriteString(horizontal)
	for i := 1; i < s.height-1; i++ {
		builStr.WriteString(U.GetAnsiMoveTo(s.xPos, s.yPos+i))
		builStr.WriteString(U.VerticalLine)
	}
	builStr.WriteString(U.GetAnsiMoveTo(s.xPos, s.yPos+s.height-1))
	builStr.WriteString(U.BottomLeftCorner)
	builStr.WriteString(horizontal)
	builStr.WriteString(U.BottomRightCorner)
	for i := 1; i < s.height-1; i++ {
		builStr.WriteString(U.GetAnsiMoveTo(s.xPos+s.width-1, s.yPos+i))
		builStr.WriteString(U.VerticalLine)
	}
	builStr.WriteString(U.GetAnsiMoveTo(s.xPos+s.width-1, s.yPos))
	builStr.WriteString(U.TopRightCorner)
	builStr.WriteString(U.RestoreCursor)
	return builStr.String()
}
