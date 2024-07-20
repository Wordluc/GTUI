package main

import (
	Core "GTUI"
	"GTUI/Core/Color"
	"GTUI/Core/Drawing"
	"os"
)

func main() {
	command := make([]byte, 1)
	core, _ := Core.NewGtui()
	core.ISetGlobalColor(Color.GetDefaultColor())
	defer core.Close()

	rec := Drawing.CreateRectangle( 3, 3, 20, 10)
	textBox := Drawing.CreateTextBox( 4, 4)
	textBox1 := Drawing.CreateTextBox( 4, 5)
	textBox1.Type("hello")
	cont := Drawing.CreateContainer( 0, 0)
	cont.SetColor(Color.Get(Color.Cyan, Color.None))
	cont.AddChild(rec)
	cont.AddChild(textBox)
	cont.AddChild(textBox1)

	contHead := Drawing.CreateContainer(0, 0)
	rectHead := Drawing.CreateRectangleFull( 1, 1, 20, 10)
	rectHead.SetInsideColor(Color.Green)
	rectHead.GetBorder().SetColor(Color.Get(Color.Red, Color.Red))
	contHead.AddChild(rectHead)
	contHead.AddChild(cont)
	core.InsertEntity(contHead)
	t:=Drawing.CreateTextBox(10, 30)
	t.Type("hecwjckljklj")

	core.InsertEntity(t)
	for {
		core.IRefreshAll()
		os.Stdin.Read(command)
		if string(command) == "c" {
			textBox.SetColor(Color.Get(Color.Red, Color.None))
			continue
		} else if string(command) == "C" {
			textBox.SetColor(Color.GetNoneColor())
			continue
		}
		if string(command) == "q" {
			return
		}
		if string(command) == "a" {
			core.ISetGlobalColor(Color.Get(Color.Cyan, Color.Blue))
			continue
		}
		if string(command) == "A" {
			core.IResetGlobalColor()
			continue
		}
		if string(command) == "p" {
			rectHead.SetInsideColor(Color.Blue)
			continue
		}
		if string(command) == "h" {
			textBox.SetVisibility(!textBox.GetVisibility())
			core.IClear()
			continue
		}
		if string(command) == "P" {
			rectHead.SetInsideColor(Color.Red)
			rectHead.GetBorder().SetColor(Color.Get(Color.Green, Color.None))
			continue
		}
		if string(command) == "l" {
			Xpos, Ypos := cont.GetPos()
			cont.SetPos(Xpos+1, Ypos)
			core.IClear()
			continue
		}
		if string(command) == "k" {
			Xpos, Ypos := cont.GetPos()
			cont.SetPos(Xpos, Ypos+1)
			core.IClear()
			continue
		}
		if string(command) == "L" {
			Xpos, Ypos := contHead.GetPos()
			contHead.SetPos(Xpos+1, Ypos)
			core.IClear()
			continue
		}
		if string(command) == "K" {
			Xpos, Ypos := contHead.GetPos()
			contHead.SetPos(Xpos, Ypos+1)
			core.IClear()
			continue
		}
		textBox.Type(string(command))
	}
}
