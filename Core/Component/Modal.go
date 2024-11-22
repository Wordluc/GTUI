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
	container *Container
	nullComponent *NullComponent
}

func CreateModal(sizeX, sizeY int) *Modal {
	null := CreateNullComponent(0, 0, sizeX, sizeY)
	contComp:= CreateContainer(0,0)
	contComp.AddComponent(null)
	contComp.SetLayer(modalL1)
	return &Modal{
		nullComponent: null,
		container: contComp,
	}
}

func (b *Modal) SetBackgroundColor(color Color.ColorValue) {
	b.nullComponent.GetRect().SetInsideColor(color)
	b.nullComponent.graphics.Touch()
}

func (b *Modal) AddComponent(componentToAdd Core.IComponent) error {
	if e:=componentToAdd.SetLayer(b.container.GetLayer()+1);e!=nil{
		return e
	}
	b.container.AddComponent(componentToAdd)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
	return nil
}

func (b *Modal) AddDrawing(drawingToAdd Core.IDrawing) error {
	drawingToAdd.SetLayer(drawingToAdd.GetLayer()+modalL1+1)
	b.container.AddDrawing(drawingToAdd)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
	return nil
}

func	(b *Modal)GetComponets() []Core.IComponent{
	return b.container.GetComponents()
}

func (b *Modal) SetPos(x, y int) {
	b.container.SetPos(x, y)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
}

func (b *Modal) GetPos() (int, int) {
	return b.nullComponent.GetPos()
}

func (b *Modal) GetLayer() Core.Layer {
	return b.container.GetLayer()
}

func (b *Modal) SetLayer(layer Core.Layer) error{
	if layer<0 {
		return errors.New("layer can't be negative")
	}
	b.container.SetLayer(layer)
	EventManager.Call(EventManager.ReorganizeElements, []any{b})
	return nil
}

func (b *Modal) GetGraphics() Core.IDrawing {
	return b.container.GetGraphics()
}

func (b *Modal) OnClick() {
}
func (b *Modal) OnRelease() {
}
func (b *Modal) OnHover() {
}
func (b *Modal) OnLeave() {
}

//DO NOT USE
func (b *Modal) GetSize() (int, int) {
	panic("mustn't be called")
}

//DO NOT USE
func (b *Modal) GetVisibleArea() Core.IDrawing {
	panic("mustn't be called")
}
