package ITerminal

type ITerminal interface {
	Start() error
	Stop()
	Clear()
	Print(byte []byte)
	PrintStr(str string)
	Size() (int,int)
	GetCursor() (int, int)
	SetCursor(x, y int)
}
