package Keyboard

import (
	"github.com/eiannone/keyboard"
)

type Keyboard struct {
	key       stateKey
	IsRunning bool
	Channel   []chan string
	Loop      func(IKeyBoard) bool
}
type stateKey struct {
	key  Key
	rune rune
}

func (t *Keyboard) Start(Loop Loop) error {
	t.Channel = make([]chan string, 0)
	t.IsRunning = true
	t.Loop = Loop
	eventKey, e := keyboard.GetKeys(1)
	if e != nil {
		return e
	}
	t.keyListening(eventKey)
	return nil
}
func (t *Keyboard) GetChannels() []chan string {
	return t.Channel
}
func (t *Keyboard) NewChannel() int {
	channel := make(chan string)
	t.Channel = append(t.Channel, channel)
	i := len(t.Channel) - 1
	return i
}
func (t *Keyboard) DeleteChannel(i int) {
	if i <= 0 {
		t.Channel = []chan string{}
		return
	}
	t.Channel = append(t.Channel[:i], t.Channel[i+1:]...)
}

func (t *Keyboard) keyListening(eventKey <-chan keyboard.KeyEvent) {
	for {
		v := <-eventKey
		if v.Key == keyboard.KeySpace {
			v.Rune = ' '
		}
		t.key = stateKey{key: Key(v.Key), rune: v.Rune}
		for _, b := range t.GetChannels() {
			key, e := t.GetKey()
			if e != nil {
				continue
			}
			b <- string(key)
		}
		if !t.Loop(t) {
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

func (t *Keyboard) IsKeySPressed(key Key) bool {
	return t.key.key==key
}

func (t *Keyboard) IsKeyPressed(key rune) bool {
	return t.key.rune == key
}
