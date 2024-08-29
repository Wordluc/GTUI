package Drawing

import "testing"

func types(textBlock *TextBlock, text string) {
	for _, char := range text {
		textBlock.Type(char)
	}
}
func TestTypingInTextBlock(t *testing.T) {
	textBlock := CreateTextBlock(0, 0, 10, 100)
	textBlock.Type('a')
	textBlock.Type('2')
	textBlock.Type('3')
	textBlock.Type('\n')
	textBlock.Type('n')
	expLine0 := []rune{'a', '2', '3'}
	expLine1 := []rune{'n'}
	for i := 0; i < 2; i++ {
		if textBlock.lines[0].line[i] != expLine0[i] {
			t.Errorf("expected %v, got %v", string(expLine0[i]), string(textBlock.lines[0].line[i]))
		}
	}
	for i, char := range textBlock.lines[0].getText() {
		if i >= len(expLine0) {
			break
		}
		if char != expLine0[i] {
			t.Errorf("expected %v, got %v", string(expLine0[i]), string(char))
		}
	}
	for i, char := range textBlock.lines[1].getText() {
		if i >= len(expLine1) {
			break
		}
		println(char)
		if char != expLine1[i] {
			t.Errorf("expected %v, got %v", string(expLine1[i]), string(char))
		}
	}
}
func TestComeBackTypingAfterNewLine(t *testing.T) {
	textBlock := CreateTextBlock(0, 0, 10, 100)
	textBlock.Type('a')
	if textBlock.currentLine != 0 {
		t.Errorf("expected %v, got %v", 0, textBlock.currentLine)
	}
	if textBlock.currentCharacter != 1 {
		t.Errorf("expected %v, got %v", 1, textBlock.currentCharacter)
	}
	textBlock.Type('\n')
	if textBlock.currentLine != 1 {
		t.Errorf("expected %v, got %v", 1, textBlock.currentLine)
	}
	if textBlock.currentCharacter != 0 {
		t.Errorf("expected %v, got %v", 0, textBlock.currentCharacter)
	}
	textBlock.ForceSetCurrentLine(0)
	if textBlock.currentLine != 0 {
		t.Errorf("expected %v, got %v", 0, textBlock.currentLine)
	}
	if textBlock.lines[textBlock.currentLine].totalChar != 1 {
		t.Errorf("expected %v, got %v", 1, textBlock.lines[textBlock.currentLine].totalChar)
	}
}
func TestSetCurrentTextBlock(t *testing.T) {
	x, y := 0, 0
	textBlock := CreateTextBlock(70, 0, 10, 100)
	types(textBlock, "a2311\nnnn\nrrrr")
	x, y = textBlock.GetCurrentCursor()
	y--
	textBlock.SetCurrentCursor(x, y)
	if textBlock.preLenght != 4 {
		t.Errorf("expected %v, got %v", 4, textBlock.preLenght)
	}
	if textBlock.currentCharacter != x-1 || textBlock.currentLine != y {
		t.Errorf("expected %v, %v, got %v, %v", x-1, y, textBlock.currentCharacter, textBlock.currentLine)
	}
	y--
	textBlock.SetCurrentCursor(x, y)
	if textBlock.preLenght != 4 {
		t.Errorf("expected %v, got %v", 4, textBlock.preLenght)
	}
	if textBlock.currentCharacter != 4 || textBlock.currentLine != y {
		t.Errorf("expected %v, %v, got %v, %v", 4, y, textBlock.currentCharacter, textBlock.currentLine)
	}
	x++
	textBlock.SetCurrentCursor(x, y)
	y++
	textBlock.SetCurrentCursor(x, y)
	if textBlock.preLenght != 5 {
		t.Errorf("expected %v, got %v", 5, textBlock.preLenght)
	}
	if textBlock.currentCharacter != 3 || textBlock.currentLine != y {
		t.Errorf("expected %v, %v, got %v, %v", 3, y, textBlock.currentCharacter, textBlock.currentLine)
	}
	y--
	textBlock.SetCurrentCursor(x, y)
	if textBlock.preLenght != 5 {
		t.Errorf("expected %v, got %v", 5, textBlock.preLenght)
	}
	if textBlock.currentCharacter != 5 || textBlock.currentLine != y {
		t.Errorf("expected %v, %v, got %v, %v", 5, y, textBlock.currentCharacter, textBlock.currentLine)
	}
	x = 3
	y = 0
	textBlock.SetCurrentCursor(x, y)
	if textBlock.currentCharacter != 3 || textBlock.currentLine != y {
		t.Errorf("expected %v, %v, got %v, %v", 3, y, textBlock.currentCharacter, textBlock.currentLine)
	}
}
func TestComeBackTyping(t *testing.T) {
	textBlock := CreateTextBlock(0, 0, 10, 100)
	textBlock.Type('u')
	textBlock.Type('c')
	textBlock.Type('a')
	textBlock.SetCurrentCursor(0, 0)
	textBlock.Type('l')
	if textBlock.lines[0].getText() != "luca" {
		t.Errorf("expected %v, got %v", "luca", string(textBlock.lines[0].getText()))
	}
	textBlock.SetCurrentCursor(0, 0)
	textBlock.Type('c')
	textBlock.Type('o')
	textBlock.SetCurrentCursor(1, 0)
	textBlock.Type('i')
	textBlock.Type('a')
	textBlock.SetCurrentCursor(4, 0)
	textBlock.Type(' ')
	if textBlock.lines[0].getText() != "ciao luca" {
		t.Errorf("expected %v, got %v", "ciao luca", string(textBlock.lines[0].getText()))
	}
}

