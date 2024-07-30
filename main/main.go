package main

import (
	Core "GTUI"
	"GTUI/Core/Component"
	"GTUI/Core/Utils/Color"
	"os"
)

func main() {
	command := make([]byte, 1)
	core, _ := Core.NewGtui()
	core.ISetGlobalColor(Color.GetDefaultColor())
	defer core.Close()
	b := Component.CreateButton(1, 1, 10, 5, "prova")
	core.InsertComponent(b)
	x, y := 0, 0
			core.ISetCursor(x, y)
	for {
		core.IRefreshAll()
		os.Stdin.Read(command)
		core.ISetCursor(x, y)
		if string(command) == "c" {
			core.Interact(x, y)
			continue
		}
		switch string(command) {
		case "w":
			y--
			break
		case "a":
			x--
			break
		case "s":
			y++
			break
		case "d":
			x++
			break
		}
		if string(command) == "q" {
			break
		}
		core.ISetCursor(x, y)
	}
}
