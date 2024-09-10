package Drawing

import (
	U "GTUI/Core/Utils"
	"GTUI/Core/Utils/Color"
	"slices"
	"strings"
)

type TextBlock struct {
	visible                  bool
	ansiCode                 string
	isChanged                bool
	lines                    []*lineText
	xPos                     int
	yPos                     int
	xSize                    int
	ySize                    int
	xRelativeMaxSize         int
	xRelativeMinSize         int
	absoluteCurrentCharacter int
	currentLine              int
	initialCapacity          int
	totalLine                int
	preLenght                int
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
		visible:          true,
		ansiCode:         "",
		isChanged:        true,
		lines:            lines,
		xPos:             x,
		yPos:             y,
		xSize:            xSize+1,
		ySize:            ySize,
		xRelativeMinSize: 0,
		xRelativeMaxSize: xSize+1,
		currentLine:      0,
		initialCapacity:  initialCapacity,
		totalLine:        1,
		absoluteCurrentCharacter:0,
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
///Set the absolute character position, 
///if x==size the text is slided to the right
///if x==-1 the text is slided to the left
func (t *TextBlock) ForceSetRelativeCharacter(x int) {
	t.Touch()
	if x==t.xSize{
		if t.absoluteCurrentCharacter==t.lines[t.currentLine].totalChar{
			return
		}
		t.xRelativeMaxSize++
		t.xRelativeMinSize++
		t.absoluteCurrentCharacter++
		return
	}
	if x==-1 {
		t.absoluteCurrentCharacter--
		t.xRelativeMaxSize--
		t.xRelativeMinSize--
		if t.xRelativeMinSize<0{
			t.xRelativeMinSize=0
		}
		if t.xRelativeMaxSize<t.xSize{
			t.xRelativeMaxSize=t.xSize
		}
		if t.absoluteCurrentCharacter<0{
			t.absoluteCurrentCharacter=0
		}
		return
	}
	t.absoluteCurrentCharacter = t.xRelativeMinSize+x
}

func (t *TextBlock) ResetCurrentCharacter() {
	t.absoluteCurrentCharacter = 0
	t.xRelativeMinSize = 0
	t.xRelativeMaxSize = t.xSize
	t.preLenght=0
}
func (t *TextBlock) GetRelativeCurrentCharacter() int {
	if r:= t.absoluteCurrentCharacter - t.xRelativeMinSize;r<0{
		return 0
	}else{
    return r
	}
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

func (t *TextBlock) SetCurrentCursor(x, y int) (int, int) {//sistemare la logica di cambio y
	xRelative := x - t.xPos
	yRelative := y - t.yPos

	if yRelative < 0 {
		yRelative = 0
	}
	if yRelative >= t.ySize {
		return x, t.ySize
	}
	isYChanged := yRelative != t.currentLine
	isXChanged := xRelative != t.GetRelativeCurrentCharacter()
	if yRelative >= len(t.lines) {
		yRelative = len(t.lines) - 1
	}
	if xRelative >= t.lines[yRelative].totalChar {
		xRelative = t.lines[yRelative].totalChar
	}
	if isXChanged {
		t.ForceSetRelativeCharacter(xRelative)
		t.preLenght = t.absoluteCurrentCharacter
	}
	if isYChanged {//TODO: da refattorizzare
		if t.lines[yRelative].totalChar > t.preLenght {
			nSlide:=t.preLenght-t.xSize
			if nSlide<0{
				nSlide=0
			}
			t.xRelativeMinSize=nSlide
			if t.preLenght>=t.xSize{
				t.xRelativeMinSize=nSlide+1
			}
			if t.xRelativeMinSize<0{
				t.xRelativeMinSize=0
			}
			t.xRelativeMaxSize=nSlide+t.xSize
			t.absoluteCurrentCharacter = t.preLenght
		} else {
			nSlide:=t.lines[yRelative].totalChar-t.xSize
			if nSlide<0{
				nSlide=0
			}
			t.xRelativeMinSize=nSlide
			if t.lines[yRelative].totalChar>=t.xSize{
				t.xRelativeMinSize=nSlide+1
			}
			if t.xRelativeMinSize<0{
				t.xRelativeMinSize=0
			}
			t.xRelativeMaxSize=nSlide+t.xSize
			t.absoluteCurrentCharacter = t.lines[yRelative].totalChar
		}
		t.ForceSetCurrentLine(yRelative)
	}
	t.Touch()
	return t.GetRelativeCurrentCharacter() + t.xPos, yRelative + t.yPos

}

func (t *TextBlock) GetTotalCursor() (int, int) {
	return t.lines[t.currentLine].totalChar, t.totalLine
}
func (t *TextBlock) GetCurrentCursor() (int, int) {
	return t.GetRelativeCurrentCharacter(), t.currentLine
}
func (t *TextBlock) Type(char rune) {
	defer t.Touch()
	if char == '\n' {
		t.totalLine++
		t.currentLine++
		if t.currentLine >= len(t.lines) {
			t.lines = append(t.lines, CreateLineText(t.initialCapacity))
			t.ResetCurrentCharacter()
			return
		}	
		if t.absoluteCurrentCharacter <= t.lines[t.currentLine-1].totalChar {
			newLine := t.lines[t.currentLine-1].split(t.absoluteCurrentCharacter)
			t.lines = slices.Concat(t.lines[:t.currentLine], []*lineText{newLine}, t.lines[t.currentLine:])
			t.ResetCurrentCharacter()
			return
		}
	} else {
		t.lines[t.currentLine].digit(char, t.absoluteCurrentCharacter)
		t.ForceSetRelativeCharacter(t.GetRelativeCurrentCharacter()+1)
		t.preLenght = t.absoluteCurrentCharacter
	}
}
func printArray(array []rune) {
	for i, v := range array {
		println(i, ":", string(v))
	}
}
func (t *TextBlock) Delete() {
	if t.totalLine == 1 && t.absoluteCurrentCharacter == 0 {
		return
	}
	defer t.Touch()
	if t.absoluteCurrentCharacter == 0 {
		if t.currentLine-1 >= 0 {
			t.ForceSetRelativeCharacter(t.lines[t.currentLine-1].totalChar)
			t.preLenght = t.absoluteCurrentCharacter
			t.lines[t.currentLine-1].merge(t.lines[t.currentLine])
			t.lines = slices.Concat(t.lines[:t.currentLine], t.lines[t.currentLine+1:])
			t.currentLine--
			t.totalLine--
		}
		return
	}
	deleteLine := t.lines[t.currentLine].delete(t.absoluteCurrentCharacter)
	t.preLenght = t.absoluteCurrentCharacter - 1
	t.ForceSetRelativeCharacter(t.absoluteCurrentCharacter - 1)
	if deleteLine {
		t.totalLine--
		t.lines = slices.Concat(t.lines[:t.currentLine], t.lines[t.currentLine+1:])
		t.currentLine--
		if t.totalLine == 0 {
			t.lines = []*lineText{CreateLineText(t.initialCapacity)}
			t.currentLine = 0
			t.totalLine = 1
			t.ResetCurrentCharacter()
			return
		}
		t.ForceSetRelativeCharacter(t.lines[t.currentLine].totalChar)
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
func (t *TextBlock) parseText(text string) string {
	start := 0
	size := len(text)
	if size > t.xRelativeMinSize{
		start = t.xRelativeMinSize
	}else{
		return ""
	}
	if size > t.xRelativeMaxSize-1 {
		size = t.xRelativeMaxSize-1
	}//i valori size and start devono globali
	if diff:=size-start;diff>t.xSize{
		start+=diff-t.xSize
	}
	return text[start:size]
}
func (t *TextBlock) getTextWithAnsi() string {
	full := strings.Builder{}
	for i, line := range t.lines {
		if line == nil {
			break
		}
		full.WriteString(U.GetAnsiMoveTo(t.xPos, t.yPos+i))
		text := line.getText()
		full.WriteString(t.parseText(text))
	}
	return full.String()
}

func (t *TextBlock) getText() string {
	full := strings.Builder{}
	for _, line := range t.lines {
		if line == nil {
			break
		}
		text := line.getText()
		full.WriteString(t.parseText(text) + "\n")
	}
	return full.String()
}
