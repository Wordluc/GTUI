package main

import (
	Core "GTUI"
	"GTUI/Core/Component"
	"GTUI/Core/Utils/Color"
	Kd "GTUI/Keyboard"
)

var core *Core.Gtui 
var keyb Kd.IKeyBoard = &Kd.Keyboard{}
var x, y int
func Creation(k Kd.IKeyBoard) Component.StreamCharacter {
	stream := Component.StreamCharacter{}
	stream.Get = func() chan string {
		i := k.NewChannel()
		stream.IChannel = i
		return k.GetChannels()[i]
	}
	stream.Delete = func() {
		 k.DeleteChannel(stream.IChannel) 
	}
	return stream
}
func main() {
	core,_=Core.NewGtui()
	defer keyb.Stop()
	defer core.Close()
	core.ISetGlobalColor(Color.GetDefaultColor())
	b := Component.CreateButton(5, 5, 10, 3, "Press")
	c := Component.CreateTextField(50, 5, 20, 10, Creation(keyb))
	core.InsertComponent(c)
	core.InsertComponent(b)
	core.IRefreshAll()
	core.SetCur(x, y)
	keyb.Start(loop)
}

func loop(keyb Kd.IKeyBoard) bool{
	core.IRefreshAll()
	if keyb.IsKeySPressed(Kd.KeyArrowDown) {
		y++
	}
	if keyb.IsKeySPressed(Kd.KeyArrowUp) {
		y--
	}
	if keyb.IsKeySPressed(Kd.KeyArrowRight) {
		x++
	}
	if keyb.IsKeySPressed(Kd.KeyArrowLeft) {
		x--
	}
	if keyb.IsKeySPressed(Kd.KeyEnter) {
		y++
	}
	core.SetCur(x, y)
	if keyb.IsKeyPressed('c') {
		core.Interact(x, y, Component.OnClick)
	}
	if keyb.IsKeyPressed('q') {
		return false
	}
	core.IRefreshAll()
	return true
}
