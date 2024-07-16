package Drawing

import (
	"GTUI/Core/Color"
	U "GTUI/Core/Utils"
	"strings"
)

type TextBox struct {
	isChanged    bool
	Color Color.Color
	ansiCode     string
	name         string
	XPos         int
	YPos         int
	text         strings.Builder
}

func CreateTextBox(name string, x, y int) *TextBox {
	return &TextBox{
		name:         name,
		XPos:         x,
		YPos:         y,
		isChanged:    true,
	}
}

func (s *TextBox) Type(text string) {
	  s.text.WriteString(text)
		s.Touch()
}

func (s *TextBox) GetAnsiCode() string {
	if s.isChanged {
		s.isChanged = false
		s.ansiCode = s.getAnsiTextBox()
	}
	return s.ansiCode
}

func (s *TextBox) SetPos(x, y int) {
	s.XPos = x
	s.YPos = y
}

func (s *TextBox) GetPos() (int, int) {
	return s.XPos, s.YPos
}

func (s *TextBox) getAnsiTextBox() string {
	var str strings.Builder
	str.WriteString(U.GetAnsiMoveTo(s.XPos, s.YPos))
	str.WriteString(s.Color.GetAnsiColor())
	str.WriteString(s.text.String())
	return str.String()
}

func (s *TextBox) Touch() {
	s.isChanged = true
}

func (s *TextBox) GetName() string {
	return s.name
}

