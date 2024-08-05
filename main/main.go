package main

import (
	Core "GTUI"
	"GTUI/Core/Component"
	"GTUI/Core/Utils/Color"
	Kd "GTUI/Keyboard"
)

var x, y = 0, 0
var core *Core.Gtui 
var keyb Kd.IKeyBoard = &Kd.Keyboard{}

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
	b := Component.CreateButton(5, 5, 10, 5, "prova")
	c := Component.CreateTextField(50, 5, 10, 5, Creation(keyb))
	core.InsertComponent(c)
	core.InsertComponent(b)
	core.IRefreshAll()
	core.ISetCursor(x, y)
	keyb.Start(loop)
}

func loop(keyb Kd.IKeyBoard) bool{
	if keyb.TokenPressed(Kd.KeyArrowDown) {
		y++
	}
	if keyb.TokenPressed(Kd.KeyArrowUp) {
		y--
	}
	if keyb.TokenPressed(Kd.KeyArrowRight) {
		x++
	}
	if keyb.TokenPressed(Kd.KeyArrowLeft) {
		x--
	}
	if keyb.TokenPressed(Kd.KeyEnter) {
		y++
	}
	if keyb.IsKeyPressed('c') {
		core.Click(x, y)
	}
	if keyb.IsKeyPressed('q') {
		return false
	}
	core.ISetCursor(x, y)
	core.IRefreshAll()
	core.ISetCursor(x, y)
	return true
}
