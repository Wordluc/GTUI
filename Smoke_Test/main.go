package main

import (
	Core "github.com/Wordluc/GTUI"
	"github.com/Wordluc/GTUI/Core/Component"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
	"github.com/Wordluc/GTUI/Keyboard"
	Kd "github.com/Wordluc/GTUI/Keyboard"
	"github.com/Wordluc/GTUI/Terminal"
)

var core *Core.Gtui
var comp Component.IComponent
func main() {
	kbr := Keyboard.NewKeyboard()
	core, _ = Core.NewGtui(loop, kbr, &Terminal.Terminal{})
	xS, yS := 50, 40

	c,e := Component.CreateTextBox(0, 0, xS, yS, core.CreateStreamingCharacter())
	if e != nil {
		panic(e)
	}
	c.SetOnOut(func() {
		c.GetVisibleArea().SetColor(Color.Get(Color.Gray, Color.None))
	})
	c.SetOnHover(func() {
		c.GetVisibleArea().SetColor(Color.Get(Color.White, Color.None))
	})
   b:=Component.CreateButton(70, 0, 20, 30, "test")
	b.SetOnClick(func() {
		c.ClearAll()
	})
	b.SetOnLeave(func() {
		b.GetVisibleArea().SetColor(Color.Get(Color.Gray, Color.None))
	})
	b.SetOnHover(func() {
		b.GetVisibleArea().SetColor(Color.Get(Color.White, Color.None))
	})
	compComponent:=Component.CreateContainer(0,0)
	if e := compComponent.AddComponent(b); e != nil {
		panic(e)
	}
	if e := compComponent.AddComponent(c); e != nil {
		panic(e)
	}
	comp=compComponent
	button1:=Component.CreateButton(85,0,10,10,"test")
	button1.SetOnLeave(func() {
		button1.GetVisibleArea().SetColor(Color.Get(Color.Gray, Color.None))
	})
	button1.SetOnHover(func() {
		button1.GetVisibleArea().SetColor(Color.Get(Color.White, Color.None))
	})
	if e := core.InsertComponent(compComponent); e != nil {
		panic(e)
	}
	if e:=core.InsertComponent(button1);e!=nil{
		panic(e)
	}
	if e := core.SetCur(1, 1); e != nil {
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

	if keyb.IsKeySPressed(Kd.KeyCtrlS) {
		core.IClear()
		comp.SetPos(x,y)
		core.RefreshComponents()
	}

	core.SetCur(x, y)
	if keyb.IsKeySPressed(Kd.KeyEsc) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(Component.IWritableComponent); ok {
				c.StopTyping()
			}
		})
	}

	if keyb.IsKeySPressed(Kd.KeyEnter) {
		core.Click(x, y)
	}

	if keyb.IsKeySPressed(Kd.KeyCtrlV) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					c.Paste(keyb.GetClickboard())
					core.AllineCursor()
				}
			}
		})
	}

	if keyb.IsKeySPressed(Kd.KeyCtrlA) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					c.SetWrap(!c.GetWrap())
				}
			}
		})
	}

	if keyb.IsKeySPressed(Kd.KeyCtrlC) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					keyb.InsertClickboard(c.Copy())
				}
			}
		})
	}

	if keyb.IsKeySPressed(Kd.KeyCtrlQ) {
		return false
	}
	return true
}
