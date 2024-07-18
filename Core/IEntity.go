package Core

import "GTUI/Core/Color"

type IEntity interface {
	Touch()
	GetAnsiCode(defaultColor Color.Color) string
	SetPos(x, y int)
	GetPos() (int, int)
}
