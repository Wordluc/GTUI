package Component

type StreamCharacter struct {
	Get      func() chan string
	IChannel int
	Delete   func()
}
type IInteractiveShape interface {
	isOn(x, y int) bool
	getShapes() []BaseInteractiveShape
}
type BaseInteractiveShape struct {
	xPos   int
	yPos   int
	Width  int
	Height int
}
func (b BaseInteractiveShape) GetPos() (int, int) {
	return b.xPos, b.yPos
}
func (b BaseInteractiveShape) GetSize() (int, int) {
	return b.Width, b.Height
}

func (b *BaseInteractiveShape) isOn(x, y int) bool {
	return x >= b.xPos && x < b.xPos+b.Width && y >= b.yPos && y < b.yPos+b.Height
}

func (b *BaseInteractiveShape) getShapes() []BaseInteractiveShape {
	return []BaseInteractiveShape{*b}
}

type ComplexInteractiveShape struct {
	shapes []IInteractiveShape
}

func (c *ComplexInteractiveShape) isOn(x, y int) bool {
	for _, shape := range c.shapes {
		if shape.isOn(x, y) {
			return true
		}
	}
	return false
}

func (c *ComplexInteractiveShape) getShapes() []BaseInteractiveShape {
	var shapes []BaseInteractiveShape
	for _, shape := range c.shapes {
		shapes = append(shapes, shape.getShapes()...)
	}
	return shapes
}

func (c *ComplexInteractiveShape) addShape(shape IInteractiveShape) {
	c.shapes = append(c.shapes, shape)
}
