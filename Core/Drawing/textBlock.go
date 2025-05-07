package Drawing

import (
	"errors"
	"slices"
	"strings"

	"github.com/Wordluc/GTUI/Core"
	"github.com/Wordluc/GTUI/Core/EventManager"
	U "github.com/Wordluc/GTUI/Core/Utils"
	"github.com/Wordluc/GTUI/Core/Utils/Color"
)

type lineText struct {
	line            []rune
	initialCapacity int
	totalChar       int
}

func createLineText(initialCapacity int) *lineText {
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
		return createLineText(t.initialCapacity)
	}
	splited := t.line[i:t.totalChar]
	t.line = t.line[:i]
	newLine := createLineText(t.initialCapacity)
	newLine.totalChar = t.totalChar - i
	newLine.line = splited
	t.totalChar = i
	return newLine
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
	xStartingWrapping        int
	yStartingWrapping        int
	absoluteCurrentCharacter int
	currentLine              int
	initialCapacity          int
	totalLine                int
	preLenght                int
	tabSize                  int
	wrap                     bool
	layer                    Core.Layer
	color                    Color.Color
}

func CreateTextBlock(x, y int, xSize, ySize int, initialCapacity int) *TextBlock {
	lines := make([]*lineText, 1)
	lines[0] = createLineText(initialCapacity)
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
		tabSize:                  4,
		wrap:                     false,
		color:                    Color.GetDefaultColor(),
	}
}

func (t *TextBlock) Touch() {
	t.isChanged = true
}
func (t *TextBlock) IsTouched() bool {
	return t.isChanged
}
func (t *TextBlock) GetAnsiCode(defaultColor Color.Color) string {
	if t.isChanged {
		t.ansiCode = t.makeAnsiCode(defaultColor)
		t.isChanged = false
	}
	return t.ansiCode
}
func (t *TextBlock) SetColor(color Color.Color) {
	t.color = color
	t.Touch()
}
func (t *TextBlock) SetPos(x, y int) {
	t.xPos = x
	t.yPos = y
	t.Touch()
	EventManager.Call(EventManager.ReorganizeElements, t)
}

func (b *TextBlock) SetLayer(layer Core.Layer) error {
	if layer < 0 {
		return errors.New("layer can't be negative")
	}
	b.layer = layer
	b.Touch()
	EventManager.Call(EventManager.ReorganizeElements, b)
	return nil
}

func (c *TextBlock) GetLayer() Core.Layer {
	return c.layer
}

func (t *TextBlock) GetSize() (int, int) {
	return t.xSize - 1, t.ySize
}

