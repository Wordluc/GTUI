package Drawing

import (
	U "GTUI/Core/Utils"
	"GTUI/Core/Utils/Color"
	"errors"
	"slices"
	"strings"
)

type lineText struct {
	line            []rune
	initialCapacity int
	totalChar       int
}

func CreateLineText(initialCapacity int) (*lineText,error) {
	if initialCapacity < 0 {
		return nil, errors.New("invalid capacity")
	}
	return &lineText{
		line:            make([]rune, initialCapacity),
		initialCapacity: initialCapacity,
		totalChar:       0,
	},nil
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
///delete the character i, if the line is empty return true
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

func (t *lineText) split(i int) (*lineText,error) {
	var e error
	if i < 0{
		return nil,errors.New("index lineText is negative")
	}
	if i >= t.totalChar {
		return CreateLineText(t.initialCapacity)
	}
	if t.totalChar > len(t.line) {
		return nil,errors.New("index lineText is out of range")
	}
	if i >len(t.line){
		return nil,errors.New("index lineText is out of range")
	}
	splited := t.line[i:t.totalChar]
	t.line = t.line[:i]
	newLine,e := CreateLineText(t.initialCapacity)
	if e != nil {
		return nil, e
	}
	newLine.totalChar = t.totalChar - i
	newLine.line = splited
	t.totalChar = i
	return newLine,nil
}

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
	yRelativeMaxSize         int
	yRelativeMinSize         int
	absoluteCurrentCharacter int
	currentLine              int
	initialCapacity          int
	totalLine                int
	preLenght                int
}

func CreateTextBlock(x, y int, xSize, ySize int, initialCapacity int) (*TextBlock ,error){
	lines := make([]*lineText, 1)
	var e error
	if lines == nil {
		return nil, errors.New("failed to create text block")
	}
	lines[0],e = CreateLineText(initialCapacity)
	if e != nil {
		return nil, e
	}
	return &TextBlock{
		visible:                  true,
		ansiCode:                 "",
		isChanged:                true,
		lines:                    lines,
		xPos:                     x,
		yPos:                     y,
		xSize:                    xSize + 1,
		ySize:                    ySize,
		xRelativeMinSize:         0,
		yRelativeMinSize:         0,
		xRelativeMaxSize:         xSize + 1,
		yRelativeMaxSize:         ySize,
		currentLine:              0,
		initialCapacity:          initialCapacity,
		totalLine:                1,
		absoluteCurrentCharacter: 0,
	},nil
}

func (t *TextBlock) Touch() {
	t.isChanged = true
}

