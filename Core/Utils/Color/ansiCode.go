package Color


type IColor int8
const (//Foreground
	Black IColor = 30
	Red        = iota + 30
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Gray=90
	Reset = 0
	None=-1
)

