package Component

import (
	"GTUI/Core"
	"GTUI/Core/Drawing"
	"GTUI/Core/Utils/Color"
)

type TextBox struct {
	graphics    *Drawing.Container
	visibleArea *Drawing.Rectangle
	textBlock   *Drawing.TextBlock
	isTyping    bool
	streamText  StreamCharacter
}

func CreateTextBox(x, y, sizeX, sizeY int, streamText StreamCharacter) *TextBox {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	textBox := Drawing.CreateTextBlock(1, 1, 10, 100)
	cont.AddChild(rect)
	cont.AddChild(textBox)
	cont.SetPos(x, y)
	return &TextBox{
		graphics:    cont,
		visibleArea: rect,
		isTyping:    false,
		streamText:  streamText,
		textBlock:   textBox,
	}
}

func (b *TextBox) loopTyping() {
	channel := b.streamText.Get()
	for str := range channel {
		if !b.isTyping {
			return
		}
		b.textBlock.Type(rune(str[0]))
	}
}

func (b *TextBox) OnClick() {
	if b.isTyping {
		return
	}
	b.isTyping = true
	b.visibleArea.SetColor(Color.Get(Color.Green, Color.None))
	go b.loopTyping()
}

func (b *TextBox) OnLeave() {
	b.isTyping = false
	b.visibleArea.SetColor(Color.GetDefaultColor())
	b.streamText.Delete()
}

func (b *TextBox) OnRelease() {}

func (b *TextBox) OnHover() {
	if b.isTyping {
		return
	}
	b.visibleArea.SetColor(Color.Get(Color.Gray, Color.None))
}

func (b *TextBox) GetGraphics() Core.IEntity {
	return b.graphics
}

func (b *TextBox) IsWritable() bool {
	return b.isTyping
}

func (b *TextBox) DiffTotalToXY(x, y int) (int, int) {
	xP, yP := b.textBlock.GetPos()
	x = x - xP
	y = y - yP
	diffX, diffY := b.textBlock.GetTotalCursor()
	diffX = diffX - x
	diffY = diffY - y
	return diffX, diffY
}

func (b *TextBox) DiffCurrentToXY(x, y int) (int, int) {
	xP, yP := b.textBlock.GetPos()
	x = x - xP
	y = y - yP
	diffX, diffY := b.textBlock.GetCurrentCursor()
	diffX = diffX - x
	diffY = diffY - y
	return diffX, diffY
}
func (b *TextBox) SetCurrentPos(x, y int) {
	xP, yP := b.textBlock.GetPos()
	b.textBlock.SetCurrent(x-xP, y-yP)
	return
}
func (b *TextBox) getShape() (InteractiveShape, error) {
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
