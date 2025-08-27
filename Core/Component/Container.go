package Component

import (
	"errors"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/EventManager"
)

type Container struct {
	onClick    func()
	active     bool
	visible    bool
	components []Core.IComponent
	drawings   []Core.IDrawing
	layer      Core.Layer
	x, y       int
}

func CreateContainer(x, y int) *Container {
	return &Container{
		active:     true,
		x:          x,
		y:          y,
		drawings:   make([]Core.IDrawing, 0),
		components: make([]Core.IComponent, 0),
	}
}
func (c *Container) AddComponent(components ...Core.IComponent) error {
	c.components = append(c.components, components...)
	return nil
}

func (c *Container) AddDrawing(drawings ...Core.IDrawing) error {
	c.drawings = append(c.drawings, drawings...)
	return nil
}

func (c *Container) AddContainer(containers ...Core.IContainer) error {
	for _, conp := range containers {
		c.drawings = append(c.drawings, conp.GetGraphics()...)
		c.components = append(c.components, conp.GetComponents()...)
	}
	return nil
}

func (c *Container) GetComponents() []Core.IComponent {
	return c.components
}

func (c *Container) GetGraphics() (result []Core.IDrawing) {
	result = append(result, c.drawings...)

	for _, drawing := range c.components {
		result = append(result, drawing.GetGraphics()...)
	}
	return result
}

func (b *Container) SetLayer(layer Core.Layer) error {
	if layer < 0 {
		return errors.New("layer can't be negative")
	}
	for _, comp := range b.components {
		comp.SetLayer(comp.GetLayer() + layer)
	}
	for _, draw := range b.drawings {
		draw.SetLayer(draw.GetLayer() + layer)
	}
	b.layer = layer
	EventManager.Call(EventManager.ReorganizeElements, b)
	return nil
}
func (c *Container) GetLayer() Core.Layer {
	return c.layer
}

func (c *Container) GetVisibility() bool {
	return c.visible
}

func (c *Container) SetVisibility(visible bool) {
	for _, comp := range c.drawings {
		comp.SetVisibility(visible)
	}
	for _, comp := range c.components {
		for _, drawing := range comp.GetGraphics() {
			drawing.SetVisibility(visible)
		}
	}
	c.visible = visible
}

func (c *Container) SetonClick(onClick func()) {
	c.onClick = onClick
}

func (c *Container) SetActive(active bool) {
	c.active = active
	for _, comp := range c.components {
		comp.SetActive(active)
	}
}

func (c *Container) GetActivity() bool {
	return c.active
}

func (c *Container) SetPos(x, y int) {
	diffX := x - c.x
	diffY := y - c.y
	for _, comp := range c.components {
		xE, yE := comp.GetPos()
		comp.SetPos(xE+diffX, yE+diffY)
	}
	for _, draw := range c.drawings {
		cD, cE := draw.GetPos()
		draw.SetPos(cD+diffX, cE+diffY)
		draw.Touch()
	}
	c.x = x
	c.y = y
	EventManager.Call(EventManager.ReorganizeElements, c)
}

func (c *Container) GetPos() (int, int) {
	return c.x, c.y
}
