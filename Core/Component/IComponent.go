package Component

import (
	"github.com/Wordluc/GTUI/Core"
)

type IComponent interface {
	OnClick(int,int)
	OnRelease(int,int)
	OnHover(int,int)
	OnOut(int,int)
	GetGraphics() Core.IEntity
	getShape() (IInteractiveShape, error)
	SetPos(x,y int)
	GetPos() (int,int)
	GetSize() (int,int)
}

type IWritableComponent interface {
	IsTyping() bool
	StartTyping()
	StopTyping()
	//return offset between actual character and line to x,y,(AnCharacter-x,AnLine-y)
	DiffCurrentToXY(x,y int) (int,int) 
	SetCurrentPosCursor(x, y int)(int,int)
}