func TestDeleteTextline(t *testing.T) {
	textLine := CreateLineText(10)
	textLine.digit('1', 0)
	textLine.digit('t', 1)
	textLine.digit('2', 2)
	textLine.digit('3', 3)
	textLine.digit('4', 4)
	textLine.digit('f', 5)
	textLine.delete(2)
	if textLine.getText() != "1234f" {
		t.Errorf("expected %v, got %v", "1234f", textLine.getText())
	}
	textLine.delete(5)
	if textLine.getText() != "1234" {
		t.Errorf("expected %v, got %v", "1234", textLine.getText())
	}
}
func TestDeleteLine(t *testing.T) {
	textBlock := CreateTextBlock(0, 0, 10, 100)
	textBlock.Type('u')
	textBlock.Type('\n')
	textBlock.Type('a')
	textBlock.Delete()
	if textBlock.totalLine != 2 {
		t.Errorf("expected %v, got %v", 2, textBlock.totalLine)
	}
	textBlock.Delete()
	if textBlock.totalLine != 1 {
		t.Errorf("expected %v, got %v", 1, textBlock.totalLine)
	}
	textBlock.Type('c')
	if textBlock.lines[0].totalChar != 2 {
		t.Errorf("expected %v, got %v", 1, textBlock.lines[0].totalChar)
	}
	textBlock.Delete()
	if textBlock.totalLine != 1 {
		t.Errorf("expected %v, got %v", 1, textBlock.totalLine)
	}
	textBlock.Delete()
	if textBlock.totalLine != 1 {
		t.Errorf("expected %v, got %v", 1, textBlock.totalLine)
	}
	textBlock.Delete()
	textBlock.Delete()
	textBlock.Delete()
	if textBlock.totalLine != 1 {
		t.Errorf("expected %v, got %v", 1, textBlock.totalLine)
	}
	if textBlock.lines[0].totalChar != 0 {
		t.Errorf("expected no characters")
	}
}
func TestFromTwoLinesDoesOnelIne(t *testing.T) {
	textBlock := CreateTextBlock(0, 0, 10, 100)
	textBlock.Type('c')
	textBlock.Type('\n')
	textBlock.Type('a')
	textBlock.SetCurrentCursor(0, 1)
	textBlock.Delete()
	if textBlock.totalLine != 1 {
		t.Errorf("expected %v, got %v", 1, textBlock.totalLine)
	}
	if textBlock.lines[0].totalChar != 2 {
		t.Errorf("expected %v, got %v", 1, textBlock.lines[0].totalChar)
	}
	if textBlock.lines[0].getText() != "ca" {
		t.Errorf("expected %v, got %v", "ca", textBlock.lines[0].getText())
	}
}
func TestNewLineInTheMiddleOfText1(t *testing.T) {
	textBlock := CreateTextBlock(0, 0, 10, 100)
	types(textBlock, "123a1")
	textBlock.SetCurrentCursor(3, 0)
	textBlock.Type('\n')
	if textBlock.lines[0].getText() != "123" {
		t.Errorf("expected %v, got %v", "123", textBlock.lines[0].getText())
	}
	if textBlock.lines[1].getText() != "a1" {
		t.Errorf("expected %v, got %v", "a1", textBlock.lines[1].getText())
	}
}
func TestNewLineInTheMiddleOfText2(t *testing.T) {
	textBlock := CreateTextBlock(0, 0, 10, 100)
	types(textBlock, "123a1\nciao2 prova\ngggg")
	textBlock.SetCurrentCursor(3, 1)
	textBlock.Type('\n')
	if textBlock.lines[1].getText() != "ciao" {
		t.Errorf("expected %v, got %v", "ciao", textBlock.lines[1].getText())
	}
	if textBlock.lines[2].getText() != "2 prova" {
		t.Errorf("expected %v, got %v", "2 prova", textBlock.lines[2].getText())
	}
	if textBlock.lines[3].getText() != "gggg" {
		t.Errorf("expected %v, got %v", "gggg", textBlock.lines[3].getText())
	}
	textBlock.SetCurrentCursor(0, 3)
	textBlock.SetCurrentCursor(1, 3)
	textBlock.Type('\n')
	if textBlock.lines[3].getText() != "g" {
		t.Errorf("expected %v, got %v", "g", textBlock.lines[3].getText())
	}
	if textBlock.lines[4].getText() != "ggg" {
		t.Errorf("expected %v, got %v", "ggg", textBlock.lines[4].getText())
	}
}
