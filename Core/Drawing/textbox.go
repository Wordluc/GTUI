package Drawing

import (
	"GTUI/Core/Color"
	U "GTUI/Core/Utils"
	"strings"
)

type TextBox struct {
	isChanged    bool
	defaultColor Color.Color
	ansiCode     string
	name         string
	xPos         int
	yPos         int
	text         strings.Builder
}

func CreateTextBox(name string, x, y int) *TextBox {
	return &TextBox{
		name:         name,
		xPos:         x,
		yPos:         y,
		isChanged:    true,
	}
}

func (s *TextBox) Type(text string) {
	  s.text.WriteString(text)
}

func (s *TextBox) GetAnsiCode() string {
	if s.isChanged {
		s.isChanged = false
		s.ansiCode = s.getAnsiTextBox()
	}
	return s.ansiCode
}

func (s *TextBox) getAnsiTextBox() string {
	var str strings.Builder
	str.WriteString(U.GetAnsiMoveTo(s.xPos, s.yPos))
	str.WriteString(s.text.String())
	return str.String()
}

func (s *TextBox) Touch() {
	s.isChanged = true
}

func (s *TextBox) GetName() string {
	return s.name
}

func (s *TextBox) SetColor(c Color.Color) {
	s.defaultColor = c
}

func (s *TextBox) GetColor() Color.Color {
	return s.defaultColor
}
