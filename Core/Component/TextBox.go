package Component

import (
	"GTUI/Core"
	"GTUI/Core/Drawing"
	"GTUI/Core/Utils/Color"
	"errors"
)

type TextBox struct {
	graphics        *Drawing.Container
	interactiveArea *Drawing.Rectangle
	textBox         *Drawing.TextField
	isTyping        bool
	streamText      StreamCharacter
	xTextPosition   int
	nTextLine       int
}

func CreateTextBox(x, y, sizeX, sizeY int, streamText StreamCharacter) *TextBox {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	textBox := Drawing.CreateTextField(1, 1)
	cont.AddChild(rect)
	cont.AddChild(textBox)
	cont.SetPos(x, y)
	return &TextBox{
		graphics:        cont,
		interactiveArea: rect,
		isTyping:        false,
		streamText:      streamText,
		textBox:         textBox,
		nTextLine:       0,
		xTextPosition:   0,
	}
}

func (b *TextBox) loopTyping() {
	channel := b.streamText.Get()
	for str := range channel {
		if !b.isTyping {
			return
		}
		if string(str) != "\n" {
			b.xTextPosition++
		} else {
			b.nTextLine++
			b.xTextPosition = 0
		}
		b.textBox.Type(string(str))
	}
}
func (b *TextBox) Clear() {
	b.textBox.ClearText()
}

func (b *TextBox) OnClick() {
	if b.isTyping {
		return
	}
	b.isTyping = true
	b.interactiveArea.SetColor(Color.Get(Color.Green, Color.None))
	go b.loopTyping()
}

func (b *TextBox) OnLeave() {
	b.isTyping = false
	b.interactiveArea.SetColor(Color.GetDefaultColor())
	b.streamText.Delete()
}

func (b *TextBox) OnRelease() {}

func (b *TextBox) OnHover() {
	b.interactiveArea.SetColor(Color.Get(Color.Gray, Color.None))
}

func (b *TextBox) GetGraphics() Core.IEntity {
	return b.graphics
}

func (b *TextBox) IsOn() bool {
	return b.isTyping
}

func (b *TextBox) CanMoveXCursor(x int) int {
	xP,_:=b.textBox.GetPos()
	x =  x-xP
	if x > b.xTextPosition {
		return x - b.xTextPosition
	}
	return 0
}

func (b *TextBox) CanMoveYCursor(y int) int{
	_,yP:=b.textBox.GetPos()
	y =  y-yP
	if y > b.nTextLine {
	return y - b.nTextLine

	}
	return 0
}

func (b *TextBox) getShape() (InteractiveShape, error) {
	if b.interactiveArea == nil {
		return InteractiveShape{}, errors.New("interactive area is nil")
	}
	shape := InteractiveShape{}
	shape.xPos, shape.yPos = b.interactiveArea.GetPos()
	shape.Width, shape.Height = b.interactiveArea.GetSize()
	return shape, nil
}
