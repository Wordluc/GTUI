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
	SetActive(bool)
	GetActive() bool
}

type IContainer interface {
	GetGraphics() []IDrawing
	GetComponents() []IComponent
	AddContainer(...IContainer) error
}
type IComplexElement interface {
	GetGraphics() []IDrawing
	GetComponents() []IComponent
	AddContainer(...IContainer) error
	AddComponent(...IComponent) error
	AddDrawing(...IDrawing) error
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
