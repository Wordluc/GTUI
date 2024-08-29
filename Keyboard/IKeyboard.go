package Keyboard

type Token struct{}
type Loop func(IKeyBoard) bool
type IKeyBoard interface {
	Start(Loop Loop) error
	Stop()
	GetKey() (byte, error)
	IsKeyPressed(key rune) bool
	IsKeySPressed(key Key) bool
	GetChannels() []chan string
	NewChannel() int
	DeleteChannel(int)
}
