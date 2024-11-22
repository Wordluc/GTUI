package Component

import (
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/EventManager"
)
const (
	modalL1 Core.Layer= Core.LMax+iota
	modalL2
	
)
type Modal struct {
	component *Container
	visibleArea Core.IDrawing
}
func CreateModal(x, y, sizeX, sizeY int,content *Container) *Modal {
	null := CreateNullComponent(0, 0, sizeX, sizeY)
	contComp:= CreateContainer(0,0)

	contComp.AddComponent(null)
	contComp.AddComponent(content)

	content.SetLayer(modalL2)
	contComp.SetLayer(modalL1)

	content.SetPos(1,1)
	contComp.SetPos(x, y)
	
	return &Modal{
		visibleArea: null.GetGraphics(),
		component: contComp,
	}
}

func	(b *Modal)GetComponets() []Core.IComponent{
	return b.component.GetComponents()
}

func (b *Modal) SetPos(x, y int) {
	b.component.SetPos(x, y)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
}

func (b *Modal) GetPos() (int, int) {
	return b.visibleArea.GetPos()
}

func (b *Modal) GetLayer() Core.Layer {
	return b.component.GetLayer()
}

func (b *Modal) SetLayer(layer Core.Layer) {
	b.component.SetLayer(layer)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
}

func (b *Modal) GetGraphics() Core.IDrawing {
	return b.component.GetGraphics()
}

func (b *Modal) GetSize() (int, int) {
	panic("mustn't be called")
}

func (b *Modal) GetVisibleArea() Core.IDrawing {
	panic("mustn't be called")
}

func (b *Modal) OnClick() {
}
func (b *Modal) OnRelease() {
}
func (b *Modal) OnHover() {
}
func (b *Modal) OnLeave() {
}
