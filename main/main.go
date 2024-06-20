package main

import (
	Core "GTUI"
	Color "GTUI/internal/Color"
	"os"
)

func main() {
	command := make([]byte, 1)
	core, _ := Core.NewCore()
	defer core.Close()	
	for {
		os.Stdin.Read(command)
		switch string(command) {
		case "a":
			core.Print("cioa")
		case "z":
			core.Printf("Hello World", Core.MetaData{Color: Color.Color{Foreground: Color.RedF, Background: Color.BlueB}})
		case "s":
			core.SetColor(Color.Color{Foreground: Color.CyanF, Background: Color.BlueB})
		case "r":
			core.ResetColor()
		case "p":
			core.SetCursor(10,5)
		default:
			return
		}
	}
}
