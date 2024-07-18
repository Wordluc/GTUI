package Drawing

import (
	"GTUI/Core/Color"
	U "GTUI/Core/Utils"
	"strings"
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
	}
}

func (r *RectangleFull) Touch() {
	r.isChanged = true
	r.border.Touch()
}
func (r *RectangleFull) GetAnsiCode(defaultColor Color.Color) string {
	if r.isChanged {
		var str strings.Builder
		str.WriteString(r.getAnsiRectangle(defaultColor))
		str.WriteString(r.border.GetAnsiCode(defaultColor))
		str.WriteString(Color.GetResetColor())
		r.ansiCode = str.String()
		r.isChanged = false
	}
	return r.ansiCode
}

func (r *RectangleFull) SetColor(color Color.Color) {
	r.insideColor = color
	r.Touch()
	r.border.Touch()
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
}

func (r *RectangleFull) SetSize(x, y int) {
	r.Width = x
	r.Height = y
	r.border.SetSize(x, y)
	r.Touch()
}

func (s *RectangleFull) getAnsiRectangle(defaultColor Color.Color) string {
	var str strings.Builder
	var line = strings.Repeat(" ", s.Width-2)
	color:=s.insideColor.GetMixedColor(defaultColor).GetAnsiColor()
	y := 0
	for y < s.Height-2 {
		str.WriteString(color)
		str.WriteString(U.GetAnsiMoveTo(s.XPos+1, y+s.YPos+1))
		str.WriteString(line)
		y++
	}
	return str.String()
}
