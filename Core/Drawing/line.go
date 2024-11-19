package Drawing

import (
	"errors"
	"math"
	"strings"

	"github.com/Wordluc/GTUI/Core/EventManager"
	"github.com/Wordluc/GTUI/Core"
	U "github.com/Wordluc/GTUI/Core/Utils"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

type Line struct {
	isChanged bool
	ansiCode  string
	color     Color.Color
	xPos      int
	yPos      int
	Len       int
	angle     int16
	visible   bool
	layer     Core.Layer
}

func CreateLine(x, y, len int, angle int16) *Line {
	return &Line{
		ansiCode:  "",
		xPos:      x,
		yPos:      y,
		Len:       len,
		angle:     angle,
		isChanged: true,
		color:     Color.GetNoneColor(),
		visible:   true,
	}
}

func (r *Line) Touch() {
	r.isChanged = true
}
func (r *Line) IsTouched() bool {
	return r.isChanged
}
func (r *Line) GetSize() (int, int) {
   switch(r.angle){
		case 0:return r.Len,1
		case 45:
			var x = int(math.Sqrt(float64(r.Len^2)+float64(r.Len^2)))
			return x,x
		case 90:return 1,r.Len
		default:return 0,0
   }
}
func (r *Line) GetAnsiCode(defaultColor Color.Color) string {
	if !r.visible {
		return ""
	}
	if r.isChanged {
		r.ansiCode = r.getAnsiLine(defaultColor)
		r.isChanged = false
	}
	return r.ansiCode
}

func (r *Line) SetAngle(angle int16) {
	switch angle {
	case 0, 45, 90:
		r.angle = angle
	default:
		return
	}
	r.Touch()
}

func (l *Line) SetPos(x, y int) {
	l.xPos = x
	l.yPos = y
	l.Touch()
	EventManager.Call(EventManager.ReorganizeElements, []any{l})
}

func (b *Line) SetLayer(layer Core.Layer) {
	if layer>Core.LMax{
		layer=Core.LMax
	}
	b.layer = layer
	b.Touch()
	EventManager.Call(EventManager.ReorganizeElements,[]any{b})
}
func (c *Line) GetLayer() Core.Layer {
	return c.layer
}
func (c *Line) SetColor(color Color.Color) {
	c.color = color
	c.Touch()
}

func (l *Line) SetLen(len int) error {
	if len < 0 {
		return errors.New("len < 0")
	}
	l.Len = len
	l.Touch()
	return nil
}

func (s *Line) SetVisibility(visible bool) {
	s.visible = visible
}

func (s *Line) GetVisibility() bool {
	return s.visible
}

func (l *Line) GetPos() (int, int) {
	return l.xPos, l.yPos
}

func (l *Line) getAnsiLine(defaultColor Color.Color) string {
	var str *strings.Builder = new(strings.Builder)
	str.WriteString(U.SaveCursor)
	str.WriteString(l.color.GetMixedColor(defaultColor).GetAnsiColor())
	switch l.angle {
	case 0:
		ansiLineAngle0(l, str)
	case 45:
		ansiLineAngle45(l, str)
	case 90:
		ansiLineAngle90(l, str)
	case 270:
		ansiLineAngle270(l, str)
	}
	str.WriteString(U.RestoreCursor)
	str.WriteString(Color.GetResetColor())
	return str.String()
}

func ansiLineAngle0(l *Line, str *strings.Builder) {
	str.WriteString(U.GetAnsiMoveTo(l.xPos, l.yPos))
	str.WriteString(strings.Repeat(U.HorizontalLine, l.Len))
}

func ansiLineAngle270(l *Line, str *strings.Builder) {
	for i := 1; i < l.Len; i++ {
		str.WriteString(U.GetAnsiMoveTo(l.xPos+i, l.yPos+i))
		str.WriteString(U.HorizontalLine)
	}
}

func ansiLineAngle90(l *Line, str *strings.Builder) {
	for i := 0; i < l.Len; i++ {
		str.WriteString(U.GetAnsiMoveTo(l.xPos, l.yPos+i))
		str.WriteString(U.VerticalLine)
	}
}

func ansiLineAngle45(l *Line, str *strings.Builder) {
	for i := 0; i < l.Len; i++ {
		str.WriteString(U.GetAnsiMoveTo(l.xPos+i, l.yPos-i))
		str.WriteString(U.HorizontalLine)
	}
}
