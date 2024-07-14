package Core

import "GTUI/Core/Color"

type IEntity interface {
	Touch()
	GetAnsiCode() string
	GetName() string
	SetColor(color Color.Color)
	GetColor() Color.Color
}
