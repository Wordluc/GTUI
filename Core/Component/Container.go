package Component

import (
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
)

type Container struct {
	onClick    func()
	active     bool
	components []IComponent
	drawing    *Drawing.Container
}

func CreateContainer(x, y int) *Container {
	return &Container{
		active:     true,
		drawing:    Drawing.CreateContainer(x, y),
		components: make([]IComponent, 0),
	}
}
func (c *Container) AddComponent(component IComponent)error {
	c.components = append(c.components, component)
	return c.drawing.AddChild(component.GetGraphics())
}

func (c *Container) AddDrawing(container Drawing.Container)error {
	return c.drawing.AddChild(&container)
}

func (c *Container) GetSize() (int,int) {
	return c.drawing.GetSize()
}

func (c *Container) GetComponents() []IComponent {
	return c.components
}
func (c *Container) SetonClick(onClick func()) {
	c.onClick = onClick
}

func (c *Container) SetActivity(active bool) {
	c.active = active
}
func (c *Container) GetActivity() bool {
	return c.active
}

func (c *Container) SetPos(x, y int) {
	c.drawing.SetPos(x, y)
}

func (c *Container) GetPos() (int, int) {
	return c.drawing.GetPos()
}

func (c *Container) OnClick(x, y int) {
	if !c.active {
		return
	}
	if c.onClick != nil {
		c.onClick()
	}
	for _, component := range c.components {
		shape, _ := component.getShape()
		if shape.isOn(x, y) {
			component.OnClick(x, y)
		}
	}
}

func (c *Container) OnRelease(x, y int) {
	if !c.active {
		return
	}
	for _, component := range c.components {
		shape, _ := component.getShape()
		if shape.isOn(x, y) {
			component.OnRelease(x, y)
		}
	}
}

func (c *Container) OnHover(x, y int) {
	if !c.active {
		return
	}
	for _, component := range c.components {
		shape, _ := component.getShape()
		if shape.isOn(x, y) {
			component.OnHover(x, y)
		}
	}
}

func (c *Container) OnOut(x, y int) {
	if !c.active {
		return
	}
	for _, component := range c.components {
		shape, _ := component.getShape()
		if !shape.isOn(x, y) {
			component.OnOut(x, y)
		}
	}
}

func (c *Container) GetGraphics() Core.IEntity {
	return c.drawing
}

func (c *Container) getShape() (IInteractiveShape, error) {
	complexShape := &ComplexInteractiveShape{}
	for _, component := range c.components {
		shape, err := component.getShape()
		if err != nil {
			return &BaseInteractiveShape{}, err
		}
		complexShape.addShape(shape)
	}
	return complexShape, nil
}
