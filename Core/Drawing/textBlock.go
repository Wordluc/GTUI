package Drawing

import (
	U "GTUI/Core/Utils"
	"GTUI/Core/Utils/Color"
	"slices"
	"strings"

)

type TextBlock struct {
	visible          bool
	ansiCode         string
	isChanged        bool
	lines            []*lineText
	xPos             int
	yPos             int
	currentCharacter int
	currentLine      int
	initialCapacity  int
	totalLine        int
	preLenght        int
}
type lineText struct {
	line            []rune
	initialCapacity int
	totalChar       int
	isChanged       bool
}

func CreateLineText(initialCapacity int) *lineText {
	return &lineText{
		line:            make([]rune, initialCapacity),
		initialCapacity: initialCapacity,
		totalChar:       0,
	}
}
func (t *lineText) getText() string {
	var str strings.Builder
	for i := range t.totalChar {
		str.WriteRune(t.line[i])
	}
	return str.String()
}

func (t *lineText) digit(char rune, i int) {
	if i > t.totalChar {
		t.line[i] = char
		t.totalChar = i
	} else {
		t.line = slices.Concat(t.line[:i], []rune{char}, t.line[i:])
		t.totalChar++
	}
	if t.totalChar >= len(t.line) {
		t.line = slices.Concat(t.line, make([]rune, t.initialCapacity))
	}
}

// /delete the character i, if the line is empty return true
func (t *lineText) delete(i int) bool {
	t.totalChar--
	if  t.totalChar<=0 {
		return true
	}
	if i <= 0 {
		return false
	}
	if i > t.totalChar {
		return false
	}
	t.line = slices.Concat(t.line[:i-1], t.line[i:])
	if t.totalChar <= 0 {
		t.totalChar = 0
	}
	return false
}
func CreateTextBlock(x, y int, initialLine int, initialCapacity int) *TextBlock {
	lines := make([]*lineText, initialLine)
	lines[0] = CreateLineText(initialCapacity)
	return &TextBlock{
		visible:         true,
		ansiCode:        "",
		isChanged:       true,
		lines:           lines,
		xPos:            x,
		yPos:            y,
		currentLine:     0,
		initialCapacity: initialCapacity,
		totalLine:       1,
	}
}

func (t *TextBlock) Touch() {
	t.isChanged = true
}
func (t *TextBlock) GetAnsiCode(defaultColor Color.Color) string {
	if t.isChanged {
		t.ansiCode = t.getFullText()
		t.isChanged = false
	}
	return t.ansiCode
}
func (t *TextBlock) SetPos(x, y int) {
	t.xPos = x
	t.yPos = y
	t.Touch()
}

func (t *TextBlock) ForceSetCurrentLine(line int) {
	if line >= len(t.lines) {
		return
	}
	t.currentLine = line
	if t.lines[t.currentLine] == nil {
		t.lines[t.currentLine] = CreateLineText(t.initialCapacity)
	}
}

func (t *TextBlock) ForceSetCurrentChar(ichar int) {
	t.currentCharacter = ichar
}

func (t *TextBlock) SetCurrentCursor(ichar int, line int) {
	isYChanged := line != t.currentLine
	isXChanged := line == t.currentLine && ichar != t.currentCharacter
	if isXChanged {
		if ichar > t.currentCharacter {
			t.preLenght = t.currentCharacter + 1
		} else {
			t.preLenght = t.currentCharacter - 1
		}
		t.ForceSetCurrentChar(ichar)
	}
	if isYChanged {
		t.ForceSetCurrentLine(line)
		maxX, _ := t.GetTotalCursor()
		if t.preLenght > maxX {
			t.ForceSetCurrentChar(maxX)
		} else {
			t.ForceSetCurrentChar(t.preLenght)
		}
	}
}

func (t *TextBlock) GetTotalCursor() (int, int) {
	return t.lines[t.currentLine].totalChar, t.totalLine
}
func (t *TextBlock) GetCurrentCursor() (int, int) {
	return t.currentCharacter, t.currentLine
}
func (t *TextBlock) Type(char rune) {
	if char == '\n' {
		t.totalLine++
		t.currentLine++
		t.currentCharacter = 0
		if t.currentLine >= len(t.lines) {
			t.lines = append(t.lines, CreateLineText(t.initialCapacity))
		}
		if t.lines[t.currentLine] == nil {
			t.lines[t.currentLine] = CreateLineText(t.initialCapacity)
		}
	} else {
		t.lines[t.currentLine].digit(char, t.currentCharacter)
		t.currentCharacter++
		t.preLenght = t.currentCharacter
	}
	t.Touch()
}
func (t *TextBlock) Delete() {
	if t.currentCharacter == 0 && t.currentLine == 0 {
		return
	}
	deleteLine := t.lines[t.currentLine].delete(t.currentCharacter)
	defer t.Touch()
	t.preLenght = t.currentCharacter - 1
	t.currentCharacter--
	if deleteLine || t.currentCharacter < 0 {
		t.totalLine--
		t.lines = slices.Concat(t.lines[:t.currentLine], t.lines[t.currentLine+1:])
		t.currentLine--
		if t.totalLine==0{
      t.lines = []*lineText{CreateLineText(t.initialCapacity)}
			t.currentLine = 0
			t.totalLine = 1
			return
		}
		t.currentCharacter = t.lines[t.currentLine].totalChar
	}
}
func (t *TextBlock) GetPos() (int, int) {
	return t.xPos, t.yPos
}
func (t *TextBlock) SetVisibility(visible bool) {
	t.visible = visible
}
func (t *TextBlock) GetVisibility() bool {
	return t.visible
}

// todo da inserire i colori
func (t *TextBlock) getFullText() string {
	full := strings.Builder{}
	for i, line := range t.lines {
		if line == nil {
			break
		}
		full.WriteString(U.GetAnsiMoveTo(t.xPos, t.yPos+i))
		full.WriteString(line.getText())
	}
	return full.String()
}
