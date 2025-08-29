package Component

import (
	"errors"
	"time"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/EventManager"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

type TextBox struct {
	graphics           *Drawing.Container
	visibleArea        *Drawing.Rectangle
	textBlock          *Drawing.TextBlock
	isTyping           bool
	isClicked          bool
	streamText         StreamCharacter
	wrap               bool
	onClick            func()
	onLeave            func()
	onHover            func()
	OnTypingColor      Color.Color
	OnLeaveTypingColor Color.Color
	OnHoverColor       Color.Color
	IsOneLine          bool
	active             bool
}

func CreateTextBox(x, y, sizeX, sizeY int, streamText StreamCharacter) (*TextBox, error) {
	cont := Drawing.CreateContainer()
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
	textBox := Drawing.CreateTextBlock(1, 1, sizeX-2, sizeY-2, 10)
	if e := cont.AddDrawings(rect); e != nil {
		return nil, e
	}
	rect.SetLayer(Core.L1)
	textBox.SetLayer(Core.L2)
	if e := cont.AddDrawings(textBox); e != nil {
		return nil, e
	}
	cont.SetPos(x, y)
	return &TextBox{
		graphics:           cont,
		visibleArea:        rect,
		isTyping:           false,
		isClicked:          false,
		streamText:         streamText,
		textBlock:          textBox,
		OnTypingColor:      Color.Get(Color.Blue, Color.None),
		OnLeaveTypingColor: Color.Get(Color.Gray, Color.None),
		OnHoverColor:       Color.GetDefaultColor(),
		active:             true,
	}, nil
}
func (b *TextBox) GetSize() (int, int) {
	return b.visibleArea.GetSize()
}

func (b *TextBox) SetPos(x, y int) {
	b.graphics.SetPos(x, y)
	EventManager.Call(EventManager.ReorganizeElements, b)
}
func (b *TextBox) SetActive(active bool) {
	b.active = active
}
func (b *TextBox) GetActive() bool {
	return b.active
}
func (b *TextBox) GetPos() (int, int) {
	return b.visibleArea.GetPos()
}
func (b *TextBox) SetLayer(layer Core.Layer) error {
	if layer < 0 {
		return errors.New("layer can't be negative")
	}
	b.graphics.SetLayer(layer)
	EventManager.Call(EventManager.ReorganizeElements, b)
	return nil
}
func (b *TextBox) GetLayer() Core.Layer {
	return b.graphics.GetLayer()
}
func (b *TextBox) loopTyping() {
	EventManager.Call(EventManager.CursorAlign)
	channel := b.streamText.Get()
	var key rune
	for str := range channel {
		if !b.isTyping {
			return
		}
		if b.IsOneLine && str == "\n" {
			channel <- ""
			b.StopTyping()
			break
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
	v.textBlock.Paste(text)
}

func (v *TextBox) GetSelectedText() string {
	return v.textBlock.GetSelectedText()
}

func (t *TextBox) GetText() string {
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
	b.GetVisibleArea().SetColor(b.OnTypingColor)
	b.isTyping = true
	//Align cursor
	b.textBlock.Type(' ')
	b.textBlock.Delete()
	//
	go b.loopTyping()
	b.streamText.Delete()
}
func (b *TextBox) StopTyping() {
	b.streamText.Delete()
	b.isTyping = false
	b.graphics.Touch()
	EventManager.Call(EventManager.Refresh)
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
	time.AfterFunc(time.Millisecond*100, func() {
		b.OnRelease()
		EventManager.Call(EventManager.Refresh, b)
	})
}

func (b *TextBox) OnLeave() {
	if b.onLeave != nil {
		b.onLeave()
	}
	b.streamText.Delete()
	b.isTyping = false
	b.GetVisibleArea().SetColor(b.OnLeaveTypingColor)

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

	b.GetVisibleArea().SetColor(b.OnHoverColor)
}

func (b *TextBox) GetGraphics() []Core.IDrawing {
	return b.graphics.GetGraphics()
}

func (b *TextBox) GetGraphic() *Drawing.Container {
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
func (b *TextBox) SetCurrentPosCursor(x, y int) (int, int) {
	return b.textBlock.SetCursor_Relative(x, y)
}
