package main

import (
	Core "GTUI"
	"GTUI/Core/Component"
	"GTUI/Keyboard"
	Kd "GTUI/Keyboard"
	"GTUI/Terminal"
)

var core *Core.Gtui 
func main() {
	core,_=Core.NewGtui(loop,&Keyboard.Keyboard{},&Terminal.Terminal{})
	b := Component.CreateButton(0, 0, 10, 5, "Press")
	xS,yS:=core.Size()
	c := Component.CreateTextBox(0, 0, xS, yS, core.CreateStreamingCharacter())
	b.SetOnClick(func() {
		core.IClear()
	})
	core.InsertComponent(c)
	core.Start()
}

func loop(keyb Kd.IKeyBoard) bool{
	var x, y = core.GetCur()
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
	return true
}
