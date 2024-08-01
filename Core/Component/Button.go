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
		b.interactiveArea.SetColor(Color.GetDefaultColor())
		b.isClicked = false
	} else {
		if b.onClick != nil {
			b.onClick()
		}
		b.interactiveArea.SetColor(Color.Get(Color.Blue, Color.None))
		b.isClicked = true
	}
}
func (b *Button) OnRelease() {
	if b.onRelease != nil {
		b.onRelease()
	}
}
func (b *Button) OnHover() {
	if b.onHover != nil {
		b.onHover()
	}
}
func (b *Button) OnLeave() {
	if b.onLeave != nil {
		b.OnLeave()
	}
}
func (b *Button) isOn(xCur, yCur int) bool {
	x, y := b.interactiveArea.GetPos()
	return xCur >= x && xCur < x+b.interactiveArea.Width && yCur >= y && yCur < y+b.interactiveArea.Height
}
func (b *Button) GetGraphics() Core.IEntity {
	return b.graphics
}
func (b *Button) insertingToMap(s *ComponentM) error {
	if s == nil {
		return errors.New("component manager is nil")
	}
	x, y := b.interactiveArea.GetPos()
  finalX,finalY:=x+b.interactiveArea.Width,y+b.interactiveArea.Height
	for i := x; i <= finalX; i+=s.ChunkSize {
		for j := y; j <= finalY; j+=s.ChunkSize {
			xC, yC := i/s.ChunkSize, j/s.ChunkSize
			ele := (*s.Map)[xC][yC]
			if ele == nil {
				(*s.Map)[xC][yC] = &[]IComponent{b}
			} else {
				(*ele) = append(*ele, b)
			}
		}
	}
	return nil
}
