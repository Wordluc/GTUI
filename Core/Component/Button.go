package Component

import (
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

type OnEvent func()
type Button struct {
	graphics    *Drawing.Container
	visibleArea *Drawing.Rectangle
	onClick     OnEvent
	onRelease   OnEvent
	onHover     OnEvent
	onLeave     OnEvent
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
	return b.graphics.GetPos()
}
func (b *Button) GetSize() (int, int) {
	return b.graphics.GetSize()
}
func (b *Button) SetOnClick(onClick OnEvent) {
	b.onClick = onClick
}
func (b *Button) SetOnLeave(onLeave OnEvent) {
	b.onLeave = onLeave
}
func (b *Button) SetOnRelease(onRelease OnEvent) {
	b.onRelease = onRelease
}
func (b *Button) SetOnHover(onHover OnEvent) {
	b.onHover = onHover
}
func (b *Button) OnClick(_,_ int) {
	if !b.isClicked {
		if b.onClick != nil {
			b.onClick() 
		}
		b.isClicked = true
		b.updateColorByClick()
	}
}
func (b *Button) OnRelease(_,_ int) {
	if b.onRelease != nil {
		b.onRelease()
	}
	b.isClicked = false
	b.visibleArea.SetColor(Color.Get(Color.Gray, Color.None))
}
func (b *Button) OnHover(_,_ int) {
	if !b.isClicked {
		b.visibleArea.SetColor(Color.Get(Color.Gray, Color.None))
	}
	if b.onHover != nil {
		b.onHover()
	}
}
func (b *Button) OnOut(_,_ int) {
	if !b.isClicked {
		b.visibleArea.SetColor(Color.Get(Color.Gray, Color.None))
	}
	if b.onLeave != nil {
		b.onLeave()
	}
}
func (b *Button) updateColorByClick() {
	if b.isClicked {
		b.visibleArea.SetColor(Color.Get(Color.Blue, Color.None))
	} else {
		b.visibleArea.SetColor(Color.Get(Color.Gray, Color.None))
	}
}
func (b *Button) GetGraphics() Core.IEntity {
	return b.graphics
}
func (b *Button) GetVisibleArea() *Drawing.Rectangle {
	return b.visibleArea
}
func (b *Button) getShape() (IInteractiveShape, error) {
	x, y := b.visibleArea.GetPos()
	xDim, yDim := b.visibleArea.GetSize()
	shape := BaseInteractiveShape{
		xPos:   x + 1,
		yPos:   y + 1,
		Width:  xDim - 1,
		Height: yDim - 1,
	}
	return &shape, nil
}
