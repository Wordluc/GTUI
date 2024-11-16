package Component

import (
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
)

type Container struct {
	onClick    func()
	active     bool
	components []Core.IComponent
	drawing    *Drawing.Container
}

func CreateContainer(x, y int) *Container {
	return &Container{
		active:     true,
		drawing:    Drawing.CreateContainer(x, y),
		components: make([]Core.IComponent, 0),
	}
}
func (c *Container) AddComponent(component Core.IComponent)error {
	c.components = append(c.components, component)
	return c.drawing.AddChild(component.GetGraphics())
}

func (c *Container) AddDrawing(container *Drawing.Container)error {
	return c.drawing.AddChild(container)
}
//DO NOT USE
func (c *Container) GetSize() (int,int) {
	panic("mustn't be called")
}

func (c *Container) GetComponents() []Core.IComponent {
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

func (c *Container) OnClick() {
	panic("mustn't be called")
}

func (c *Container) OnRelease() {
	panic("mustn't be called")
}

func (c *Container) OnHover() {
	panic("mustn't be called")
}

func (c *Container) OnLeave() {
	panic("mustn't be called")
}

func (c *Container) GetGraphics() Core.IDrawing {
	return c.drawing
}

func (c *Container) GetShape() (Core.IInteractiveShape, error) {
	complexShape := &Core.ComplexInteractiveShape{}
	for _, component := range c.components {
		shape, err := component.GetShape()
		if err != nil {
			return &Core.BaseInteractiveShape{}, err
		}
		complexShape.AddShape(shape)
	}
	return complexShape, nil
}
