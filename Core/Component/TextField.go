package Component

import "GTUI/Core/Drawing"

type TextField struct {
	graphics *Drawing.Container
	interactiveArea *Drawing.Rectangle
	isTyping bool
}

func CreateTextField(x, y, sizeX, sizeY int) *TextField {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	rectIn := Drawing.CreateRectangle(1, 1, sizeX-2, sizeY-2)
	cont.AddChild(rect)
	cont.AddChild(rectIn)
	cont.SetPos(x, y)
	return &TextField{
		graphics:        cont,
		interactiveArea: rectIn,
		isTyping:        false,
	}
}
