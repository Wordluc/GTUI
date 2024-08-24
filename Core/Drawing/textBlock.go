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
	getText         string
}

func CreateLineText(initialCapacity int) *lineText {
	return &lineText{
		line:            make([]rune, initialCapacity),
		initialCapacity: initialCapacity,
		totalChar:       0,
		isChanged:       true, //da usare
	}
}
func (t *lineText) getFullText() string {
	var str strings.Builder
	for i := range t.totalChar {
		str.WriteRune(t.line[i])
	}
	return str.String()
}

func (t *lineText) digit(char rune) {
	t.line[t.totalChar] = char
	t.totalChar++
	if t.totalChar >= len(t.line) {
		t.line = slices.Concat(t.line, make([]rune, t.initialCapacity))
	}
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
		totalLine:       0,
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

func (t *TextBlock) SetCurrentLine(line int) {
	if line >= len(t.lines) {
		return
	}
	t.currentLine = line
	if t.lines[t.currentLine] == nil {
		t.lines[t.currentLine] = CreateLineText(t.initialCapacity)
	}
}

func (t *TextBlock) SetCurrentChar(ichar int) {
	t.currentCharacter = ichar
}
func (t *TextBlock) SetCurrent(ichar int, line int) {
	isYChanged := line != t.currentLine
	isXChanged := line == t.currentLine && ichar != t.currentCharacter
	if !isYChanged && !isXChanged {
		t.preLenght= ichar
	}
	if isXChanged {
		t.preLenght = t.currentCharacter
		t.SetCurrentChar(ichar)
	}		
	if isYChanged {
		t.SetCurrentLine(line)
		maxX, _ := t.GetTotalCursor()
		if t.preLenght > maxX {
			t.SetCurrentChar(maxX)
		} else {
			t.SetCurrentChar(t.preLenght)
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
		t.lines[t.currentLine].digit(char)
		t.currentCharacter++
	}
	t.Touch()
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
		full.WriteString(U.GetAnsiMoveTo(t.xPos, t.yPos+i))
		if line == nil {
			break
		}
		full.WriteString(line.getFullText())
	}
	return full.String()
}
