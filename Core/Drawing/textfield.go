package Drawing

import (
	"errors"
	"strings"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/EventManager"
	U "github.com/Wordluc/GTUI/Core/Utils"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

type TextField struct {
	isChanged bool
	color     Color.Color
	ansiCode  string
	XPos      int
	YPos      int
	text      string
	visible   bool
	layer     Core.Layer
}

func CreateTextField(x, y int,text string) *TextField {
	return &TextField{
		XPos:      x,
		YPos:      y,
		isChanged: true,
		color:     Color.GetDefaultColor(),
		visible:   true,
		text:      text,
		layer:     Core.L1,
	}
}

func (s *TextField) GetSize() (int,int) {
	return len(s.text),1
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
	s.Touch()
	EventManager.Call(EventManager.ReorganizeElements,[]any{s})
}
func (s *TextField) SetLayer(layer Core.Layer) error{
	if s.layer == layer{
		return errors.New("layer already set")
	}
	s.layer = layer
	s.Touch()
	EventManager.Call(EventManager.ReorganizeElements,[]any{s})
	return nil
}
func (s *TextField) GetLayer() Core.Layer {
	return s.layer
}
func (s *TextField) GetPos() (int, int) {
	return s.XPos, s.YPos
}

func (s *TextField) SetColor(color Color.Color) {
	s.color = color
	s.Touch()
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
	str.WriteString(s.text)
	str.WriteString(Color.GetResetColor())
	return str.String()
}

func (s *TextField) Touch() {
	s.isChanged = true
}
func (s *TextField) IsTouched() bool {
	return s.isChanged
}
