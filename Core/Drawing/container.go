package Drawing

import (
	"strings"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/EventManager"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

type Container struct {
	xPos     int
	yPos     int
	drawings []Core.IDrawing
	color    Color.Color
	visible  bool
	layer    Core.Layer
}

func CreateContainer(x, y int) *Container {
	return &Container{
		xPos:    x,
		yPos:    y,
		color:   Color.GetNoneColor(),
		visible: true,
	}
}

func (c *Container) GetComponents() []Core.IComponent {
	return []Core.IComponent{}
}

func (c *Container) GetGraphics() []Core.IDrawing {
	return c.drawings
}

func (c *Container) AddDrawings(eles ...Core.IDrawing) error {
	for _, ele := range eles {
		c.drawings = append(c.drawings, ele)
	}
	return nil
}

func (c *Container) AddContainer(containers ...Core.IContainer) error {
	for _, conp := range containers {
		c.drawings = append(c.drawings, conp.GetGraphics()...)
	}
	return nil
}

func (c *Container) Touch() {
	for _, child := range c.drawings {
		child.Touch()
	}
}
func (c *Container) IsTouched() bool {
	for _, child := range c.drawings {
		if child.IsTouched() {
			return true
		}
	}
	return false
}
func (c *Container) SetVisibility(visible bool) {
	for _, child := range c.drawings {
		child.SetVisibility(visible)
	}
	c.visible = visible
	c.Touch()
	EventManager.Call(EventManager.ForceRefresh)
}

func (c *Container) GetVisibility() bool {
	return c.visible
}

func (c *Container) getAnsiCode(defaultColor Color.Color) string {
	var str strings.Builder
	str.WriteString(c.color.GetAnsiColor())
	for _, child := range c.drawings {
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
	for _, child := range c.drawings {
		xChild, yChild := child.GetPos()
		child.SetPos(xChild+deltaX, yChild+deltaY)
	}
	c.Touch()
	EventManager.Call(EventManager.ReorganizeElements, c)
}

func (c *Container) GetPos() (int, int) {
	return c.xPos, c.yPos
}

func (b *Container) SetLayer(layer Core.Layer) error {
	if layer < 0 {
		panic("layer can't be negative")
	}
	diff := layer - b.layer
	for _, comp := range b.drawings {
		comp.SetLayer(comp.GetLayer() + diff)
	}
	b.layer = layer
	b.Touch()
	EventManager.Call(EventManager.ReorganizeElements, b)
	return nil
}

func (c *Container) GetLayer() Core.Layer {
	return c.layer
}
func (c *Container) SetColor(color Color.Color) {
	c.color = color
	c.Touch()
}
