package Component

import (
	"GTUI/Core/Drawing"
	"errors"
)

type OnEvent func()
type Button struct {
	graphics  *Drawing.Container
	interactiveArea      *Drawing.Rectangle
	onClick   OnEvent
	onRelease OnEvent
	onHover   OnEvent
	onLeave   OnEvent
}

func CreateButton(x, y, sizeX, sizeY int, text string) *Button {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	textD := Drawing.CreateTextBox(0, 0)
	textD.Type(text)
	cont.AddChild(rect)
	cont.AddChild(textD)
	return &Button{
		graphics: cont,
		interactiveArea:     rect,
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
	if b.onClick != nil {
		b.onClick()
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
func (b *Button)isOn(xCur,yCur int)bool{
	x,y:=b.interactiveArea.GetPos()
	return xCur>=x && xCur<=x+b.interactiveArea.Width && yCur>=y && yCur<=y+b.interactiveArea.Height
}
func (b *Button) insertingToMap(s *Search) error {
	for _, e := range b.graphics.GetChildren() {
		x, y := e.GetPos()
		xiChunk := x / s.ChunkSize
		yiChunk := y / s.ChunkSize
		ele := (*s.Map)[xiChunk][yiChunk]
		if ele != nil {
			return errors.New("error adding component")
		}

		ele = &[]IComponent{b}
		(*ele) = append(*ele, b)
	}
	return nil
}
