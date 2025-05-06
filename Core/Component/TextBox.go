package Component

import (
	"errors"
	"time"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/EventManager"
)

type TextBox struct {
	graphics    *Drawing.Container
	visibleArea *Drawing.Rectangle
	textBlock   *Drawing.TextBlock
	isTyping    bool
	isClicked   bool
	streamText  StreamCharacter
	wrap        bool
	onClick     func()
	onLeave     func()
	onHover     func()
}

func CreateTextBox(x, y, sizeX, sizeY int, streamText StreamCharacter) (*TextBox, error) {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	textBox := Drawing.CreateTextBlock(1, 1, sizeX-2, sizeY-2, 10)
	if e := cont.AddDrawings(rect); e != nil {
		return nil, e
	}
	if e := cont.AddDrawings(textBox); e != nil {
		return nil, e
	}
	cont.SetPos(x, y)
	return &TextBox{
		graphics:    cont,
		visibleArea: rect,
		isTyping:    false,
		isClicked:   false,
		streamText:  streamText,
		textBlock:   textBox,
	}, nil
}
func (b *TextBox) GetSize() (int, int) {
	return b.visibleArea.GetSize()
}

func (b *TextBox) SetPos(x, y int) {
	b.graphics.SetPos(x, y)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
}
func (b *TextBox) GetPos() (int, int) {
	return b.visibleArea.GetPos()
}
func (b *TextBox) SetLayer(layer Core.Layer)error {
	if layer<0 {
		return errors.New("layer can't be negative")
	}
	b.graphics.SetLayer(layer)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
	return nil
}
func (b *TextBox) GetLayer() Core.Layer {
	return b.graphics.GetLayer()
}
func (b *TextBox) loopTyping() {
	channel := b.streamText.Get()
	var key rune
	for str := range channel {
		if !b.isTyping {
			return
		}
		for _, key = range []rune(str) {
			if key == '\b' {
				b.textBlock.Delete()
				continue
			}
			b.textBlock.Type(key)
			break
		}
		channel <- ""
	}
}

func (v *TextBox) ClearAll() {
	v.textBlock.ClearAll()
}

func (v *TextBox) DeleteLastCharacter() {
	v.textBlock.Delete()
}

func (v *TextBox) Paste(text string) {
	for _, char := range []rune(text) {
		if char == '\r' {
			continue
		}
		v.textBlock.Type(char)
	}
}

func (v *TextBox) GetSelectedText() string {
	return v.textBlock.GetSelectedText()
}

func (t *TextBox) GetText()string {
	return t.textBlock.GetText()
}

func (t *TextBox) SetWrap(isOn bool) {
	t.textBlock.SetWrap(isOn)
}

func (t *TextBox) IsInSelectingMode() bool {
	return t.textBlock.GetWrap()
}
func (b *TextBox) StartTyping() {
	if b.isTyping {
		return
	}
	b.isTyping = true
	go b.loopTyping()
	b.streamText.Delete()
}
func (b *TextBox) StopTyping() {
	b.streamText.Delete()
	b.isTyping = false
}

func (b *TextBox) OnClick() {
	if b.isClicked {
		return
	}
	b.isClicked = true
	if b.onClick != nil {
		b.onClick()
	}
	b.StartTyping()
	time.AfterFunc(time.Millisecond*1000, func() {
		b.OnRelease()
		EventManager.Call(EventManager.Refresh, []any{b})
	})
}

func (b *TextBox) OnLeave() {
	if b.onLeave != nil {
		b.onLeave()	
	}
	b.StopTyping()
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

func (b *TextBox) OnRelease() {
	b.isClicked = false
}

func (b *TextBox) OnHover() {
	if b.onHover != nil {
		b.onHover()	
	}
}

func (b *TextBox) GetGraphics() []Core.IDrawing {
	return b.graphics.GetDrawings()
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
func (b *TextBox) SetCurrentPosCursor(x, y int) (int, int) {
	return b.textBlock.SetCursor_Relative(x, y)
}
