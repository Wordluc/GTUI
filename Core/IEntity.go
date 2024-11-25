package Core

import "github.com/Wordluc/GTUI/Core/Utils/Color"

type IDrawing interface {
	//Notify that the element has been touched,so it needs to be redrawn
	Touch()
	IsTouched() bool
	GetAnsiCode(defaultColor Color.Color) string
	SetPos(x, y int)
	GetPos() (int, int)
	SetVisibility(visible bool)
	GetVisibility() bool
	GetSize() (int, int)
	GetLayer() Layer
	SetLayer(layer Layer)error
}

type IEntity interface {
	GetPos() (int, int)
	GetSize() (int, int)
	GetLayer() Layer
}
