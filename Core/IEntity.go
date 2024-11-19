package Core

import "github.com/Wordluc/GTUI/Core/Utils/Color"

type IDrawing interface {
	Touch()
	IsTouched() bool
	GetAnsiCode(defaultColor Color.Color) string
	SetPos(x, y int)
	GetPos() (int, int)
	SetVisibility(visible bool)
	GetVisibility() bool
	GetSize() (int, int)
	GetLayer() Layer
	SetLayer(layer Layer)
}

type IEntity interface {
	GetPos() (int, int)
	GetSize() (int, int)
	GetLayer() Layer
	SetLayer(layer Layer)
}
