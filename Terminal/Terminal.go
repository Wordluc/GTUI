package Terminal

import (
	"os"
	"strconv"
	"golang.org/x/term"
)

type Terminal struct {
	term *term.State
}

func (t *Terminal) Start() error {
	tt, e := term.MakeRaw(int(os.Stdin.Fd()))
	if e != nil {
		return e
	}
	t.term = tt
	term.NewTerminal(os.Stdin, "")
	t.Clear()
	return nil
}

func (t *Terminal) Stop() {
	term.Restore(int(os.Stdin.Fd()), t.term)
}

func (t *Terminal) Clear() {
	t.Print([]byte("\033[0J"))
	t.Print([]byte("\033[1J"))
}


func (t *Terminal) Print(byte []byte) {
	os.Stdout.Write(byte)
}

func (t *Terminal) PrintStr(str string) {
	os.Stdout.Write([]byte(str))
}

func (t *Terminal) Size() (int,int) {
	var e error
	x,y,e := term.GetSize(int(os.Stdout.Fd()))
	if e != nil {
		panic(e)
	}
	return x,y
}
func (t *Terminal) HideCursor() {
	t.PrintStr("\033[?25l")
}
func (t *Terminal) ShowCursor() {
	t.PrintStr("\033[?25h")
}
func (t *Terminal) SetCursor(x, y int) {
	t.PrintStr("\033[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H")
}
func (t *Terminal) GetSetCursor(x, y int) string {
	return "\033[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H"
}
