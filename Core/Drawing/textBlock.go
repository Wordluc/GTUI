package Drawing

import (
	U "GTUI/Core/Utils"
	"GTUI/Core/Utils/Color"
	"slices"
	"strings"
)

type TextBlock struct {
	visible         bool
	ansiCode        string
	isChanged       bool
	lines           []*lineText
	xPos            int
	yPos            int
	actualCharacter int
	actualLine       int
	initialCapacity int
	totalLine        int
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
	for _, char := range t.line {
		if char == '\n' {
			break
		}
		str.WriteRune(char)
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
		actualLine:     0,
		initialCapacity: initialCapacity,
		totalLine:        0,
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
	t.actualLine = line
	if t.lines[t.actualLine] == nil {
		t.lines[t.actualLine] = CreateLineText(t.initialCapacity)
	}
}

func (t *TextBlock) SetCurrentChar(ichar int) {
	t.actualCharacter=ichar
}

func (t *TextBlock) GetTotalChar() int {
	return t.lines[t.actualLine].totalChar
}

func (t *TextBlock) GetLastLine() int {
	return t.totalLine
}
func (t *TextBlock) GetActualCursor()(int,int){
   return t.actualCharacter,t.actualLine
}
func (t *TextBlock) Type(char rune) {
	if char == '\n' {
		t.lines[t.actualLine].digit(char)
		t.totalLine++
		t.actualLine++
		if t.actualLine >= len(t.lines) {
			t.lines = append(t.lines, CreateLineText(t.initialCapacity))
		}
		if t.lines[t.actualLine] == nil {
			t.lines[t.actualLine] = CreateLineText(t.initialCapacity)
		}
	} else {
		t.lines[t.actualLine].digit(char)
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
