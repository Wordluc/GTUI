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
	x,y=core.GetCur()
	core.ISetGlobalColor(Color.GetDefaultColor())
	b := Component.CreateButton(0, 0, 10, 5, "Press")
	xS,yS:=core.Size()
	c := Component.CreateTextBox(0, 0, xS, yS, Creation(keyb))
	b.SetOnClick(func() {
		core.IClear()
	})
	core.InsertComponent(c)
	//core.InsertComponent(b)
	core.SetCur(x, y)
	core.IRefreshAll()
	keyb.Start(loop)
}

func loop(keyb Kd.IKeyBoard) bool{
	core.IRefreshAll()
	x, y = core.GetCur()
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
	core.SetCur(x, y)
	if keyb.IsKeyPressed('c') {
		core.Click(x, y)
	}
	if keyb.IsKeyPressed('q') {
		return false
	}
	core.IClear()
	core.IRefreshAll()
	return true
}
