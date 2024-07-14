package Color


type iColor int8
const (
	BlackF iColor = 30
	RedF        = iota + 30
	GreenF
	YellowF
	BlueF
	MagentaF
	CyanF
	WhiteF
	ResetF = 0
)

const (
	BlackB iColor = 40
	RedB        = iota + 40
	GreenB
	YellowB
	BlueB
	MagentaB
	CyanB
	WhiteB
	ResetB = 0
)
