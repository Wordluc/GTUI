package Core


type IEntity interface {
	Touch()
	GetAnsiCode() string
	GetName() string
	SetPos(x, y int)
	GetPos() (int, int)
}
