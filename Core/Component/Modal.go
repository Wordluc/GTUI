package Component

import (
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
)
const (
	modalL1 Core.Layer= Core.LMax+iota
	modalL2
	
)

func CreateModal(x, y, sizeX, sizeY int,content *Container) *Container {
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	null := CreateNullComponent(0, 0, sizeX, sizeY)
	contComp:= CreateContainer(0,0)
	contComp.AddComponent(null)
	content.SetLayer(modalL2)
	contComp.AddDrawing(rect)
	content.SetPos(1,1)
	contComp.AddComponent(content)
	contComp.SetLayer(modalL1)
	contComp.SetPos(x, y)
	return contComp
}
