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

type ICursorInteragibleComponent interface {
	IsWritable() bool
	/// return offset between total character and line to x,y,(AnCharacter-x,AnLine-y)	
	DiffTotalToXY(x,y int) (int,int)
	SetCurrentPos(x, y int)
  ///return offset between actual character and line to x,y,(AnCharacter-x,AnLine-y)
  DiffCurrentToXY(x,y int) (int,int) 
}
