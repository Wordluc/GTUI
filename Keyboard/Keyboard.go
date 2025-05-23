package Keyboard

import (
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/atotto/clipboard"
)

type Keyboard struct {
	key       stateKey
	IsRunning bool
	channel   []chan string
	loop      Loop
}
type stateKey struct {
	key  Key
	rune rune
}

func NewKeyboard() *Keyboard {
	return &Keyboard{channel: make([]chan string, 0), IsRunning: false}
}
func (t *Keyboard) Start(Loop Loop) error {
	t.IsRunning = true
	t.loop = Loop

	keyboard.Listen(t.keyListening)
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
func (t *Keyboard) DeleteChannel(i int) {
	if i <= 0 {
		t.channel = []chan string{}
		return
	}
	t.channel = append(t.channel[:i], t.channel[i+1:]...)
}

func (t *Keyboard) keyListening(v keys.Key) (bool, error) {
	if v.Code == keys.Space {
		v.Runes =[]rune{' '}
	}
	if v.Code == keys.Tab {
		v.Runes = []rune{'	'}
	}
	if v.Code == keys.Enter {
		v.Runes = []rune{'\n'}
	}
	if v.Code == keys.Backspace {
		v.Runes = []rune{'\b'}
	}
	if rune:=v.Runes;len(rune)>0{
		t.key = stateKey{key: Key(v.Code), rune: rune[0]}
	}else
	{
		t.key = stateKey{key: Key(v.Code), rune: 0}
	}
	for _, b := range t.GetChannels() {
		key, e := t.GetKey()
		if e != nil {
			continue
		}
		if key == 0 {
			break
		}
		b <- string(v.Runes)
		<-b
	}
	if !t.loop(t) {
		return true, nil
	}
	return false, nil
}
 
func (t *Keyboard) GetClickboard() string {
	text, e := clipboard.ReadAll()
	if e != nil {
		text = ""
	}
	return text
}
func (t *Keyboard) InsertClickboard(text string) {
	clipboard.WriteAll(text)
}
func (t *Keyboard) Stop() {
}

func (t *Keyboard) GetKey() (byte, error) {
	return byte(t.key.rune), nil
}

func (t *Keyboard) IsKeySPressed(key Key) bool {
	return t.key.key == key
}

func (t *Keyboard) IsKeyPressed(key rune) bool {
	return t.key.rune == key
}
//Clean the keybaord state
func (t *Keyboard) CleanKeyboardState() {
	t.key = stateKey{}
}
