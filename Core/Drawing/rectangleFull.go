package Drawing

import (
	"errors"
	"strings"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/EventManager"
	U "github.com/Wordluc/GTUI/Core/Utils"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

type RectangleFull struct {
	isChanged   bool
	ansiCode    string
	border      *Rectangle
	insideColor Color.Color
	XPos        int
	YPos        int
	Width       int
	Height      int
	visible     bool
	layer       Core.Layer
}

func CreateRectangleFull(x, y, width, height int) *RectangleFull {
	border := CreateRectangle(x, y, width, height)
	return &RectangleFull{
		ansiCode:    "",
		XPos:        x,
		YPos:        y,
		Width:       width,
		Height:      height,
		isChanged:   true,
		insideColor: Color.GetNoneColor(),
		border:      border,
		visible:     true,
	}
}

func (r *RectangleFull) Touch() {
	r.isChanged = true
	r.border.Touch()
}

func (r *RectangleFull) IsTouched() bool {
	return r.isChanged || r.border.IsTouched()
}
func (r *RectangleFull) GetAnsiCode(defaultColor Color.Color) string {
	if !r.visible {
		return ""
	}
	if r.isChanged {
		var str strings.Builder
		str.WriteString(r.border.GetAnsiCode(defaultColor))
		str.WriteString(r.getAnsiRectangle(defaultColor))
		str.WriteString(Color.GetResetColor())
		r.ansiCode = str.String()
		r.isChanged = false
	}
	return r.ansiCode
}

func (r *RectangleFull) SetInsideColor(color Color.ColorValue) {
	r.insideColor = Color.Get(Color.None, color)
	r.Touch()
}

func (r *RectangleFull) SetBorderColor(color Color.Color) {
	r.border.SetColor(color)
	r.Touch()
}

func (r *RectangleFull) GetColor() Color.Color {
	return r.insideColor
}

func (r *RectangleFull) GetPos() (int, int) {
	return r.XPos, r.YPos
}

func (r *RectangleFull) GetBorder() *Rectangle {
	return r.border
}

func (r *RectangleFull) SetPos(x, y int) {
	r.XPos = x
	r.YPos = y
	r.border.SetPos(x, y)
	r.Touch()
	EventManager.Call(EventManager.ReorganizeElements, r)
}

func (b *RectangleFull) SetLayer(layer Core.Layer) error {
	if layer < 0 {
		return errors.New("layer can't be negative")
	}
	b.layer = layer
	b.Touch()
	EventManager.Call(EventManager.ReorganizeElements, b)
	return nil
}
func (c *RectangleFull) GetLayer() Core.Layer {
	return c.layer
}
func (r *RectangleFull) SetSize(x, y int) {
	r.Width = x
	r.Height = y
	r.border.SetSize(x, y)
	r.Touch()
}

func (r *RectangleFull) GetSize() (int, int) {
	return r.Width, r.Height
}

func (s *RectangleFull) SetVisibility(visible bool) {
	s.visible = visible
	s.Touch()
	EventManager.Call(EventManager.ForceRefresh)
}

func (s *RectangleFull) GetVisibility() bool {
	return s.visible
}

func (s *RectangleFull) getAnsiRectangle(defaultColor Color.Color) string {
	var str strings.Builder
	var line = strings.Repeat(" ", s.Width-2)
	color := s.insideColor.GetMixedColor(defaultColor).GetAnsiColor()
	y := 0
	for y < s.Height-2 {
		str.WriteString(color)
		str.WriteString(U.GetAnsiMoveTo(s.XPos+1, y+s.YPos+1))
		str.WriteString(line)
		y++
	}
	return str.String()
}
