package Color


type ColorValue int8
//for the backgroud will be used a function to convert it,gBg
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

