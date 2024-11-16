package Drawing

import (
	"errors"
	"strings"

	"github.com/Wordluc/GTUI/Core/EventManager"
	U "github.com/Wordluc/GTUI/Core/Utils"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

type Rectangle struct {
	isChanged bool
	ansiCode  string
	color     Color.Color
	xPos      int
	yPos      int
	width     int
	height    int
	visible   bool
}

func CreateRectangle(x, y, width, height int) *Rectangle {
	return &Rectangle{
		ansiCode:  "",
		xPos:      x,
		yPos:      y,
		width:     width,
		isChanged: true,
		height:    height,
		color:     Color.GetNoneColor(),
		visible:   true,
	}
}

func (r *Rectangle) Touch() {
	r.isChanged = true
}
func (r *Rectangle) IsTouched() bool {
	return r.isChanged
}
func (r *Rectangle) GetAnsiCode(defaultColor Color.Color) string {
	if !r.visible {
		return ""
	}
	if r.isChanged {
		r.ansiCode = r.getAnsiRectangle(defaultColor)
		r.isChanged = false
	}
	return r.ansiCode
}

func (l *Rectangle) SetPos(x, y int) {
	l.xPos = x
	l.yPos = y
	l.Touch()
	EventManager.Call(EventManager.ReorganizeElements,[]any{l})
}

func (l *Rectangle) GetSize() (int, int) {
	return l.width, l.height
}

func (l *Rectangle) SetSize(x, y int) error {
	if x < 0 || y < 0 {
		return errors.New("invalid size")
	}
	l.width = x
	l.height = y
	l.Touch()
	return nil
}

func (c *Rectangle) SetColor(color Color.Color) {
	c.color = color
	c.Touch()
}

func (l *Rectangle) GetPos() (int, int) {
	return l.xPos, l.yPos
}

func (s *Rectangle) SetVisibility(visible bool) {
	s.visible = visible
}

func (l *Rectangle) GetVisibility() bool {
	return l.visible
}
 
func (s *Rectangle) getAnsiRectangle(defaultColor Color.Color) string {
	var horizontal = strings.Repeat(U.HorizontalLine, s.width-2)
	var color = s.color.GetMixedColor(defaultColor).GetAnsiColor()
	var builStr strings.Builder
	builStr.WriteString(color)
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
