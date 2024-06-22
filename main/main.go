package main

import (
	Core "GTUI"
	Color "GTUI/internal/Color"
	Figures "GTUI/internal/Drawing"
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
			core.InsertStrs("cioa\n")
		case "z":
			core.InsertStrsf(Core.MetaData{Color: Color.Color{Foreground: Color.RedF, Background: Color.BlueB}},"Hello World", )
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
			if str,e:=Figures.GetAnsiLine(Figures.Line{Angle: 90, Len: 10, XPos: 10, YPos: 1});e==nil{
				core.InsertStrs(str)
			}
		case "p":
			if str,e:=Figures.GetAnsiRectangle(Figures.Rectangle{Width: 10, Height: 10, XPos: 1, YPos: 1});e==nil{
				core.InsertStrs(str)
			}
		default:
			return
		}
	}
}
