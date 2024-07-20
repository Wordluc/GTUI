package impl

import (
	IKeyboard "GTUI/Keyboard"

	"github.com/eiannone/keyboard"
)

type Keyboard struct {
	key  stateKey
	IsRunning bool
}
type stateKey struct {
	key  keyboard.Key
	rune rune
}

func (t *Keyboard) Start(loop IKeyboard.Loop) error {
	keyboard.Open()
	eventKey, e := keyboard.GetKeys(10)
	if e != nil {
		return e
	}
	for {
		v := <-eventKey
		t.key = stateKey{key: v.Key, rune: v.Rune}
		if !t.IsRunning{
       break
		}
	}
	return nil
}

func (t *Keyboard) Stop() {
	keyboard.Close()
}

func (t *Keyboard) GetKey() (byte, error) {
	return byte(t.key.rune), nil
}

func (t *Keyboard) IsTokenPressed(token IKeyboard.Token) bool {
	key := t.key.key
	if v, e := mapTokenToKey(key); e == nil {
		if v == key {
			return true
		}
	}
	return false
}

func (t *Keyboard) IsKeyPressed(key byte) bool {
	return byte(t.key.rune)==key 
}

func mapTokenToKey(token keyboard.Key) (keyboard.Key, error) {
	switch token {
	case keyboard.KeyArrowLeft:
		return keyboard.KeyArrowLeft, nil
	case keyboard.KeyArrowRight:
		return keyboard.KeyArrowRight, nil
	case keyboard.KeyArrowUp:
		return keyboard.KeyArrowUp, nil
	case keyboard.KeyArrowDown:
		return keyboard.KeyArrowDown, nil
	}
	return 0, nil
}
