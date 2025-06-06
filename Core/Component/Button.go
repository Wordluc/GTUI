package Component

import (
	"errors"
	"time"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/EventManager"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

type Button struct {
	graphics       *Drawing.Container
	visibleArea    *Drawing.RectangleFull
	onClick        Core.OnEvent
	onRelease      Core.OnEvent
	onHover        Core.OnEvent
	onLeave        Core.OnEvent
	isClicked      bool
	isActive       bool
	OnClickColor   Color.Color
	OnReleaseColor Color.Color
	OnLeaveColor   Color.Color
	OnHoverColor   Color.Color
}

func CreateButton(x, y, sizeX, sizeY int, text string) *Button {
	cont := Drawing.CreateContainer(0, 0)
	rect := Drawing.CreateRectangleFull(0, 0, sizeX, sizeY)
	textD := Drawing.CreateTextField(0, 0, text)
	rect.SetLayer(Core.L1)
	textD.SetLayer(Core.L2)
	xC, yC := sizeX/2-len(text)/2, sizeY/2
	textD.SetPos(xC, yC)
	cont.AddDrawings(rect)
	cont.AddDrawings(textD)
	cont.SetPos(x, y)
	return &Button{
		graphics:       cont,
		visibleArea:    rect,
		isClicked:      false,
		isActive:       true,
		OnClickColor:   Color.Get(Color.Red, Color.None),
		OnReleaseColor: Color.Get(Color.Gray, Color.None),
		OnLeaveColor:   Color.Get(Color.Gray, Color.None),
		OnHoverColor:   Color.GetDefaultColor(),
	}
}

func (b *Button) SetActive(isActive bool) {
	b.isActive = isActive
	if !b.isActive {
		time.AfterFunc(0, func() {
			EventManager.Call(EventManager.Refresh, b)
			b.isActive = true
			b.OnLeave()
			b.isActive = false
		})
	}
}

func (b *Button) GetActive() bool {
	return b.isActive
}

func (b *Button) SetPos(x, y int) {
	b.graphics.SetPos(x, y)
	EventManager.Call(EventManager.ReorganizeElements, b)
}
func (b *Button) GetPos() (int, int) {
	return b.visibleArea.GetPos()
}
func (b *Button) GetLayer() Core.Layer {
	return b.graphics.GetLayer()
}
func (b *Button) SetLayer(layer Core.Layer) error {
	if layer < 0 {
		return errors.New("layer can't be negative")
	}
	b.graphics.SetLayer(layer)
	EventManager.Call(EventManager.ReorganizeElements, b)
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
	if b.isClicked {
		return
	}
	b.GetVisibleArea().SetBorderColor(b.OnClickColor)
	if b.onClick != nil {
		b.onClick()
	}
	time.AfterFunc(time.Millisecond*500, func() {
		b.OnRelease()
		EventManager.Call(EventManager.Refresh, b)
	})
	b.isClicked = true
}

func (b *Button) OnRelease() {
	if b.onRelease != nil {
		b.onRelease()
	}
	b.GetVisibleArea().SetBorderColor(b.OnReleaseColor)
	b.isClicked = false
}

func (b *Button) OnHover() {
	if b.onHover != nil {
		b.onHover()
	}

	b.GetVisibleArea().SetBorderColor(b.OnHoverColor)
}

func (b *Button) OnLeave() {
	if b.onLeave != nil {
		b.onLeave()
	}
	b.GetVisibleArea().SetBorderColor(b.OnLeaveColor)
}

func (b *Button) GetGraphics() []Core.IDrawing {
	return b.graphics.GetDrawings()
}

func (b *Button) GetGraphic() *Drawing.Container {
	return b.graphics
}

func (b *Button) GetVisibleArea() *Drawing.RectangleFull {
	return b.visibleArea
}
