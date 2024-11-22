package Component

import (
	"errors"
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/EventManager"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)
const (
	modalL1 Core.Layer= Core.LMax+iota
	modalL2
	
)
type Modal struct {
	component *Container
	nullComponent *NullComponent
}

func CreateModal(sizeX, sizeY int) *Modal {
	null := CreateNullComponent(0, 0, sizeX, sizeY)
	contComp:= CreateContainer(0,0)
	contComp.AddComponent(null)
	contComp.SetLayer(modalL1)
	return &Modal{
		nullComponent: null,
		component: contComp,
	}
}

func (b *Modal) SetBackgroundColor(color Color.ColorValue) {
	b.nullComponent.GetRect().SetInsideColor(color)
	b.nullComponent.graphics.Touch()
}

func (b *Modal) AddComponent(componentToAdd Core.IComponent) error {
	if e:=componentToAdd.SetLayer(b.component.GetLayer()+1);e!=nil{
		return e
	}
	b.component.AddComponent(componentToAdd)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
	return nil
}

func (b *Modal) AddDrawing(drawingToAdd Core.IDrawing) error {
	drawingToAdd.SetLayer(drawingToAdd.GetLayer()+modalL1+1)
	b.component.AddDrawing(drawingToAdd)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
	return nil
}

func	(b *Modal)GetComponets() []Core.IComponent{
	return b.component.GetComponents()
}

func (b *Modal) SetPos(x, y int) {
	b.component.SetPos(x, y)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
}

func (b *Modal) GetPos() (int, int) {
	return b.nullComponent.GetPos()
}

func (b *Modal) GetLayer() Core.Layer {
	return b.component.GetLayer()
}

func (b *Modal) SetLayer(layer Core.Layer) error{
	if layer<0 {
		return errors.New("layer can't be negative")
	}
	b.component.SetLayer(layer)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
	return nil
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
