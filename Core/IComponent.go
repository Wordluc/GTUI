package Core

type OnEvent func()
type IComponent interface {
	OnClick()
	OnRelease()
	OnHover()
	OnLeave()
	GetGraphics() IDrawing
	GetShape() (IInteractiveShape, error)
	SetPos(x,y int)
	GetPos() (int,int)
	GetSize() (int,int)
	GetLayer() Layer
	SetLayer(layer Layer)
}

type IWritableComponent interface {
	IsTyping() bool
	StartTyping()
	StopTyping()
	//return offset between actual character and line to x,y,(AnCharacter-x,AnLine-y)
	DiffCurrentToXY(x,y int) (int,int) 
	SetCurrentPosCursor(x, y int)(int,int)
}
type IInteractiveShape interface {
	IsOn(x, y int) bool
	GetShapes() []BaseInteractiveShape
}
type BaseInteractiveShape struct {
	XPos   int
	YPos   int
	Width  int
	Height int
}
func (b BaseInteractiveShape) GetPos() (int, int) {
	return b.XPos, b.YPos
}
func (b BaseInteractiveShape) GetSize() (int, int) {
	return b.Width, b.Height
}

func (b *BaseInteractiveShape) IsOn(x, y int) bool {
	return x >= b.XPos && x < b.XPos+b.Width && y >= b.YPos && y < b.YPos+b.Height
}

func (b *BaseInteractiveShape) GetShapes() []BaseInteractiveShape {
	return []BaseInteractiveShape{*b}
}

type ComplexInteractiveShape struct {
	shapes []IInteractiveShape
}

func (c *ComplexInteractiveShape) IsOn(x, y int) bool {
	for _, shape := range c.shapes {
		if shape.IsOn(x, y) {
			return true
		}
	}
	return false
}

func (c *ComplexInteractiveShape) GetShapes() []BaseInteractiveShape {
	var shapes []BaseInteractiveShape
	for _, shape := range c.shapes {
		shapes = append(shapes, shape.GetShapes()...)
	}
	return shapes
}

func (c *ComplexInteractiveShape) AddShape(shape IInteractiveShape) {
	c.shapes = append(c.shapes, shape)
}
