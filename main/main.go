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
	defer core.Close()
	rec := Drawing.CreateRectangle("rec", 3, 3, 20, 10)
	textBox := Drawing.CreateTextBox("test", 4, 4)
	cont := Drawing.CreateContainer("testo", 0, 0)
	cont.AddChild(rec)
	cont.AddChild(textBox)

	contHead := Drawing.CreateContainer("head", 0, 0)
	rectHead := Drawing.CreateRectangle("testo", 0, 0, 10, 5)
	rectHead.SetColor(Color.Color{Foreground: Color.RedF, Background: Color.BlueB})
	contHead.AddChild(rectHead)
	contHead.AddChild(cont)
	core.InsertEntity(contHead)
	core.IRefreshAll()
	for {
		core.IRefreshAll()
		os.Stdin.Read(command)
		if string(command) == "c" {
			textBox.SetColor(Color.Color{Foreground: Color.CyanF, Background: Color.BlueB})
			continue
		} else if string(command) == "C" {
			textBox.SetColor(Color.GetDefaultColor())
			continue
		}
		if string(command) == "q" {
			return
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
