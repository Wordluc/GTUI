package Drawing

import (
	"GTUI/Core/Color"
	U "GTUI/Core/Utils"
	"strings"
)

type TextBox struct {
	isChanged    bool
	color Color.Color
	ansiCode     string
	XPos         int
	YPos         int
	text         strings.Builder
}

func CreateTextBox(x, y int) *TextBox {
	return &TextBox{
		XPos:         x,
		YPos:         y,
		isChanged:    true,
		color:        Color.GetNoneColor(),
	}
}

func (s *TextBox) Type(text string) {
	  s.text.WriteString(text)
		s.Touch()
}

func (s *TextBox) GetAnsiCode(defaultColor Color.Color) string {
	if s.isChanged {
		s.isChanged = false
		s.ansiCode = s.getAnsiTextBox(defaultColor)
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
func (s *TextBox) SetColor(color Color.Color) {
	s.color = color
	s.Touch()
}
func (s *TextBox) getAnsiTextBox(defaultColor Color.Color) string {
	var str strings.Builder
	str.WriteString(U.GetAnsiMoveTo(s.XPos, s.YPos))
	str.WriteString(s.color.GetMixedColor(defaultColor).GetAnsiColor())
	str.WriteString(s.text.String())
	str.WriteString(Color.GetResetColor())
	return str.String()
}

func (s *TextBox) Touch() {
	s.isChanged = true
}
