package Component

import (
	"GTUI/Core"
	"GTUI/Core/Drawing"
	"GTUI/Core/Utils/Color"
	"errors"
)

type TextField struct {
	graphics        *Drawing.Container
	interactiveArea *Drawing.Rectangle
	textBox         *Drawing.TextBox
	isTyping        bool
	streamText      StreamCharacter
}

func CreateTextField(x, y, sizeX, sizeY int, streamText StreamCharacter) *TextField {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	textBox := Drawing.CreateTextBox(1, 1)
	cont.AddChild(rect)
	cont.AddChild(textBox)
	cont.SetPos(x, y)
	return &TextField{
		graphics:        cont,
		interactiveArea: rect,
		isTyping:        false,
		streamText:      streamText,
		textBox:         textBox,
	}
}
func (b *TextField) loopTyping() {
	channel := b.streamText.Get()
	for str := range channel {
		if !b.isTyping {
			return
		}
		b.textBox.Type(string(str))
	}
}

func (b *TextField) OnClick() {
	if b.isTyping {
		return
	}
	b.isTyping = true
	b.interactiveArea.SetColor(Color.Get(Color.Green, Color.None))
	go b.loopTyping()
}

func (b *TextField) OnLeave() {
	b.isTyping = false
	b.interactiveArea.SetColor(Color.GetDefaultColor())
	b.streamText.Delete()
}

func (b *TextField) OnRelease() {}

func (b *TextField) OnHover() {
	b.interactiveArea.SetColor(Color.Get(Color.Gray, Color.None))
}

func (b *TextField) GetGraphics() Core.IEntity {
	return b.graphics
}

func (b *TextField) getShape() (InteractiveShape, error) {
	if b.interactiveArea == nil {
		return InteractiveShape{}, errors.New("interactive area is nil")
	}
	shape := InteractiveShape{}
	shape.xPos, shape.yPos = b.interactiveArea.GetPos()
	shape.Width, shape.Height = b.interactiveArea.GetSize()
	return shape, nil
}
