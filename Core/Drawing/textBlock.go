package Drawing

import (
	U "GTUI/Core/Utils"
	"GTUI/Core/Utils/Color"
	"slices"
	"strings"
)

type TextBlock struct {
	visible           bool
	ansiCode          string
	isChanged         bool
	lines             []*lineText
	xPos              int
	yPos              int
	xSize             int
	ySize             int
	currentCharacter  int
	currentOffsetChar int
	currentLine       int
	initialCapacity   int
	totalLine         int
	preLenght         int
}
type lineText struct {
	line            []rune
	initialCapacity int
	totalChar       int
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
	if t.totalChar <= 0 {
		return true
	}
	if i <= 0 {
		return false
	}
	t.totalChar--
	if i > t.totalChar {
		return false
	}
	t.line = slices.Concat(t.line[:i-1], t.line[i:])
	if t.totalChar <= 0 {
		t.totalChar = 0
	}
	return false
}

func (t *lineText) merge(add *lineText) {
	t.line = slices.Concat(t.line[:t.totalChar], add.line)
	t.totalChar += add.totalChar
}
func (t *lineText) split(i int) *lineText {
	if i >= t.totalChar {
		return CreateLineText(t.initialCapacity)
	}
	splited := t.line[i:t.totalChar]
	t.line = t.line[:i]
	newLine := CreateLineText(t.initialCapacity)
	newLine.totalChar = t.totalChar - i
	newLine.line = splited
	t.totalChar = i
	return newLine
}
func CreateTextBlock(x, y int, xSize, ySize int, initialCapacity int) *TextBlock {
	lines := make([]*lineText, 1)
	lines[0] = CreateLineText(initialCapacity)
	return &TextBlock{
		visible:         true,
		ansiCode:        "",
		isChanged:       true,
		lines:           lines,
		xPos:            x,
		yPos:            y,
		xSize:           xSize,
		ySize:           ySize,
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
		t.ansiCode = t.getTextWithAnsi()
		t.isChanged = false
	}
	return t.ansiCode
}
func (t *TextBlock) SetPos(x, y int) {
	t.xPos = x
	t.yPos = y
	t.Touch()
}

func (t *TextBlock) GetSize() (int, int) {
	return t.xSize, t.ySize
}
func (t *TextBlock) ForceSetCurrentCharacter(x int) {
	t.currentCharacter = x
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

func (t *TextBlock) SetCurrentCursor(x, y int) (int, int) {
	xRelative := x - t.xPos
	yRelative := y - t.yPos
	if yRelative < 0 {
		yRelative = 0
	}
	if yRelative >= t.ySize {
		return x, t.ySize
	}
	isYChanged := yRelative != t.currentLine
	isXChanged := xRelative != t.currentCharacter

	if yRelative >= len(t.lines) {
		yRelative = len(t.lines) - 1
	}
	if xRelative >= t.lines[yRelative].totalChar {
		xRelative = t.lines[yRelative].totalChar
	}

	if isXChanged {
		t.ForceSetCurrentCharacter(xRelative)
		t.preLenght = xRelative
	}
	if isYChanged {
		if t.lines[yRelative].totalChar > t.preLenght {
			t.ForceSetCurrentCharacter(t.preLenght)
		} else {
			t.ForceSetCurrentCharacter(t.lines[yRelative].totalChar)
		}
		t.ForceSetCurrentLine(yRelative)
	}
	return t.currentCharacter + t.xPos, yRelative + t.yPos

}

func (t *TextBlock) GetTotalCursor() (int, int) {
	return t.lines[t.currentLine].totalChar, t.totalLine
}
func (t *TextBlock) GetCurrentCursor() (int, int) {
	return t.currentCharacter - t.currentOffsetChar, t.currentLine
}
func (t *TextBlock) Type(char rune) {
	if char == '\n' {
		t.totalLine++
		t.currentLine++
		if t.currentLine >= len(t.lines) {
			t.lines = append(t.lines, CreateLineText(t.initialCapacity))
		}
		if t.currentCharacter > 0 && t.currentCharacter < t.lines[t.currentLine-1].totalChar {
			newLine := t.lines[t.currentLine-1].split(t.currentCharacter)
			t.lines = slices.Concat(t.lines[:t.currentLine], []*lineText{newLine}, t.lines[t.currentLine:])
		}
		t.ForceSetCurrentCharacter(0)
	} else {
		t.lines[t.currentLine].digit(char, t.currentCharacter)
		t.ForceSetCurrentCharacter(t.currentCharacter + 1)
		t.preLenght = t.currentCharacter
	}
	t.Touch()
}
func printArray(array []rune) {
	for i, v := range array {
		println(i, ":", string(v))
	}
}
func (t *TextBlock) Delete() {
	if t.totalLine == 1 && t.currentCharacter == 0 {
		return
	}
	defer t.Touch()
	if t.currentCharacter == 0 {
		if t.currentLine-1 >= 0 {
			t.ForceSetCurrentCharacter(t.lines[t.currentLine-1].totalChar)
			t.preLenght = t.currentCharacter
			t.lines[t.currentLine-1].merge(t.lines[t.currentLine])
			t.lines = slices.Concat(t.lines[:t.currentLine], t.lines[t.currentLine+1:])
			t.currentLine--
			t.totalLine--
		}
		return
	}
	deleteLine := t.lines[t.currentLine].delete(t.currentCharacter)
	t.preLenght = t.currentCharacter - 1
	t.ForceSetCurrentCharacter(t.currentCharacter - 1)
	if deleteLine {
		t.totalLine--
		t.lines = slices.Concat(t.lines[:t.currentLine], t.lines[t.currentLine+1:])
		t.currentLine--
		if t.totalLine == 0 {
			t.lines = []*lineText{CreateLineText(t.initialCapacity)}
			t.currentLine = 0
			t.totalLine = 1
			t.ForceSetCurrentCharacter(0)
			return
		}
		t.ForceSetCurrentCharacter(t.lines[t.currentLine].totalChar)
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
func (t *TextBlock) getTextWithAnsi() string {
	full := strings.Builder{}
	for i, line := range t.lines {
		if line == nil {
			break
		}
		full.WriteString(U.GetAnsiMoveTo(t.xPos, t.yPos+i))
		text := line.getText()
		if len(text) > t.currentOffsetChar+t.xSize {
			text = text[t.currentOffsetChar : t.currentOffsetChar+t.xSize]
		} else {
			text = text[t.currentOffsetChar:]
		}
		full.WriteString(text)
	}
	return full.String()
}

func (t *TextBlock) getText() string {
	full := strings.Builder{}
	for _, line := range t.lines {
		if line == nil {
			break
		}
		full.WriteString(line.getText()[t.currentOffsetChar:t.currentOffsetChar+t.xSize] + "\n")
	}
	return full.String()
}
