package Core

import "GTUI/Core/Color"

type IEntity interface {
	Touch()
	GetAnsiCode() string
	GetName() string
	SetPos(x, y int)
	GetPos() (int, int)
	SetColor(color Color.Color)
}
