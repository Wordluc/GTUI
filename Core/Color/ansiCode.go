package Color


type iColor int8
const (//Foreground
	Black iColor = 30
	Red        = iota + 30
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Reset = 0
	None=-1
)

