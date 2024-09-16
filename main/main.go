package main

import (
	Core "GTUI"
	"GTUI/Core/Component"
	"GTUI/Core/Utils"
	"GTUI/Core/Utils/Color"
	"GTUI/Keyboard"
	Kd "GTUI/Keyboard"
	"GTUI/Terminal"
)

var core *Core.Gtui
var components []Component.IComponent

func main() {
	kbr := Keyboard.NewKeyboard()
	var totalError Utils.CError
	var e error
	core, e = Core.NewGtui(loop, kbr, &Terminal.Terminal{}); 
	totalError.Add(e)
	xS, yS := 50, 40
	c, e := Component.CreateTextBox(0, 0, xS, yS, core.CreateStreamingCharacter())
	totalError.Add(e)
	c.SetOnLeave(func() {
		c.GetVisibleArea().SetColor(Color.Get(Color.Gray, Color.None))
	})
	c.SetOnHover(func() {
		c.GetVisibleArea().SetColor(Color.Get(Color.White, Color.None))
	})
	components = append(components, c)
	totalError.Add(core.InsertComponent(c))
	totalError.Add(core.SetCur(0, 0))
	if totalError.HasError() {
		panic(totalError)
	}
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
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(Component.IWritableComponent); ok {
				if !c.IsTyping() {
					c.StartTyping()
				}
			}
		})
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
