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
	for {
		os.Stdin.Read(command)
		switch string(command) {
		case "s":
			core.ISetGlobalColor(Color.Color{Foreground: Color.CyanF, Background: Color.BlueB})
		case "r":
			core.IResetGlobalColor()
		case "t":
			core.IRefreshAll()
		case "c":
			textBox := Drawing.CreateTextBox("textbox", 0, 0)
			textBox.Type("1234567890abcdefghijklmnopqrstuvwxyz\n")
			textBox.Type("1234567890abcdefghijklmnopqrstuvwxyz\n")
			core.InsertEntity(textBox)
		case "l":
			line := Drawing.CreateLine("line", 1, 10, 10, 0)
			core.InsertEntity(line)
		case "p":
			rectangle := Drawing.CreateRectangle("rectangle", 1, 1, 10, 10)
			core.InsertEntity(rectangle)
		default:
			return
		}
	}
}
