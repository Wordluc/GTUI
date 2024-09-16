package Component

import (
	"GTUI/Core"
	"GTUI/Core/Drawing"
	"GTUI/Core/Utils"
	"errors"
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

func CreateButton(x, y, sizeX, sizeY int, text string) (*Button, error) {
	totalError := Utils.NewError()
	totalError.AddIf(sizeX <= 0 || sizeY <= 0, errors.New("invalid size"))
	cont, e := Drawing.CreateContainer(0, 0)
	totalError.Add(e) 
	rect, e := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	totalError.Add(e) 
	textD,e := Drawing.CreateTextField(0, 0)
	totalError.Add(e)
	if totalError.HasError() {
		return nil, totalError
	}
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
	}, nil
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
func (b *Button) SetOnLeave(onLeave OnEvent) {
	b.onLeave = onLeave
}
func (b *Button) GetVisibleArea() *Drawing.Rectangle {
	return b.visibleArea
}
func (b *Button) OnClick() {
	if !b.isClicked {
		b.isClicked = true
		if b.onClick != nil {
			b.onClick()
		}
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
func (b *Button) getShape() (InteractiveShape, error) {
	x, y := b.visibleArea.GetPos()
	xDim, yDim := b.visibleArea.GetSize()
	shape := InteractiveShape{
		xPos:   x + 1,
		yPos:   y + 1,
		Width:  xDim - 1,
		Height: yDim - 1,
	}
	return shape, nil
}
