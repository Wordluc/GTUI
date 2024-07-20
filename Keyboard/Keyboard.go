package Keyboard

import (

)
type Token struct {
	
}
type Loop func() bool
type IKeyBoard interface {
	Start(loop Loop) error
	Stop()
	GetKey() (byte, error)
	IsKeyPressed(key byte) bool
	IsTokenPressed(token Token) bool
}
