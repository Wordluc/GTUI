package main

import (
	"github.com/Wordluc/GTUI"
	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/Component"
	"github.com/Wordluc/GTUI/Core/Drawing"
	AlignCursor "github.com/Wordluc/GTUI/Core/EventManager"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
	"github.com/Wordluc/GTUI/Keyboard"
	Kd "github.com/Wordluc/GTUI/Keyboard"
)

var comp *Component.Container
var button2 *Component.Button
var modal *Component.Modal

func main() {
	core, _ := GTUI.NewGtui(loop)
	xS, yS := 50, 40
	oneLineTextbox, e := Component.CreateTextBox(40, 38, 20, 3, core.CreateStreamingCharacter())
	oneLineTextbox.IsOneLine = true
	core.AddComponent(oneLineTextbox)
	c, e := Component.CreateTextBox(30, 5, 20, 10, core.CreateStreamingCharacter())
	c.SetLayer(Core.L3)
	if e != nil {
		panic(e)
	}
	rect := Drawing.CreateRectangle(2, 2, xS-1, yS-1)
	rect1 := Drawing.CreateRectangle(16, 26, 10, 10)
	contDraw := Drawing.CreateContainer()
	contDraw.AddDrawings(rect)
	contDraw.AddDrawings(rect1)
	contDraw.SetVisibility(true)
	compComponent := Component.CreateContainer()
	compComponent.AddDrawing(contDraw.GetGraphics()...)
	if e := compComponent.AddComponent(c); e != nil {
		panic(e)
	}
	comp = compComponent
	button1 := Component.CreateButton(100, 0, 10, 10, "test1")
	button1.SetActive(false)
	button1.SetLayer(Core.L4)
	button1.GetVisibleArea().SetInsideColor(Color.White)
	button2 = Component.CreateButton(62, 15, 10, 10, "test")
	button2.SetLayer(Core.L1)
	button2.GetVisibleArea().SetInsideColor(Color.Blue)
	comp.SetPos(10, 10)
	if e := core.AddContainer(compComponent); e != nil {
		panic(e)
	}
	if e := core.AddComponent(button1); e != nil {
		panic(e)
	}
	if e := core.AddComponent(button2); e != nil {
		panic(e)
	}
	if e := core.SetCur(1, 1); e != nil {
		panic(e)
	}

	text := Drawing.CreateTextField(10, 1, "modal")
	ok := Component.CreateButton(2, 5, 10, 3, "ok")
	esplodi := Component.CreateButton(18, 5, 10, 3, "esplodi")

	modal = Component.CreateModal(30, 10) //the buttons dont work inside the modal
	modal.SetBackgroundColor(Color.Gray)
	modal.AddDrawing(text)
	modal.AddComponent(esplodi)
	modal.AddComponent(ok)
	modal.SetPos(60, 10)
	if e := core.AddComplexElement(modal); e != nil {
		panic(e)
	}
	core.Start()
}

func loop(keyb Kd.IKeyBoard, core *GTUI.Gtui) bool {
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
	core.SetCur(x, y)
	if keyb.IsKeyPressed('S') {
		core.GoTo(Core.Down)
	}
	if keyb.IsKeyPressed('W') {
		core.GoTo(Core.Up)
	}
	if keyb.IsKeyPressed('D') {
		core.GoTo(Core.Right)
	}
	if keyb.IsKeyPressed('A') {
		core.GoTo(Core.Left)
	}
	if keyb.IsKeySPressed(Keyboard.CtrlS) {
		button2.SetPos(x, y)
	}

	if keyb.IsKeySPressed(Keyboard.CtrlP) {
		modal.SetPos(x, y)
	}

	if keyb.IsKeyPressed('v') {
		modal.SetActive(!modal.GetVisibility())
		modal.SetVisibility(!modal.GetVisibility())
	}
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
					AlignCursor.Call(AlignCursor.CursorAlign)
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
