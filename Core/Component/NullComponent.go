package Component

import (
	"errors"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/EventManager"
)

// To use when you wanna avoid the interaction under the component
type NullComponent struct {
	graphics    *Drawing.Container
	visibleArea *Drawing.RectangleFull
	onLeave     func()
	onHover     func()
	onRelease   func()
	active      bool
}

func CreateNullComponent(x, y, sizeX, sizeY int) *NullComponent {
	cont := Drawing.CreateContainer(x, y)
	rect := Drawing.CreateRectangleFull(0, 0, sizeX, sizeY)
	rect.SetPos(x, y)
	cont.AddDrawings(rect)
	return &NullComponent{
		graphics: cont,
		active:   true,
	}
}
func (b *NullComponent) GetSize() (int, int) {
	return b.visibleArea.GetSize()
}

func (b *NullComponent) SetPos(x, y int) {
	b.visibleArea.SetPos(x, y)
	EventManager.Call(EventManager.ReorganizeElements, b)
}

func (b *NullComponent) GetPos() (int, int) {
	return b.visibleArea.GetPos()
}

func (b *NullComponent) SetActive(active bool) {
	b.active = active
}

func (b *NullComponent) GetActive() bool {
	return b.active
}

func (b *NullComponent) GetRect() *Drawing.RectangleFull {
	return b.visibleArea
}

func (b *NullComponent) SetLayer(layer Core.Layer) error {
	if layer < 0 {
		return errors.New("layer can't be negative")
	}
	b.visibleArea.SetLayer(layer)
	EventManager.Call(EventManager.ReorganizeElements, b)
	return nil
}

func (b *NullComponent) GetLayer() Core.Layer {
	return b.visibleArea.GetLayer()
}

func (b *NullComponent) GetGraphics() []Core.IDrawing {
	return b.graphics.GetDrawings()
}

func (b *NullComponent) GetGraphic() *Drawing.Container {
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
