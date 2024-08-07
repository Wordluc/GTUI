package Component

import (
	"GTUI/Core"
	"GTUI/Core/Drawing"
	"GTUI/Core/Utils/Color"
	"errors"
)

type OnEvent func()
type Button struct {
	graphics        *Drawing.Container
	interactiveArea *Drawing.Rectangle
	onClick         OnEvent
	onRelease       OnEvent
	onHover         OnEvent
	onLeave         OnEvent
	isClicked       bool
}

func CreateButton(x, y, sizeX, sizeY int, text string) *Button {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	textD := Drawing.CreateTextBox(0, 0)
	textD.Type(text)
	xC, yC := sizeX/2-len(text)/2, sizeY/2
	textD.SetPos(xC, yC)
	cont.AddChild(rect)
	cont.AddChild(textD)
	cont.SetPos(x, y)
	return &Button{
		graphics:        cont,
		interactiveArea: rect,
		isClicked:       false,
	}
}

func (b *Button) SetOnClick(onClick OnEvent) {
	b.onClick = onClick
}
func (b *Button) SetOnRelease(onRelease OnEvent) {
	b.onRelease = onRelease
}
func (b *Button) SetOnHover(onHover OnEvent) {
	b.onHover = onHover
}

func (b *Button) OnClick() {
	if b.isClicked {
		b.isClicked = false
	} else {
		if b.onClick != nil {
			b.onClick()
		}
		b.isClicked = true
	}
	b.updateColorByClick()
}
func (b *Button) OnRelease() {
	if b.onRelease != nil {
		b.onRelease()
	}
}
func (b *Button) OnHover() {
	if !b.isClicked {
		b.interactiveArea.SetColor(Color.Get(Color.Gray, Color.None))
	}
	if b.onHover != nil {
		b.onHover()
	}
}
func (b *Button) OnLeave() {
	if !b.isClicked {
		b.interactiveArea.SetColor(Color.GetDefaultColor())
	}
	if b.onLeave != nil {
		b.OnLeave()
	}
}
func (b *Button) updateColorByClick() {
   if b.isClicked{
			b.interactiveArea.SetColor(Color.Get(Color.Blue, Color.None))
	 }else{
			b.interactiveArea.SetColor(Color.Get(Color.Gray, Color.None))
	 }
}
func (b *Button) GetGraphics() Core.IEntity {
	return b.graphics
}
func (b *Button) getShape() (InteractiveShape,error) {
	if b.interactiveArea == nil {
		return InteractiveShape{}, errors.New("interactive area is nil")
	}
	shape:=InteractiveShape{}
	shape.xPos,shape.yPos= b.interactiveArea.GetPos()
	shape.Width,shape.Height= b.interactiveArea.GetSize()
	return shape,nil
}