func (t *TextBlock) GetAnsiCode(defaultColor Color.Color) string {
	if t.isChanged {
		t.ansiCode = t.GetText(true)
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

func (t *TextBlock) GetPos() (int, int) {
	return t.xPos, t.yPos
}

func (t *TextBlock) SetVisibility(visible bool) {
	t.visible = visible
}

func (t *TextBlock) GetVisibility() bool {
	return t.visible
}

func (t *TextBlock) GetTotalCursor() (int, int) {
	return t.lines[t.currentLine].totalChar, t.totalLine
}

func (t *TextBlock) GetCursor_Relative() (int, int) {
	return t.getXCursor_Relative(), t.getYCursor_Relative()
}

///Delete the current character
func (t *TextBlock) Delete() error {
	if t.totalLine == 1 && t.absoluteCurrentCharacter == 0 {
		return nil
	}
	defer t.Touch()
	if t.absoluteCurrentCharacter == 0 {
		if t.currentLine-1 >= 0 {
			t.setXCursor_Absolute(t.lines[t.currentLine-1].totalChar)
			t.preLenght = t.absoluteCurrentCharacter
			t.lines[t.currentLine-1].merge(t.lines[t.currentLine])
			t.lines = slices.Concat(t.lines[:t.currentLine], t.lines[t.currentLine+1:])
			t.setYCursor_Relative(t.getYCursor_Relative() - 1)
			t.totalLine--
		}
		return nil
	}
	deleteLine := t.lines[t.currentLine].delete(t.absoluteCurrentCharacter)
	t.preLenght = t.absoluteCurrentCharacter - 1
	t.setXCursor_Absolute(t.absoluteCurrentCharacter - 1)
	if deleteLine {
		t.totalLine--
		t.lines = slices.Concat(t.lines[:t.currentLine], t.lines[t.currentLine+1:])
		t.setYCursor_Relative(t.getYCursor_Relative() - 1)
		if t.totalLine == 0 {
			line,e:= CreateLineText(t.initialCapacity)
			if e != nil {
				return e
			}
			t.lines = []*lineText{line}
			t.currentLine = 0 //Reset y and min max
			t.totalLine = 1
			t.resetXCursor()
			return nil
		}
		t.setXCursor_Absolute(t.lines[t.currentLine].totalChar)
	}
	return nil
}

func (t *TextBlock) GetText(withAnsiCode bool) string {
	full := strings.Builder{}
	y := 0
	for i, line := range t.lines {
		if i < t.yRelativeMinSize {
			continue
		}
		if i >= t.yRelativeMaxSize {
			break
		}
		if line == nil {
			break
		}
		if withAnsiCode {
			full.WriteString(U.GetAnsiMoveTo(t.xPos, t.yPos+y))
		}
		text := line.getText()
		full.WriteString(t.parseText(text))
		if !withAnsiCode {
			full.WriteRune('\n')
		}
		y++
	}
	return full.String()
}
// /Set the absolute character position by looking at the relative position in the window,
// /if x>=size the text is slided to the right
// /if x<0 the text is slided to the left
func (t *TextBlock) setXCursor_Relative(x int) {
	defer t.Touch()
	if x >= t.xSize {
		if t.absoluteCurrentCharacter+x - t.xSize+1>t.lines[t.currentLine].totalChar {
			return
		}
		t.xRelativeMaxSize += x - t.xSize + 1
		t.xRelativeMinSize += x - t.xSize + 1
		t.absoluteCurrentCharacter += x - t.xSize + 1
		return
	}
	if x < 0 {
		t.absoluteCurrentCharacter += x
		t.xRelativeMaxSize += x
		t.xRelativeMinSize += x
		if t.xRelativeMinSize < 0 {
			t.xRelativeMinSize = 0
		}
		if t.xRelativeMaxSize < t.xSize {
			t.xRelativeMaxSize = t.xSize
		}
		if t.absoluteCurrentCharacter < 0 {
			t.absoluteCurrentCharacter = 0
		}
		return
	}
	t.absoluteCurrentCharacter = t.xRelativeMinSize + x
}

func (t *TextBlock) setXCursor_Absolute(x int) {
	defer t.Touch()
	nSlide := x - t.xSize
	if nSlide < 0 {
		nSlide = 0
	}
	t.xRelativeMinSize = nSlide
	if x >= t.xSize {
		nSlide++
	}
	t.xRelativeMinSize = nSlide
	t.xRelativeMaxSize = nSlide + t.xSize
	if t.xRelativeMinSize < 0 {
		t.xRelativeMinSize = 0
	}
	t.absoluteCurrentCharacter = x
}

// /Set the absolute character position by looking at the relative position in the window,
// /if y>=size the text is slided to the right
// /if y<0 the text is slided to the left
func (t *TextBlock) setYCursor_Relative(y int) {
	defer t.Touch()
	if y >= t.ySize {
		if t.currentLine+y-t.ySize >= t.totalLine {
			return
		}
		//		if diff:=t.currentLine - t.totalLine;diff>0 {
		//			y=-diff
		//		}
		t.yRelativeMaxSize += y - t.ySize + 1
		t.yRelativeMinSize += y - t.ySize + 1
		t.currentLine += y - t.ySize + 1
		return
	}
	if y < 0 {
		t.currentLine += y
		t.yRelativeMaxSize += y
		t.yRelativeMinSize += y
		if t.yRelativeMinSize < 0 {
			t.yRelativeMinSize = 0
		}
		if t.yRelativeMaxSize < t.ySize {
			t.yRelativeMaxSize = t.ySize
		}
		if t.currentLine < 0 {
			t.currentLine = 0
		}
		return
	}
	t.currentLine = t.yRelativeMinSize + y
}

func (t *TextBlock) getXCursor_Relative() int {
	if r := t.absoluteCurrentCharacter - t.xRelativeMinSize; r < 0 {
		return 0
	} else {
		return r
	}
}

func (t *TextBlock) setYCursor_Absolute(y int) error {
	defer t.Touch()
	var e error
	nSlide := y - t.ySize
	if nSlide < 0 {
		nSlide = 0
	}
	t.yRelativeMinSize = nSlide
	if y >= t.ySize {
		nSlide++
	}
	t.yRelativeMinSize = nSlide
	t.yRelativeMaxSize = nSlide + t.ySize
	if t.yRelativeMinSize < 0 {
		t.yRelativeMinSize = 0
	}
	t.currentLine = y
	if y >= len(t.lines) {
		return nil
	}
	if t.lines[t.currentLine] == nil {
		t.lines[t.currentLine],e = CreateLineText(t.initialCapacity)
	}
	return e
}

func (t *TextBlock) getYCursor_Relative() int {
	if r := t.currentLine - t.yRelativeMinSize; r < 0 {
		return 0
	} else {
		return r
	}
}

func (t *TextBlock) SetCursor_Relative(x, y int) (int, int) {
	xRelative := x - t.xPos
	yRelative := y - t.yPos
	isYChanged := yRelative != t.getYCursor_Relative()
	isXChanged := xRelative != t.getXCursor_Relative()
  
	if t.yRelativeMinSize+yRelative >= len(t.lines) {
		yRelative = len(t.lines) - 1 - t.yRelativeMinSize
	}

	if t.yRelativeMinSize+yRelative < 0 {
		yRelative = 0
	}

	if t.xRelativeMinSize+xRelative < 0 {
		xRelative = 0
	}

	if xRelative + t.xRelativeMinSize >= t.lines[t.yRelativeMinSize+yRelative].totalChar {
		xRelative = t.lines[t.yRelativeMinSize+yRelative].totalChar - t.xRelativeMinSize 
	}
	if isYChanged {
		t.setYCursor_Relative(yRelative)
		if t.lines[t.currentLine].totalChar > t.preLenght {
			t.setXCursor_Absolute(t.preLenght)
		} else {
			t.setXCursor_Absolute(t.lines[t.currentLine].totalChar)
		}
	}
	if isXChanged {
		t.setXCursor_Relative(xRelative)
		t.preLenght = t.absoluteCurrentCharacter
	}
	t.Touch()
	return t.getXCursor_Relative() + t.xPos, t.getYCursor_Relative() + t.yPos
}

func (t *TextBlock) resetXCursor() {
	t.absoluteCurrentCharacter = 0
	t.xRelativeMinSize = 0
	t.xRelativeMaxSize = t.xSize
	t.preLenght = 0
}

func (t *TextBlock) Type(char rune) error {
	defer t.Touch()
	if char == '\n' {
		t.totalLine++
		t.setYCursor_Relative(t.getYCursor_Relative() + 1)
		newLine,e := t.lines[t.currentLine-1].split(t.absoluteCurrentCharacter)
		if e != nil {
			return e
		}
		t.lines = slices.Concat(t.lines[:t.currentLine], []*lineText{newLine}, t.lines[t.currentLine:])
		t.resetXCursor()
	} else {
		t.lines[t.currentLine].digit(char, t.absoluteCurrentCharacter)
		t.setXCursor_Relative(t.getXCursor_Relative() + 1)
		t.preLenght = t.absoluteCurrentCharacter
	}
	return nil
}
// TODO da inserire i colori
func (t *TextBlock) parseText(text string) string {
	start := 0
	size := len(text)
	if size > t.xRelativeMinSize {
		start = t.xRelativeMinSize
	} else {
		return ""
	}
	if size > t.xRelativeMaxSize-1 {
		size = t.xRelativeMaxSize - 1
	} //i valori size and start devono globali
	if diff := size - start; diff > t.xSize {
		start += diff - t.xSize
	}
	return text[start:size]
}

