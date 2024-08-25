package Core

import "GTUI/Core/Utils/Color"

type IEntity interface {
	Touch()
	GetAnsiCode(defaultColor Color.Color) string
	SetPos(x, y int)
	GetPos() (int, int)
	SetVisibility(visible bool)
	GetVisibility() bool
}
