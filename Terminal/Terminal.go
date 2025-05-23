package Terminal

import (
	"os"
	"strconv"

	"golang.org/x/term"
)

type Terminal struct {
	term       *term.State
	oldW, oldH int
}

func (t *Terminal) Start() error {
	tt, e := term.MakeRaw(int(os.Stdin.Fd()))
	if e != nil {
		return e
	}
	t.oldW, t.oldW = t.Size()
	t.term = tt
	term.NewTerminal(os.Stdin, "")

	t.Clear()
	return nil
}

func (t *Terminal) Stop() {
	t.Clear()
	t.PrintStr(t.ShowCursor())
	term.Restore(int(os.Stdin.Fd()), t.term)
}
func (t *Terminal) Resized() bool {
	w, h := t.Size()
	if w != t.oldW || h != t.oldH {
		t.oldW, t.oldH = w, h
		return true
	}
	return false
}
func (t *Terminal) Clear() {
	t.Print([]byte("\033[0J"))
	t.Print([]byte("\033[1J"))
}

func (t *Terminal) Print(byte []byte) {
	os.Stdout.Write(byte)
}

func (t *Terminal) PrintStr(str string) {
	os.Stdout.WriteString(str)
}

func (t *Terminal) Size() (int, int) {
	var e error
	x, y, e := term.GetSize(int(os.Stdout.Fd()))
	if e != nil {
		panic(e)
	}
	return x, y
}
func (t *Terminal) HideCursor() string {
	return "\033[?25l"
}
func (t *Terminal) ShowCursor() string {
	return "\033[?25h"
}
func (t *Terminal) SetCursor(x, y int) {
	t.PrintStr("\033[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H")
}
