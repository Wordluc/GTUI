package main

import (
	"time"

	"github.com/Wordluc/GTUI"
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Component"
	"github.com/Wordluc/GTUI/Core/Drawing"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
	"github.com/Wordluc/GTUI/Keyboard"
	Kd "github.com/Wordluc/GTUI/Keyboard"
	"github.com/Wordluc/GTUI/Terminal"
)

func setDefaultEventButton(b *Component.Button,idle Color.ColorValue,active Color.ColorValue,click Color.ColorValue) {
	b.SetOnHover(func() {
		b.GetVisibleArea().SetBorderColor(Color.Get(active, Color.None))
	})
	b.SetOnLeave(func() {
		b.GetVisibleArea().SetBorderColor(Color.Get(idle, Color.None))
	})
	b.SetOnClick(func() {
		b.GetVisibleArea().SetBorderColor(Color.Get(click, Color.None))
	})
	b.SetOnRelease(func() {
		b.GetVisibleArea().SetBorderColor(Color.Get(idle, Color.None))
	})
}
var comp Core.IComponent
var button2 *Component.Button
var modal *Component.Modal
func main() {
	kbr := Keyboard.NewKeyboard()
	core, _ := GTUI.NewGtui(loop, kbr, &Terminal.Terminal{})
	xS, yS := 50, 40
	c,e := Component.CreateTextBox(30, 5, 20, 10, core.CreateStreamingCharacter())
	c.SetLayer(Core.L3)
	if e != nil {
		panic(e)
	}
	c.SetOnLeave(func() {
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
	contDraw.SetVisibility(true)
	compComponent:=Component.CreateContainer(0,0)
	compComponent.AddDrawing(contDraw)
	if e := compComponent.AddComponent(c); e != nil {
		panic(e)
	}
	comp=compComponent
	button1:=Component.CreateButton(100,0,10,10,"test")
	button1.SetActive(false)
	button1.SetLayer(Core.L4)
	button1.GetVisibleArea().SetInsideColor(Color.White)
	button1.SetOnLeave(func() {
		button1.GetVisibleArea().SetBorderColor(Color.Get(Color.Gray, Color.None))
	})
	button1.SetOnHover(func() {
		button1.GetVisibleArea().SetBorderColor(Color.Get(Color.White, Color.None))
	})
	button1.SetOnRelease(func() {
		button1.GetVisibleArea().SetBorderColor(Color.Get(Color.Gray, Color.None))
	})
	button1.SetOnClick(func() {
		button1.GetVisibleArea().SetBorderColor(Color.Get(Color.Red, Color.None))
		time.AfterFunc(time.Millisecond*1000, func() {
			button1.OnRelease()
		})
	})
	button2=Component.CreateButton(62,15,10,10,"test")
	button2.SetLayer(Core.L1)
	button2.GetVisibleArea().SetInsideColor(Color.Blue)
	button2.SetOnLeave(func() {
		button2.GetVisibleArea().SetBorderColor(Color.Get(Color.Gray, Color.None))
	})
	button2.SetOnHover(func() {
		button2.GetVisibleArea().SetBorderColor(Color.Get(Color.White, Color.None))
	})
	button2.SetOnRelease(func() {
		button2.GetVisibleArea().SetBorderColor(Color.Get(Color.Gray, Color.None))
	})
	button2.SetOnClick(func() {
		button2.GetVisibleArea().SetBorderColor(Color.Get(Color.Red, Color.None))
		time.AfterFunc(time.Millisecond*1000, func() {
			button2.OnRelease()
		})
	})
	comp.SetPos(10,10)
	if e := core.AddComponent(compComponent); e != nil {
		panic(e)
	}
	if e:=core.AddComponent(button1);e!=nil{
		panic(e)
	}
	if e:=core.AddComponent(button2);e!=nil{
		panic(e)
	}
	if e := core.SetCur(1, 1); e != nil {
		panic(e)
	}

	text:=Drawing.CreateTextField(10,1,"modal")
	text.SetColor(Color.Get(Color.White,Color.Red))
	ok:=Component.CreateButton(2,5,10,3,"ok")
	esplodi:=Component.CreateButton(18,5,10,3,"esplodi")

	setDefaultEventButton(ok,Color.Gray,Color.White,Color.Red)
	setDefaultEventButton(esplodi,Color.Gray,Color.White,Color.Red)

	modal=Component.CreateModal(30,10)
	modal.SetBackgroundColor(Color.Gray)
	modal.AddDrawing(text)
	modal.AddComponent(esplodi)
	modal.AddComponent(ok)
	if e:=core.AddComponent(modal);e!=nil{
		panic(e)
	}
	core.Start()
}

func loop(keyb Kd.IKeyBoard,core *GTUI.Gtui) bool {
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
		button2.SetPos(x, y)
	}

	if keyb.IsKeySPressed(Keyboard.CtrlP) {
		modal.SetPos(x, y)
	}

	core.SetCur(x, y)
	if keyb.IsKeySPressed(Keyboard.Esc) {
		core.CallEventOn(x, y, func(c Core.IComponent) {
			if c, ok := c.(Core.IWritableComponent); ok {
				c.StopTyping()
			}
		})
	}

	if keyb.IsKeySPressed(Keyboard.Enter) {
		core.Click(x, y)
	}

	if keyb.IsKeySPressed(Keyboard.CtrlV) {
		core.CallEventOn(x, y, func(c Core.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					c.Paste(keyb.GetClickboard())
					core.AllineCursor()
				}
			}
		})
	}

	if keyb.IsKeySPressed(Keyboard.CtrlA) {
		core.CallEventOn(x, y, func(c Core.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					c.SetWrap(!c.IsInSelectingMode())
				}
			}
		})
	}

	if keyb.IsKeySPressed(Keyboard.CtrlC) {
		core.CallEventOn(x, y, func(c Core.IComponent) {
			if c, ok := c.(*Component.TextBox); ok {
				if c.IsTyping() {
					keyb.InsertClickboard(c.GetSelectedText())
				}
			}
		})
	}

	if keyb.IsKeySPressed(Keyboard.CtrlQ) {
		return false
	}
	return true
}
