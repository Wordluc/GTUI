package main

import (
	"atomicgo.dev/keyboard/keys"
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
	button1:=Component.CreateButton(100,0,10,10,"test")
	button1.SetOnLeave(func() {
		button1.GetVisibleArea().SetColor(Color.Get(Color.Gray, Color.None))
	})
	button1.SetOnHover(func() {
		button1.GetVisibleArea().SetColor(Color.Get(Color.White, Color.None))
	})
	comp.SetPos(10,10)
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
	if keyb.IsKeySPressed(keys.Down) {
		y++
	}
	if keyb.IsKeySPressed(keys.Up) {
		y--
	}
	if keyb.IsKeySPressed(keys.Right) {
		x++
	}
	if keyb.IsKeySPressed(keys.Left) {
		x--
	}

	if keyb.IsKeySPressed(keys.CtrlS) {
		core.IClear()
		comp.SetPos(x,y)
		core.RefreshComponents()
	}

	core.SetCur(x, y)
	if keyb.IsKeySPressed(keys.Esc) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(Component.IWritableComponent); ok {
				c.StopTyping()
			}
		})
	}

	if keyb.IsKeySPressed(keys.Enter) {
		core.Click(x, y)
	}

	if keyb.IsKeySPressed(keys.CtrlV) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					c.Paste(keyb.GetClickboard())
					core.AllineCursor()
				}
			}
		})
	}

	if keyb.IsKeySPressed(keys.CtrlA) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					c.SetWrap(!c.GetWrap())
				}
			}
		})
	}

	if keyb.IsKeySPressed(keys.CtrlC) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					keyb.InsertClickboard(c.Copy())
				}
			}
		})
	}

	if keyb.IsKeySPressed(keys.CtrlQ) {
		return false
	}
	return true
}
