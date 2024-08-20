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
	IsOn() bool
	/// check if cursor can move y absolute position, otherwise return offset
	CanMoveXCursor(x int) int
	/// check if cursor can move x absolute position, otherwise return offset
	CanMoveYCursor(y int) int
}
