package Core

type OnEvent func()
type IComponent interface {
	OnClick()
	OnRelease()
	OnHover()
	OnLeave()
	GetGraphics() []IDrawing
	SetPos(x, y int)
	GetPos() (int, int)
	GetSize() (int, int)
	GetLayer() Layer
	SetLayer(layer Layer) error
}

type IContainer interface {
	GetDrawings() []IDrawing
	GetComponents() []IComponent
	AddContainer(...IContainer) error
}
type IComplexElement interface {
	GetDrawings() []IDrawing
	GetComponents() []IComponent
}
type IWritableComponent interface {
	IsTyping() bool
	StartTyping()
	StopTyping()
	//return offset between actual character and line to x,y,(AnCharacter-x,AnLine-y)
	DiffCurrentToXY(x, y int) (int, int)
	//Set the cursor inside the component to the position (x,y)
	SetCurrentPosCursor(x, y int) (int, int)
}
