package main

import (
	"GTUI"
	Core "GTUI"
	"GTUI/Core/Component"
	"GTUI/Core/Utils/Color"
	"GTUI/Keyboard"
	"GTUI/Keyboard/impl"
)

var x, y = 0, 0
var core *Core.Gtui 

var keyb Keyboard.IKeyBoard = &impl.Keyboard{}

func Creation(k Keyboard.IKeyBoard) Component.StreamCharacter {
	stream := Component.StreamCharacter{}
	stream.Get = func() chan string {
		i := k.NewChannel()
		stream.IChannel = i
		return k.GetChannels()[i]
	}
	stream.Delete = func() {
		 k.DeleteChannel(stream.IChannel) 
	}
	return stream
}
func main() {
	core,_=Core.NewGtui()
	defer keyb.Stop()
	defer core.Close()
	core.ISetGlobalColor(Color.GetDefaultColor())
	b := Component.CreateButton(5, 5, 10, 5, "prova")
	c := Component.CreateTextField(50, 5, 10, 5, Creation(keyb))
	core.InsertComponent(c)
	core.InsertComponent(b)
	core.IRefreshAll()
	core.ISetCursor(x, y)
	keyb.Start(loop)
}

func loop(keyb Keyboard.IKeyBoard) bool{
	command, _ := keyb.GetKey()
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
		return false
	case "c":
		core.Interact(x, y, GTUI.Click)
		break
	}
	core.ISetCursor(x, y)
	core.IRefreshAll()
	core.ISetCursor(x, y)
	return true
}
