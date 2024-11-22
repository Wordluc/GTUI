package Core

type OnEvent func()
type IComponent interface {
	OnClick()
	OnRelease()
	OnHover()
	OnLeave()
	GetGraphics() IDrawing
	SetPos(x, y int)
	GetPos() (int, int)
	GetSize() (int, int)
	GetLayer() Layer
	SetLayer(layer Layer)error
}

type IComposableComponent interface {
	GetComponets() []IComponent
}

type IWritableComponent interface {
	IsTyping() bool
	StartTyping()
	StopTyping()
	//return offset between actual character and line to x,y,(AnCharacter-x,AnLine-y)
	DiffCurrentToXY(x, y int) (int, int)
	SetCurrentPosCursor(x, y int) (int, int)
}