func (t *TextBlock) GetPos() (int, int) {
	return t.xPos, t.yPos
}
func (t *TextBlock) SetVisibility(visible bool) {
	t.visible = visible
	t.Touch()
	EventManager.Call(EventManager.ForceRefresh)
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

func (t *TextBlock) SetWrap(isOn bool) {
	t.wrap = isOn
	if isOn {
		t.xStartingWrapping = t.absoluteCurrentCharacter
		t.yStartingWrapping = t.currentLine
	}
}
func (t *TextBlock) GetWrap() bool {
	return t.wrap
}

// get the wrapped text
func (t *TextBlock) GetSelectedText() string {
	if !t.wrap {
		return ""
	}
	var str strings.Builder
	xStarting := t.xStartingWrapping
	yStarting := t.yStartingWrapping
	xEnding := t.absoluteCurrentCharacter
	yEnding := t.currentLine
	if t.currentLine < t.yStartingWrapping || (t.currentLine == t.yStartingWrapping && t.absoluteCurrentCharacter < t.xStartingWrapping) {
		xStarting = t.absoluteCurrentCharacter
		yStarting = t.currentLine
		xEnding = t.xStartingWrapping
		yEnding = t.yStartingWrapping
	}
	for yi, line := range t.lines[yStarting:] {
		if yi == yStarting {
			if xStarting > line.totalChar {
				continue
			} else {
				if yEnding == yi {
					str.WriteString(string(line.line[xStarting:xEnding]))
					continue
				}
				str.WriteString(string(line.line[xStarting:]) + "\n")
				continue
			}
		}
		if yi > yEnding {
			break
		}
		if yi == yEnding {
			if xEnding > line.totalChar {
				str.WriteString(string(line.line))
			} else {
				str.WriteString(string(line.line[:xEnding]))
			}
			break
		}
		str.WriteString(string(line.line) + "\n")
	}
	t.wrap = false
	return strings.ReplaceAll(str.String(), "\x00", "")
}

type WrapperEdges struct {
	startX int
	startY int
	endX   int
	endY   int
}

func (t *TextBlock) getWrapperEdges() WrapperEdges {
	var result WrapperEdges = WrapperEdges{
		startX: t.xStartingWrapping,
		startY: t.yStartingWrapping,
		endX:   t.absoluteCurrentCharacter,
		endY:   t.currentLine,
	}
	//invert the cursor with the wrapping starting, so the cursor point will be bigger than the wrapping
	if t.currentLine < t.yStartingWrapping || (t.currentLine == t.yStartingWrapping && t.absoluteCurrentCharacter < t.xStartingWrapping) {
		result.startX = t.absoluteCurrentCharacter
		result.startY = t.currentLine
		result.endX = t.xStartingWrapping
		result.endY = t.yStartingWrapping
	}
	result.endX = result.endX - t.xRelativeMinSize
	result.endY = result.endY - t.yRelativeMinSize
	result.startY = result.startY - t.yRelativeMinSize
	result.startX = result.startX - t.xRelativeMinSize
	if result.startX < 0 {
		result.startX = 0
	}
	if result.endX < 0 {
		result.endX = 0
	}
	if result.startY < 0 {
		result.startY = 0
	}
	if result.endY < 0 {
		result.endY = 0
	}
	return result
}

func (t *TextBlock) ClearAll() {
	t.lines = make([]*lineText, 1)
	t.lines[0] = createLineText(t.initialCapacity)
	t.xRelativeMinSize = 0
	t.yRelativeMinSize = 0
	t.xRelativeMaxSize = t.xSize + 1
	t.yRelativeMaxSize = t.ySize
	t.currentLine = 0
	t.totalLine = 1
	t.absoluteCurrentCharacter = 0
	t.wrap = false
	t.Touch()
}

// delete the selected text
func (t *TextBlock) deleteWrapping() {
	defer t.Touch()
	t.wrap = false
	edge := t.getWrapperEdges()
	t.setXCursor_Absolute(edge.endX)
	t.setYCursor_Absolute(edge.endY)
	t.xStartingWrapping = edge.startX
	t.yStartingWrapping = edge.startY
	for t.absoluteCurrentCharacter > t.xStartingWrapping || t.currentLine > t.yStartingWrapping {
		t.Delete()
	}
}

// Delete the current character
func (t *TextBlock) Delete() {
	t.Touch()
	if t.wrap {
		t.deleteWrapping()
		return
	}
	if t.totalLine == 1 && t.absoluteCurrentCharacter == 0 {
		return
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
		return
	}
	deleteLine := t.lines[t.currentLine].delete(t.absoluteCurrentCharacter)
	t.preLenght = t.absoluteCurrentCharacter - 1
	t.setXCursor_Absolute(t.absoluteCurrentCharacter - 1)
	if deleteLine {
		t.totalLine--
		t.lines = slices.Concat(t.lines[:t.currentLine], t.lines[t.currentLine+1:])
		t.setYCursor_Relative(t.getYCursor_Relative() - 1)
		if t.totalLine == 0 {
			t.lines = []*lineText{createLineText(t.initialCapacity)}
			t.currentLine = 0 //Reset y and min max
			t.totalLine = 1
			t.resetXCursor()
			return
		}
		t.setXCursor_Absolute(t.lines[t.currentLine].totalChar)
	}
}

func insertTextToOrigin(origin, text string, pos int) string {
	if pos > len(origin) {
		return origin + text
	}
	return origin[:pos] + text + origin[pos:]
}
func (t *TextBlock) makeAnsiCode(defaultColor Color.Color) string {
	full := strings.Builder{}
	full.WriteString(t.color.GetMixedColor(defaultColor).GetAnsiColor())
	y := 0
	if !t.wrap {
		t.xStartingWrapping = t.absoluteCurrentCharacter
		t.yStartingWrapping = t.currentLine
	}
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
		full.WriteString(U.GetAnsiMoveTo(t.xPos, t.yPos+y))
		text := line.getText()
		text = t.parseText(text)
		edge := t.getWrapperEdges()
		if t.wrap {
			if edge.endY == i {
				text = insertTextToOrigin(text, "\033[m", edge.endX)
			}
			if edge.startY == i {
				text = insertTextToOrigin(text, "\033[100m", edge.startX)
			}
		}

		full.WriteString(text)
		y++
	}
	full.WriteString(defaultColor.GetAnsiColor())
	return strings.TrimSuffix(full.String(), "\n")
}
func (t *TextBlock) GetText() string {
	var res strings.Builder
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
		text := line.getText()
		text = t.parseText(text)
		res.WriteString(text + "\n")
	}
	return strings.TrimSuffix(res.String(), "\n")
}
func (t *TextBlock) parseText(text string) string {
	start := 0
	apperanceSize := len([]rune(text))
	logicSize := len(text)
	if apperanceSize > t.xRelativeMinSize {
		start = t.xRelativeMinSize
	} else {
		return ""
	}
	//da sistemare
	if apperanceSize > t.xRelativeMaxSize-1 {
		if logicSize == apperanceSize {
			logicSize = t.xRelativeMaxSize - 1
		} else {
			logicSize = t.xRelativeMaxSize + (logicSize - apperanceSize - 1)
		}
	}

	//	if diff := logicSize - start; diff > t.xSize {
	//		start += diff - t.xSize
	//	}
	return text[start:logicSize]
}

