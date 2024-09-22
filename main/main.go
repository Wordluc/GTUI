package main

import (
	Core "GTUI"
	"GTUI/Core/Component"
	"GTUI/Core/Utils/Color"
	"GTUI/Keyboard"
	Kd "GTUI/Keyboard"
	"GTUI/Terminal"
)

var core *Core.Gtui
var components []Component.IComponent

func main() {
	kbr := Keyboard.NewKeyboard()
	core, _ = Core.NewGtui(loop, kbr, &Terminal.Terminal{})
	xS, yS := 50, 40
	c := Component.CreateTextBox(0, 0, xS, yS, core.CreateStreamingCharacter())
	c.SetOnLeave(func() {
		c.GetVisibleArea().SetColor(Color.Get(Color.Gray,Color.None))
	})
	c.SetOnHover(func() {
		c.GetVisibleArea().SetColor(Color.Get(Color.White,Color.None))
	})
	components = append(components, c)
	if e:=core.InsertComponent(c) ;e!= nil {
		panic(e)
	}
	if e:=core.SetCur(1, 1); e != nil {
		panic(e)
	}
	core.Start()
}

func loop(keyb Kd.IKeyBoard) bool {
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
	if keyb.IsKeySPressed(Kd.KeyCtrlQ) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(Component.IWritableComponent); ok {
				c.StopTyping()
			}
		})
	}

	if keyb.IsKeySPressed(Kd.KeyEnter) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(Component.IWritableComponent); ok {
				if !c.IsTyping() {
					c.StartTyping()
				}
			}
		})
	}

	if keyb.IsKeySPressed(Kd.KeyCtrlV) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping(){
					c.Paste(keyb.GetClickboard())
					core.AllineCursor()
				}
			}
		})
	}

	if keyb.IsKeySPressed(Kd.KeyCtrlA) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping(){
					c.Wrap(true)
					
				}
			}
		})
	}
	if keyb.IsKeyPressed('c') {
		core.Click(x, y)
	}
	if keyb.IsKeyPressed('q') {
		return false
	}
	return true
}
