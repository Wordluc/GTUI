package Core

import "github.com/Wordluc/GTUI/Core/Utils/Color"

type Layer int8
const (
	//it is on the bottom
	L1 Layer = iota
	L2
	L3
	L4
	//it is on the top
	L5
)
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
