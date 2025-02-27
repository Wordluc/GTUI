package Component

import (
	"errors"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/EventManager"
)

//To use when you wanna avoid the interaction under the component
type NullComponent struct {
	graphics    *Drawing.Container
	visibleArea *Drawing.RectangleFull
	onLeave     func()
	onHover     func()
	onRelease   func()
}

func CreateNullComponent(x, y, sizeX, sizeY int) (*NullComponent) {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangleFull(0, 0, sizeX, sizeY)
	cont.AddChild(rect)
	cont.SetPos(x, y)
	return &NullComponent{
		graphics:    cont,
		visibleArea: rect,
	}
}
func (b *NullComponent) GetSize() (int, int) {
	return b.visibleArea.GetSize()
}

func (b *NullComponent) SetPos(x, y int) {
	b.visibleArea.SetPos(x, y)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
}

func (b *NullComponent) GetPos() (int, int) {
	return b.visibleArea.GetPos()
}

func (b *NullComponent) GetRect() *Drawing.RectangleFull  {
	return b.visibleArea
}

func (b *NullComponent) SetLayer(layer Core.Layer)error {
	if layer<0 {
		return errors.New("layer can't be negative")
	}
	b.graphics.SetLayer(layer)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
	return nil
}

func (b *NullComponent) GetLayer() Core.Layer {
	return b.graphics.GetLayer()
}

func (b *NullComponent) GetGraphics() Core.IDrawing {
	return b.graphics
}

func (b *NullComponent) SetOnLeave(onLeave func()) {
	b.onLeave = onLeave
}

func (b *NullComponent) SetOnHover(onHover func()) {
	b.onHover = onHover
}

func (b *NullComponent) SetOnRelease(onRelease func()) {
	b.onRelease = onRelease
}

func (b *NullComponent) OnLeave() {
	if b.onLeave != nil {
		b.onLeave()
	}
}

func (b *NullComponent) OnHover() {
	if b.onHover != nil {
		b.onHover()
	}
}

func (b *NullComponent) OnClick() {
}
func (b *NullComponent) OnRelease() {
}
