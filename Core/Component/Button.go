package Component

import (
	"errors"
	"time"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/EventManager"
)

type Button struct {
	graphics    *Drawing.Container
	visibleArea *Drawing.RectangleFull
	onClick     Core.OnEvent
	onRelease   Core.OnEvent
	onHover     Core.OnEvent
	onLeave     Core.OnEvent
	isClicked   bool
	isActive    bool
}
func CreateButton(x, y, sizeX, sizeY int, text string) *Button {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangleFull(0, 0, sizeX, sizeY)
	textD := Drawing.CreateTextField(0, 0,text)
	rect.SetLayer(Core.L1)
	textD.SetLayer(Core.L2)
	xC, yC := sizeX/2-len(text)/2, sizeY/2
	textD.SetPos(xC, yC)
	cont.AddChild(rect)
	cont.AddChild(textD)
	cont.SetPos(x, y)
	return &Button{
		graphics:    cont,
		visibleArea: rect,
		isClicked:   false,
		isActive:    true,
	}
}

func (b *Button) SetActive(isActive bool)  {
	b.isActive = isActive
}

func (b *Button) SetPos(x, y int) {
	b.graphics.SetPos(x, y)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
}
func (b *Button) GetPos() (int, int) {
	return b.visibleArea.GetPos()
}
func (b *Button) GetLayer() Core.Layer {
	return b.graphics.GetLayer()
}
func (b *Button) SetLayer(layer Core.Layer)error {
	if layer<0 {
		return errors.New("layer can't be negative")
	}
	b.graphics.SetLayer(layer)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
	return nil
}
func (b *Button) GetSize() (int, int) {
	return b.visibleArea.GetSize()
}
func (b *Button) SetOnClick(onClick Core.OnEvent) {
	b.onClick = onClick
}
func (b *Button) SetOnLeave(onLeave Core.OnEvent) {
	b.onLeave = onLeave
}
func (b *Button) SetOnRelease(onRelease Core.OnEvent) {
	b.onRelease = onRelease
}
func (b *Button) SetOnHover(onHover Core.OnEvent) {
	b.onHover = onHover
}
func (b *Button) OnClick() {
	if !b.isActive {
		return
	}
	if b.isClicked {
		return
	}
	if b.onClick == nil {
		return
	}
	b.onClick()
	time.AfterFunc(time.Millisecond*500, func() {
		b.OnRelease()
		EventManager.Call(EventManager.Refresh, []any{b})
	})
	b.isClicked = true
}
func (b *Button) OnRelease() {
	if !b.isActive {
		return
	}
	if b.onRelease == nil {
		return
	}
	b.isClicked = false
	b.onRelease()
}
func (b *Button) OnHover() {
	if !b.isActive {
		return
	}
	if b.onHover != nil {
		b.onHover()
	}
}
func (b *Button) OnLeave() {
	if !b.isActive {
		return
	}
	if b.onLeave != nil {
		b.onLeave()
	}
}
func (b *Button) GetGraphics() Core.IDrawing {
	return b.graphics
}
func (b *Button) GetVisibleArea() *Drawing.RectangleFull {
	return b.visibleArea
}
