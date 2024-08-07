package Color


type ColorValue int8
const (//Foreground
	Black ColorValue = 30
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

