package impl

import (
	"os"
	"regexp"
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
	os.Stdout.Write([]byte("\033[H\033[2J"))
}

func (t *Terminal) Print(byte []byte) {
	os.Stdout.Write(byte)
}

func (t *Terminal) PrintStr(str string) {
	os.Stdout.Write([]byte(str))
}

func (t *Terminal) Len() (int,int) {
	var e error
	x,y,e := term.GetSize(int(os.Stdout.Fd()))
	if e != nil {
		panic(e)
	}
	return x,y
}

func (t *Terminal) GetCursor() (int, int) {
	t.PrintStr("\033[6n")
	pos:= make([]byte, 32)
	os.Stdin.Read(pos)
	regex:=regexp.MustCompile("\\[?([0-9]+);([0-9]+)R")
	a:=regex.FindAllSubmatch(pos, -1)
	if len(a) == 0 || len(a[0]) != 3 {
		return t.GetCursor()
	}
	x,_:=strconv.Atoi(string(a[0][2]))
	y,_:=strconv.Atoi(string(a[0][1]))
	return x,y
}
