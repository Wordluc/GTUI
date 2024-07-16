package Drawing

import (
	"GTUI/Core"
	"GTUI/Core/Color"
	"strings"
)

type Container struct {
	name     string
	xPos     int
	yPos     int
	children []Core.IEntity
	color    Color.Color
}

func CreateContainer(name string, x, y int) *Container {
	return &Container{
		name: name,
		xPos: x,
		yPos: y,
	}
}
func (c *Container) GetChildren() []Core.IEntity {
	return c.children
}
func (c *Container) AddChild(child Core.IEntity) {
	c.children = append(c.children, child)
}
func (c *Container) Touch() {
	for _, child := range c.children {
		child.Touch()
	}
}
func (c *Container) GetAnsiCode() string {
	var str strings.Builder
	for _, child := range c.children {
		str.WriteString(c.color.GetAnsiColor())
		str.WriteString(child.GetAnsiCode())
	}
	return str.String()
}
func (c *Container) GetName() string {
	return c.name
}
func (c *Container) SetPos(x, y int) {
	deltaX := x - c.xPos
	deltaY := y - c.yPos
	c.xPos = x
	c.yPos = y
	for _, child := range c.children {
		xChild,yChild:=child.GetPos()
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
