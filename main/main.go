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
	kbr:=Keyboard.NewKeyboard()
  core,_=Core.NewGtui(loop,kbr,&Terminal.Terminal{})
	xS,yS:=core.Size()
	c := Component.CreateTextBox(0, 0, xS, yS, core.CreateStreamingCharacter())
	c.StartTyping()
	core.InsertComponent(c)
	core.SetCur(1, 1)
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
