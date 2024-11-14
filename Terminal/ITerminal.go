package Terminal

type ITerminal interface {
	Start() error
	Stop()
	Clear()
	Print(byte []byte)
	PrintStr(str string)
	Size() (int, int)
	SetCursor(x, y int)
	HideCursor() string
	ShowCursor() string
}
