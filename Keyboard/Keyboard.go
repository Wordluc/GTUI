package Keyboard

type Token struct{}
type Loop func(IKeyBoard)bool
type IKeyBoard interface {
	Start(Loop) error
	Stop()
	GetKey() (byte, error)
	IsKeyPressed(key byte) bool
	TokenPressed(token Token) bool
	GetChannels() []chan string
	NewChannel() int
	DeleteChannel(int)
}
