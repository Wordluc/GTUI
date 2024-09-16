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

func main() {
	var e error
	kbr := Keyboard.NewKeyboard()
	core, e = Core.NewGtui(loop, kbr, &Terminal.Terminal{})
	if e != nil {
		panic(e)
	}
	xS, yS := 50, 40
	c, e := Component.CreateTextBox(10, 0, xS, yS, core.CreateStreamingCharacter())
	if e != nil {
		panic(e)
	}
	if e = core.SetCur(0, 0); e != nil {
		panic(e)
	}

	c.SetOnLeave(func() {
		c.GetVisibleArea().SetColor(Color.Get(Color.Gray, Color.None))
	})
	c.SetOnHover(func() {
		c.GetVisibleArea().SetColor(Color.Get(Color.White, Color.None))
	})

	if e = core.InsertComponent(c); e != nil {
		panic(e)
	}
	button, e := Component.CreateButton(60, 1, 10, 5, "Cancella")
	button.SetOnClick(func() {
		button.GetVisibleArea().SetColor(Color.Get(Color.White, Color.None))
		c.Clear()
	})
	button.SetOnRelease(func() {
		button.GetVisibleArea().SetColor(Color.Get(Color.Gray, Color.None))
	})
	button.SetOnLeave(func() {
		button.GetVisibleArea().SetColor(Color.Get(Color.Gray, Color.None))
	})
	if e != nil {
		panic(e)
	}
	core.InsertComponent(button)
	core.Start()
}

func loop(keyb Kd.IKeyBoard) bool {
	var x, y = core.GetCur()
	if keyb.IsSpecialKeyPressed(Kd.KeyArrowDown) {
		y++
	}
	if keyb.IsSpecialKeyPressed(Kd.KeyArrowUp) {
		y--
	}
	if keyb.IsSpecialKeyPressed(Kd.KeyArrowRight) {
		x++
	}
	if keyb.IsSpecialKeyPressed(Kd.KeyArrowLeft) {
		x--
	}

	if keyb.IsSpecialKeyPressed(Kd.KeyCtrlQ) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(Component.IWritableComponent); ok {
				c.StopTyping()
			}
		})
	}

	if keyb.IsSpecialKeyPressed(Kd.KeyEnter) {
		core.Click(x, y)
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(Component.IWritableComponent); ok {
				if !c.IsTyping() {
					c.StartTyping()
				}
			}
		})
	}

	core.SetCur(x, y)
	if keyb.IsKeyPressed('q') {
		return false
	}
	return true
}
