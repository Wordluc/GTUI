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
	rectHead.Color= Color.Color{Foreground: Color.RedF, Background: Color.BlueB}
	contHead.AddChild(rectHead)
	contHead.AddChild(cont)
	core.InsertEntity(contHead)
	core.IRefreshAll()
	for {
		core.IRefreshAll()
		os.Stdin.Read(command)
		if string(command) == "c" {
			textBox.Color = Color.Color{Foreground: Color.CyanF, Background: Color.BlueB}
			textBox.Touch()
			continue
		} else if string(command) == "C" {
			textBox.Color = Color.GetDefaultColor()
			textBox.Touch()
			continue
		}
		if string(command) == "q" {
			return
		}
		if string(command) == "l" {
			cont.SetPos(cont.XPos+1, cont.YPos)
			cont.Touch()
			core.IClear()
			continue
		}
		if string(command) == "k" {
			cont.SetPos(cont.XPos, cont.YPos+1)
			cont.Touch()
			core.IClear()
			continue
		}
		if string(command) == "L" {
			contHead.SetPos(contHead.XPos+1, contHead.YPos)
			contHead.Touch()
			core.IClear()
			continue
		}
		if string(command) == "K" {
			contHead.SetPos(contHead.XPos, contHead.YPos+1)
			contHead.Touch()
			core.IClear()
			continue
		}
		textBox.Type(string(command))
	}
}
