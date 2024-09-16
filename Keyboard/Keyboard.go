package Keyboard

import (
	"errors"

	"github.com/eiannone/keyboard"
)

type Keyboard struct {
	key       stateKey
	isRunning bool
	channel   []chan string
	loop      Loop
}

type stateKey struct {
	key  Key
	rune rune
}

func NewKeyboard() *Keyboard {
	return &Keyboard{ channel: make([]chan string, 0), isRunning: false}
}

func (t *Keyboard) Start(Loop Loop) error {
	t.isRunning = true
	t.loop = Loop
	eventKey, e := keyboard.GetKeys(1)
	if e != nil {
		return e
	}
	t.keyListening(eventKey)
	return nil
}

func (t *Keyboard) GetChannels() []chan string {
	return t.channel
}

func (t *Keyboard) NewChannel() int {
	channel := make(chan string)
	t.channel = append(t.channel, channel)
	i := len(t.channel) - 1
	return i
}

func (t *Keyboard) DeleteChannel(i int) error {
	if i == 0 {
		t.channel = []chan string{}
		return nil
	}
	if i<0 || i >= len(t.channel) {
		return errors.New("index out of range")
	}
	t.channel = append(t.channel[:i], t.channel[i+1:]...)
	return nil
}

func (t *Keyboard) keyListening(eventKey <-chan keyboard.KeyEvent) {
	for {
		v := <-eventKey
		if v.Key == keyboard.KeySpace {
			v.Rune = ' '
		}
		if v.Key == keyboard.KeyEnter {
			v.Rune = '\n'
		}
		if v.Key == keyboard.KeyBackspace {
			v.Rune = '\b'
		}
		t.key = stateKey{key: Key(v.Key), rune: v.Rune}
		for _, b := range t.GetChannels() {
			key, e := t.GetKey()
			if e != nil {
				continue
			}
			if key == 0 {
				break
			}
			b <- string(v.Rune)
		}
		if !t.loop(t) {
			break
		}
	}
}
func (t *Keyboard) Stop() {
	keyboard.Close()
}

func (t *Keyboard) GetKey() (byte, error) {
	return byte(t.key.rune), nil
}

func (t *Keyboard) IsSpecialKeyPressed(key Key) bool {
	return t.key.key == key
}

func (t *Keyboard) IsKeyPressed(key rune) bool {
	return t.key.rune == key
}
