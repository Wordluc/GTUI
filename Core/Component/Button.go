package Component

import (
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
)

type Button struct {
	graphics    *Drawing.Container
	visibleArea *Drawing.Rectangle
	onClick     Core.OnEvent
	onRelease   Core.OnEvent
	onHover     Core.OnEvent
	onLeave     Core.OnEvent
	isClicked   bool
}

func CreateButton(x, y, sizeX, sizeY int, text string) *Button {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	textD := Drawing.CreateTextField(0, 0)
	textD.Type(text)
	xC, yC := sizeX/2-len(text)/2, sizeY/2
	textD.SetPos(xC, yC)
	cont.AddChild(rect)
	cont.AddChild(textD)
	cont.SetPos(x, y)
	return &Button{
		graphics:    cont,
		visibleArea: rect,
		isClicked:   false,
	}
}
func (b *Button) SetPos(x, y int) {
	b.graphics.SetPos(x, y)
}
func (b *Button) GetPos() (int, int) {
	return b.visibleArea.GetPos()
}
func (b *Button) GetSize() (int, int) {
	return b.visibleArea.GetSize()
}
func (b *Button) SetOnClick(onClick Core.OnEvent) {
	b.onClick = onClick
}
func (b *Button) SetOnLeave(onLeave Core.OnEvent) {
	b.onLeave = onLeave
}
func (b *Button) SetOnRelease(onRelease Core.OnEvent) {
	b.onRelease = onRelease
}
func (b *Button) SetOnHover(onHover Core.OnEvent) {
	b.onHover = onHover
}
func (b *Button) OnClick() {
	if !b.isClicked {
		if b.onClick != nil {
			b.onClick() 
		}
		b.isClicked = true
	}
}
func (b *Button) OnRelease() {
	if b.onRelease != nil {
		b.onRelease()
	}
	b.isClicked = false
}
func (b *Button) OnHover() {
	if b.onHover != nil {
		b.onHover()
	}
}
func (b *Button) OnLeave() {
	if b.onLeave != nil {
		b.onLeave()
	}
}
func (b *Button) GetGraphics() Core.IEntity {
	return b.graphics
}
func (b *Button) GetVisibleArea() *Drawing.Rectangle {
	return b.visibleArea
}
func (b *Button) GetShape() (Core.IInteractiveShape, error) {
	x, y := b.visibleArea.GetPos()
	xDim, yDim := b.visibleArea.GetSize()
	shape := Core.BaseInteractiveShape{
		XPos:   x + 1,
		YPos:   y + 1,
		Width:  xDim - 1,
		Height: yDim - 1,
	}
	return &shape, nil
}
