package Component

import (
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/EventManager"
)


//To use when you wanna avoid the interaction under the component
type NullComponent struct {
	graphics    *Drawing.Container
	visibleArea *Drawing.Rectangle
}

func CreateNullComponent(x, y, sizeX, sizeY int) (*NullComponent) {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangle(0, 0, sizeX, sizeY)
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
func (b *NullComponent) SetLayer(layer Core.Layer) {
	b.graphics.SetLayer(layer)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
}
func (b *NullComponent) GetLayer() Core.Layer {
	return b.graphics.GetLayer()
}

func (b *NullComponent) OnClick() {
}

func (b *NullComponent) OnLeave() {
}

func (b *NullComponent) OnRelease() {
}

func (b *NullComponent) OnHover() {
}

func (b *NullComponent) GetGraphics() Core.IDrawing {
	return b.graphics
}

func (b *NullComponent) GetVisibleArea() *Drawing.Rectangle {
	return b.visibleArea
}
