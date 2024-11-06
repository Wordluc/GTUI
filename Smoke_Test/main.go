package main

import (
	Core "github.com/Wordluc/GTUI"
	"github.com/Wordluc/GTUI/Core/Component"
	"github.com/Wordluc/GTUI/Core/Drawing"
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
	c,e := Component.CreateTextBox(5, 5, 20, 10, core.CreateStreamingCharacter())
	if e != nil {
		panic(e)
	}
	c.SetOnOut(func() {
		c.GetVisibleArea().SetColor(Color.Get(Color.Gray, Color.None))
	})
	c.SetOnHover(func() {
		c.GetVisibleArea().SetColor(Color.Get(Color.White, Color.None))
	})
	rect:=Drawing.CreateRectangle(2,2,xS-1,yS-1)
	rect1:=Drawing.CreateRectangle(16,26,10,10)
	contDraw:=Drawing.CreateContainer(0,0)
	contDraw.AddChild(rect)
	contDraw.AddChild(rect1)
	compComponent:=Component.CreateContainer(0,0)
	compComponent.AddDrawing(contDraw)
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
	button1.SetOnClick(func() {
		rect.SetColor(Color.Get(Color.Red, Color.None))
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
	if keyb.IsKeySPressed(Keyboard.Down) {
		y++
	}
	if keyb.IsKeySPressed(Keyboard.Up) {
		y--
	}
	if keyb.IsKeySPressed(Keyboard.Right) {
		x++
	}
	if keyb.IsKeySPressed(Keyboard.Left) {
		x--
	}

	if keyb.IsKeySPressed(Keyboard.CtrlS) {
		core.IClear()
		comp.SetPos(x,y)
		core.RefreshComponents()
	}

	core.SetCur(x, y)
	if keyb.IsKeySPressed(Keyboard.Esc) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(Component.IWritableComponent); ok {
				c.StopTyping()
			}
		})
	}

	if keyb.IsKeySPressed(Keyboard.Enter) {
		core.Click(x, y)
	}

	if keyb.IsKeySPressed(Keyboard.CtrlV) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					c.Paste(keyb.GetClickboard())
					core.AllineCursor()
				}
			}
		})
	}

	if keyb.IsKeySPressed(Keyboard.CtrlA) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					c.SetWrap(!c.GetWrap())
				}
			}
		})
	}

	if keyb.IsKeySPressed(Keyboard.CtrlC) {
		core.EventOn(x, y, func(c Component.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					keyb.InsertClickboard(c.Copy())
				}
			}
		})
	}

	if keyb.IsKeySPressed(Keyboard.CtrlQ) {
		return false
	}
	return true
}
