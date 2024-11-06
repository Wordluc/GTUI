package Drawing

import (
	"strings"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

type Container struct {
	xPos     int
	yPos     int
	children []Core.IEntity
	color    Color.Color
	visible  bool
}

func CreateContainer( x, y int) *Container {
	return &Container{
		xPos:  x,
		yPos:  y,
		color: Color.GetNoneColor(),
		visible: true,
	}
}

func (c *Container) GetChildren() []Core.IEntity {
	return c.children
}
//DO NOT USE
func (c *Container) GetSize()(int,int) {
	panic("mustn't be called")
}

func (c *Container) AddChild(child Core.IEntity) error {//TODO: controllare se l'errare è gestito dai caller
	c.children = append(c.children, child)
	return nil
}
func (c *Container) Touch() {
	for _, child := range c.children {
		child.Touch()
	}
}
func (c *Container) IsTouched() (bool) {
	for _, child := range c.children {
		if child.IsTouched(){
			return true
		}
	}
	return false
}
func (c *Container) SetVisibility(visible bool) {
	c.visible = visible
}

func (c *Container) GetVisibility() bool {
	return c.visible
}

func (c *Container) GetAnsiCode(defaultColor Color.Color) string {
	if !c.visible {
		return ""
	}
	var str strings.Builder
	str.WriteString(c.color.GetMixedColor(defaultColor).GetAnsiColor())
	for _, child := range c.children {
		str.WriteString(child.GetAnsiCode(defaultColor))
		str.WriteString(c.color.GetMixedColor(defaultColor).GetAnsiColor())
	}
	return str.String()
}

func (c *Container) getAnsiCode(defaultColor Color.Color) string {
	var str strings.Builder
	str.WriteString(c.color.GetAnsiColor())
	for _, child := range c.children {
		str.WriteString(child.GetAnsiCode(defaultColor))
		str.WriteString(c.color.GetAnsiColor())
	}
		str.WriteString(Color.GetResetColor())
	return str.String()
}

func (c *Container) SetPos(x, y int) {
	deltaX := x - c.xPos
	deltaY := y - c.yPos
	c.xPos = x
	c.yPos = y
	for _, child := range c.children {
		xChild, yChild := child.GetPos()
		child.SetPos(xChild+deltaX, yChild+deltaY)
	}
	c.Touch()
}

func (c *Container) GetPos() (int, int) {

	return c.xPos, c.yPos
}

func (c *Container) SetColor(color Color.Color) {
	c.color = color
	c.Touch()
}
