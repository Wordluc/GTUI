package Drawing

import (
	"GTUI/Core/Utils"
	U "GTUI/Core/Utils"
	"GTUI/Core/Utils/Color"
	"strings"
)

type TextBox struct {
	isChanged bool
	color     Color.Color
	ansiCode  string
	XPos      int
	YPos      int
	text      strings.Builder
	visible   bool
	line      int
}

func CreateTextBox(x, y int) *TextBox {
	return &TextBox{
		XPos:      x,
		YPos:      y,
		isChanged: true,
		color:     Color.GetDefaultColor(),
		visible:   true,
		text:      strings.Builder{},
	}
}

func (s *TextBox) Type(text string) {
	if text == "\n" {
		s.text.WriteString(Utils.GetAnsiMoveTo(s.XPos, s.YPos+s.line+1))
		s.line++
	} else {
		s.text.WriteString(text)
	}
	s.Touch()
}

func (s *TextBox) GetAnsiCode(defaultColor Color.Color) string {
	if !s.visible {
		return ""
	}
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

func (s *TextBox) SetVisibility(visible bool) {
	s.visible = visible
}

func (s *TextBox) GetVisibility() bool {
	return s.visible
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
