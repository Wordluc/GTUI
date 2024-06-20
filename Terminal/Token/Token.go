package Token

type Token int

const None Token=0
const (
	Arrow_Left  Token = 1
	Arrow_Right       = 2
	Arrow_Up          = 3
	Arrow_Down        = 4
)

func (t Token) String() string {
	return [...]string{
		"Arrow_Left",
		"Arrow_Right",
		"Arrow_Up",
		"Arrow_Down",
	}[t]
}
