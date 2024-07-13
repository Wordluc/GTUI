package ITerminal

type ITerminal interface {
	Start() error
	Stop()
	Clear()
	Print(byte []byte)
	PrintStr(str string)
	Len() (int,int)
	GetCursor() (int, int)
	SetCursor(x, y int)
}
