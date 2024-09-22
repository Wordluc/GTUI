package Component

import (
	"GTUI/Core"
	"GTUI/Core/Drawing"
)

type TextBox struct {
	graphics    *Drawing.Container
	visibleArea *Drawing.Rectangle
	textBlock   *Drawing.TextBlock
	isTyping    bool
	streamText  StreamCharacter
	wrap        bool
	onClick     func()
	onLeave     func()
	onHover     func()
}

func CreateTextBox(x, y, sizeX, sizeY int, streamText StreamCharacter) *TextBox {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	textBox := Drawing.CreateTextBlock(1, 1, sizeX-2, sizeY-2,  10)
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
		for _,key=range []rune(str){
			if key=='\b'{
				b.textBlock.Delete()
				continue
			} 
			b.textBlock.Type(key)
		}
	}
}

func (v *TextBox) Paste(text string) {
	for _,char:=range []rune(text){
		if char=='\r'{
       continue
		}	
		v.textBlock.Type(char)
	}
}

func (v *TextBox) Copy() string {
	return v.textBlock.GetSelectedText()
}

func (t *TextBox) SetWrap(isOn bool) {
	t.textBlock.SetWrap(isOn)
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
func (b *TextBox) SetOnHover(onHover func()) {
	b.onHover = onHover
}
func (b *TextBox) OnRelease() {}

func (b *TextBox) OnHover() {
	b.onHover()
	if b.isTyping {
		return
	}
}

func (b *TextBox) GetGraphics() Core.IEntity {
	return b.graphics
}

func (b *TextBox) GetVisibleArea() *Drawing.Rectangle {
	return b.visibleArea
}

func (b *TextBox) IsTyping() bool {
	return b.isTyping
}

func (b *TextBox) DiffCurrentToXY(x, y int) (int, int) {
	xP, yP := b.textBlock.GetPos()
	x = x - xP
	y = y - yP
	diffX, diffY := b.textBlock.GetCursor_Relative()
	diffX = diffX - x
	diffY = diffY - y
	return diffX, diffY
}
func (b *TextBox) SetCurrentPosCursor(x, y int)(int,int) {
	return b.textBlock.SetCursor_Relative(x, y)
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