// Set the absolute character position by looking at the relative position in the window,
// if x>=size the text is slided to the right
// if x<0 the text is slided to the left
func (t *TextBlock) setXCursor_Relative(x int) {
	defer t.Touch()
	if x >= t.xSize {
		if t.absoluteCurrentCharacter+x-t.xSize+1 > t.lines[t.currentLine].totalChar {
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

// Set the absolute character position by looking at the relative position in the window,
// if y>=size the text is slided to the right
// if y<0 the text is slided to the left
func (t *TextBlock) setYCursor_Relative(y int) {
	defer t.Touch()
	if y >= t.ySize {
		if t.currentLine+y-t.ySize >= t.totalLine {
			return
		}
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

func (t *TextBlock) setYCursor_Absolute(y int) {
	defer t.Touch()
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
		return
	}
	if t.lines[t.currentLine] == nil {
		t.lines[t.currentLine] = createLineText(t.initialCapacity)
	}
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

	if xRelative+t.xRelativeMinSize >= t.lines[t.yRelativeMinSize+yRelative].totalChar {
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
func (t *TextBlock) typeSilent(char rune) {
	if t.wrap {
		t.deleteWrapping()
	}
	if char == '\n' {
		t.totalLine++
		t.setYCursor_Relative(t.getYCursor_Relative() + 1)
		newLine := t.lines[t.currentLine-1].split(t.absoluteCurrentCharacter)
		t.lines = slices.Concat(t.lines[:t.currentLine], []*lineText{newLine}, t.lines[t.currentLine:])
		t.resetXCursor()
	} else {
		if char == '	' {
			for i := 0; i < t.tabSize; i++ {
				t.lines[t.currentLine].digit(' ', t.absoluteCurrentCharacter)
				t.setXCursor_Relative(t.getXCursor_Relative() + 1)
				t.preLenght = t.absoluteCurrentCharacter
			}
		} else {
			t.lines[t.currentLine].digit(char, t.absoluteCurrentCharacter)
			t.setXCursor_Relative(t.getXCursor_Relative() + 1)
			t.preLenght = t.absoluteCurrentCharacter
		}
	}
}
func (t *TextBlock) Type(char rune) {
	defer t.Touch()
	t.typeSilent(char)
}

func (t *TextBlock) SetText(s string) {
	t.ClearAll()
	for _, char := range s {
		t.typeSilent(char)
	}
	t.Touch()
}
