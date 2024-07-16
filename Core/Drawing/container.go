package Drawing

import (
	"GTUI/Core"
	"strings"
)

type Container struct {
	name     string
	XPos     int
	YPos     int
	children []Core.IEntity
}

func CreateContainer(name string, x, y int) *Container {
	return &Container{
		name: name,
		XPos: x,
		YPos: y,
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
		str.WriteString(child.GetAnsiCode())
	}
	return str.String()
}
func (c *Container) GetName() string {
	return c.name
}
func (c *Container) SetPos(x, y int) {
	deltaX := x - c.XPos
	deltaY := y - c.YPos
	c.XPos = x
	c.YPos = y
	for _, child := range c.children {
		xChild,yChild:=child.GetPos()
		child.SetPos(xChild+deltaX, yChild+deltaY)
	}
}
func (c *Container) GetPos() (int, int) {
	return c.XPos, c.YPos
}

