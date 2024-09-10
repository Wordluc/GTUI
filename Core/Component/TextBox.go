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
	onClick     func()
	onLeave     func()
}

func CreateTextBox(x, y, sizeX, sizeY int, streamText StreamCharacter) *TextBox {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	textBox := Drawing.CreateTextBlock(1, 1, sizeX-1, sizeY-1,  10)
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
	var key rune
	for str := range channel {
		if !b.isTyping {
			return
		}
    key=rune(str[0])
		if key=='\b'{
			b.textBlock.Delete()
			continue
		} 
		b.textBlock.Type(key)
	}
}

func (b *TextBox) StartTyping() {
	b.isTyping = true
	go b.loopTyping()
	b.streamText.Delete()
}
func (b *TextBox) StopTyping() {
	b.streamText.Delete()
	b.isTyping = false
}
func (b *TextBox) OnClick() {
	if b.onClick == nil {
		return
	}
	b.onClick()
}

func (b *TextBox) OnLeave() {
	if b.onLeave == nil {
		return
	}
	b.onLeave()
}
func (b *TextBox) SetOnClick(onClick func()) {
	b.onClick = onClick
}
func (b *TextBox) SetOnLeave(onLeave func()) {
	b.onLeave = onLeave
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

func (b *TextBox) IsTyping() bool {
	return b.isTyping
}
//To delete, the work has to be done in the textBlock by SetCurrentCursor
//so i can catch when move the cursor inside the textBlock but not move the real one outside of it
func (b *TextBox) DiffTotalToXY(x, y int) (int, int) {
	xP, yP := b.textBlock.GetPos()
	x = x - xP
	y = y - yP
	diffX, diffY := b.textBlock.GetTotalCursor()
	diffX = diffX - x 
	diffY = diffY - y - 1
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
func (b *TextBox) SetCurrentPosCursor(x, y int)(int,int) {
	return b.textBlock.SetCurrentCursor(x, y)
}
func (b *TextBox) getShape() (InteractiveShape, error) {
	x, y := b.textBlock.GetPos()
	xDim, yDim := b.textBlock.GetSize()
	shape := InteractiveShape{
		xPos:   x,
		yPos:   y,
		Width:  xDim,
		Height: yDim,
	}
	return shape, nil
}
