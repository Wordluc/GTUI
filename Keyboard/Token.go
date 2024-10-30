package Keyboard

import "atomicgo.dev/keyboard/keys"

type Key int

const (
	Null      Key =Key(keys.Null)
	Break     Key =Key(keys.Break)
	Enter     Key =Key(keys.Enter)
	Backspace Key =Key(keys.Backspace)
	Tab       Key =Key(keys.Tab)
	Esc       Key =Key(keys.Esc)
	Escape    Key =Key(keys.Escape)

	CtrlAt Key = Key(keys.CtrlAt)
	CtrlA  Key = Key(keys.CtrlA)
	CtrlB  Key = Key(keys.CtrlB)
	CtrlC  Key = Key(keys.CtrlC)
	CtrlD  Key = Key(keys.CtrlD)
	CtrlE  Key = Key(keys.CtrlE)
	CtrlF  Key = Key(keys.CtrlF)
	CtrlG  Key = Key(keys.CtrlG)
	CtrlH  Key = Key(keys.CtrlH)
	CtrlI  Key = Key(keys.CtrlI)
	CtrlJ  Key = Key(keys.CtrlJ)
	CtrlK  Key = Key(keys.CtrlK)
	CtrlL  Key = Key(keys.CtrlL)
	CtrlM  Key = Key(keys.CtrlM)
	CtrlN  Key = Key(keys.CtrlN)
	CtrlO  Key = Key(keys.CtrlO)
	CtrlP  Key = Key(keys.CtrlP)
	CtrlQ  Key = Key(keys.CtrlQ)
	CtrlR  Key = Key(keys.CtrlR)
	CtrlS  Key = Key(keys.CtrlS)
	CtrlT  Key = Key(keys.CtrlT)
	CtrlU  Key = Key(keys.CtrlU)
	CtrlV  Key = Key(keys.CtrlV)
	CtrlW  Key = Key(keys.CtrlW)
	CtrlX  Key = Key(keys.CtrlX)
	CtrlY  Key = Key(keys.CtrlY)
	CtrlZ  Key = Key(keys.CtrlZ)

	CtrlOpenBracket  Key = Key(keys.Escape)
	CtrlBackslash    Key = Key(keys.CtrlBackslash)
	CtrlCaret        Key = Key(keys.CtrlCaret)
	CtrlUnderscore   Key = Key(keys.CtrlUnderscore)
	CtrlQuestionMark Key = Key(keys.CtrlQuestionMark)
)

// Other keys.
const (
	RuneKey Key = -(iota + 1)
	Up
	Down
	Right
	Left
	ShiftTab
	Home
	End
	PgUp
	PgDown
	Delete
	Space
	CtrlUp
	CtrlDown
	CtrlRight
	CtrlLeft
	ShiftUp
	ShiftDown
	ShiftRight
	ShiftLeft
	CtrlShiftUp
	CtrlShiftDown
	CtrlShiftLeft
	CtrlShiftRight
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
)
