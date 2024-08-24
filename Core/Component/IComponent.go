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
	/// return offset between last character position and x, nLine-y
	CanMoveXCursor(x int) int
	/// return offset between numer of line and y, nCharacter-x
	CanMoveYCursor(y int) int

	SetCurPos(x, y int)
   ///return offset between actual character and line to x,y,(AnCharacter-x,AnLine-y)
   DiffActualToTotal(x,y int) (int,int) 
}
