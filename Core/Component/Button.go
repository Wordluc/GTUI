package Component

import (
	"GTUI/Core"
	"GTUI/Core/Drawing"
	"GTUI/Core/Utils/Color"
)

type OnEvent func()
type Button struct {
	graphics        *Drawing.Container
	visibleArea     *Drawing.Rectangle
	onClick         OnEvent
	onRelease       OnEvent
	onHover         OnEvent
	onLeave         OnEvent
	isClicked       bool
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
		graphics:        cont,
		visibleArea:     rect,
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
	if !b.isClicked {
		b.isClicked = true
		if b.onClick != nil {
			b.onClick()
		}
		b.updateColorByClick()
	}
}
func (b *Button) OnRelease() {
	if b.onRelease != nil {
		b.onRelease()
	}
	b.isClicked=false
	b.visibleArea.SetColor(Color.Get(Color.Gray, Color.None))
}
func (b *Button) OnHover() {
	if !b.isClicked {
		b.visibleArea.SetColor(Color.Get(Color.Gray, Color.None))
	}
	if b.onHover != nil {
		b.onHover()
	}
}
func (b *Button) OnLeave() {
	if !b.isClicked {
		b.visibleArea.SetColor(Color.GetDefaultColor())
	}
	if b.onLeave != nil {
		b.OnLeave()
	}
}
func (b *Button) updateColorByClick() {
   if b.isClicked{
			b.visibleArea.SetColor(Color.Get(Color.Blue, Color.None))
	 }else{
			b.visibleArea.SetColor(Color.Get(Color.Gray, Color.None))
	 }
}
func (b *Button) GetGraphics() Core.IEntity {
	return b.graphics
}
func (b *Button) getShape() (InteractiveShape,error) {
	x,y:=b.visibleArea.GetPos()
	xDim,yDim:=b.visibleArea.GetSize()
	shape:=InteractiveShape{
      xPos:x+1,
		yPos:y+1,
		Width:xDim-1,
		Height:yDim-1,
	}
	return shape,nil
}
