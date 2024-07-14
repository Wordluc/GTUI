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
			core.IRefresh()
		case "c":
			core.InsertStrs("1234567890abcdefghijklmnopqrstuvwxyz\n")
			core.InsertStrs("1234567890abcdefghijklmnopqrstuvwxyz\n")
			core.InsertStrs("1234567890abcdefghijklmnopqrstuvwxyz\n")
			core.InsertStrs("1234567890abcdefghijklmnopqrstuvwxyz\n")
			core.InsertStrs("1234567890abcdefghijklmnopqrstuvwxyz\n")
			core.ClearPortion(4,3,1,1)
		case "l":
			  line:=Drawing.CreateLine("line",1,1,10,0)
				core.InsertEntity(line)
			
		case "p":
			  rectangle:=Drawing.CreateRectangle("rectangle",1,1,10,10)
				core.InsertEntity(rectangle)
		default:
			return
		}
	}
}
