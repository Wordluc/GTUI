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
	getShape() (InteractiveShape,error)
}
type Event int8
const (
	OnClick Event= iota
	OnRelease
	OnHover
	OnLeave
)
