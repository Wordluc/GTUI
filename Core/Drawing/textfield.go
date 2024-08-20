package Drawing

import (
	"GTUI/Core/Utils"
	U "GTUI/Core/Utils"
	"GTUI/Core/Utils/Color"
	"strings"
)

type TextField struct {
	isChanged bool
	color     Color.Color
	ansiCode  string
	XPos      int
	YPos      int
	text      strings.Builder
	visible   bool
	line      int
}

func CreateTextField(x, y int) *TextField {
	return &TextField{
		XPos:      x,
		YPos:      y,
		isChanged: true,
		color:     Color.GetDefaultColor(),
		visible:   true,
		text:      strings.Builder{},
	}
}

func (s *TextField) Type(text string) {
	if text == "\n" {
		s.text.WriteString(" "+Utils.GetAnsiMoveTo(s.XPos, s.YPos+s.line+1))
		s.line++
	} else {
		s.text.WriteString(text)
	}
	s.Touch()
}

func (s *TextField) GetAnsiCode(defaultColor Color.Color) string {
	if !s.visible {
		return ""
	}
	if s.isChanged {
		s.isChanged = false
		s.ansiCode = s.getAnsiTextField(defaultColor)
	}
	return s.ansiCode
}

func (s *TextField) SetPos(x, y int) {
	s.XPos = x
	s.YPos = y
}

func (s *TextField) GetPos() (int, int) {
	return s.XPos, s.YPos
}

func (s *TextField) SetColor(color Color.Color) {
	s.color = color
	s.Touch()
}
func (s *TextField) ClearText() {
	s.text.Reset()
	s.Touch()
	s.line=0
}
func (s *TextField) SetVisibility(visible bool) {
	s.visible = visible
}
func (s *TextField) GetVisibility() bool {
	return s.visible
}

func (s *TextField) getAnsiTextField(defaultColor Color.Color) string {
	var str strings.Builder
	str.WriteString(U.GetAnsiMoveTo(s.XPos, s.YPos))
	str.WriteString(s.color.GetMixedColor(defaultColor).GetAnsiColor())
	str.WriteString(s.text.String())
	str.WriteString(Color.GetResetColor())
	s.Touch()
	return str.String()
}

func (s *TextField) Touch() {
	s.isChanged = true
}
