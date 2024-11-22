package Component

import (
	"errors"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/EventManager"
)

type Container struct {
	onClick    func()
	active     bool
	components []Core.IComponent
	drawing    *Drawing.Container
	layer      Core.Layer
	x,y int
}

func CreateContainer(x, y int) *Container {
	return &Container{
		active:     true,
		x: x,
		y: y,
		drawing:    Drawing.CreateContainer(x, y),
		components: make([]Core.IComponent, 0),
	}
}
func (c *Container) AddComponent(component Core.IComponent)error {
	c.components = append(c.components, component)
	return nil
}

func (c *Container) AddDrawing(container Core.IDrawing)error {
	return c.drawing.AddChild(container)
}

func (b *Container) SetLayer(layer Core.Layer) error{
	if layer<0 {
		return errors.New("layer can't be negative")
	}
	diff:= layer - b.layer
	for _,comp:=range b.components {
		comp.SetLayer(comp.GetLayer()+diff)
	}
	b.drawing.SetLayer(b.drawing.GetLayer()+diff)
	b.layer = layer
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
	return nil
}
func (c *Container) GetLayer() Core.Layer {
	return c.layer
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
	diffX:=x-c.x
	diffY:=y-c.y
	for _, comp := range c.components {
		xE, yE := comp.GetPos()
		comp.SetPos(xE+diffX, yE+diffY)
	}
	cD, cE := c.drawing.GetPos()
	c.drawing.SetPos(cD+diffX, cE+diffY)
	c.x=x
	c.y=y
	c.drawing.Touch()
	EventManager.Call(EventManager.ReorganizeElements, []any{c})
}

func (c *Container) GetPos() (int, int) {
	return c.drawing.GetPos()
}

func (c *Container) GetGraphics() Core.IDrawing {
	return c.drawing
}

//DO NOT USE
func (c *Container) OnClick() {
	panic("mustn't be called")
}

//DO NOT USE
func (c *Container) OnRelease() {
	panic("mustn't be called")
}

//DO NOT USE
func (c *Container) OnHover() {
	panic("mustn't be called")
}

//DO NOT USE
func (c *Container) OnLeave() {
	panic("mustn't be called")
}

//DO NOT USE
func (c *Container) GetSize() (int,int) {
	panic("mustn't be called")
}
