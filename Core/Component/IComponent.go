package Component

import (
	"github.com/Wordluc/GTUI/Core"
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
  ///return offset between actual character and line to x,y,(AnCharacter-x,AnLine-y)
  DiffCurrentToXY(x,y int) (int,int) 
	SetCurrentPosCursor(x, y int)(int,int)
}
