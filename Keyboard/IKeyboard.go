package Keyboard

type Token struct{}
type Loop func(IKeyBoard) bool
type IKeyBoard interface {
	Start(Loop Loop) error
	Stop()
	GetKey() (byte, error)
	// return true if the character is pressed
	IsKeyPressed(key rune) bool
	// return true if the key is pressed e.g: Space, Enter, Backspace, etc.
	IsKeySPressed(key Key) bool
	GetChannels() []chan string
	NewChannel() int
	DeleteChannel(int)
	GetClickboard() string
}
