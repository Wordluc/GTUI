package Component

import (
	"GTUI/Core"
)

type IComponent interface {
	OnClick()
	OnRelease()
	OnHover()
	OnLeave()
	GetGraphics() Core.IEntity
	getShape() (InteractiveShape, error)
}

type IWritableComponent interface {
	IsTyping() bool
	StartTyping()
	StopTyping()
	/// return offset between total character and line to x,y,(AnCharacter-x,AnLine-y)	
	DiffTotalToXY(x,y int) (int,int)
  ///return offset between actual character and line to x,y,(AnCharacter-x,AnLine-y)
  DiffCurrentToXY(x,y int) (int,int) 
	SetCurrentPosCursor(x, y int)(int,int)
}
