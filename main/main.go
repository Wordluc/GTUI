package main

import (
	Core "GTUI"
	"GTUI/Core/Component"
	"GTUI/Core/Utils/Color"
	"GTUI/Keyboard/impl"
	"time"
)

var x, y = 0, 0
var core, _ = Core.NewGtui()

var keyb = impl.Keyboard{}

func main() {
	core.ISetGlobalColor(Color.GetDefaultColor())
	defer core.Close()
	b := Component.CreateButton(5, 5, 10, 5, "prova")
	c := Component.CreateButton(50, 5, 11, 5, "provissima")
	keyb.Start()
	core.InsertComponent(b)
	core.InsertComponent(c)
	core.IRefreshAll()
	loop()
	defer keyb.Stop()
}

func loop() {
	command, _ := keyb.GetKey()
	core.IRefreshAll()
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
	case "q":
		return
	case "c":
		core.Interact(x, y)
		break
	}
	core.ISetCursor(x, y)
	time.Sleep(time.Millisecond * 50)
	loop()
}
