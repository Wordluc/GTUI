package Component

import (
	"testing"
)
var stream=StreamCharacter{}
func TestPaste(t *testing.T) {
	textBlock := CreateTextBox(0, 0, 10, 100, stream)
	textBlock.Paste("test\r\n")
	if textBlock.textBlock.GetText(false) != "test\n\n" {
		t.Errorf("expected %v, got %v", "test", textBlock.textBlock.GetText(false))
	}
}
func TestSpecialCharacter(t *testing.T) {
	textBlock := CreateTextBox(0, 0, 11, 100, stream)
	textBlock.Paste("sono lucà")
	if textBlock.textBlock.GetText(false) != "sono lucà\n" {
		t.Errorf("expected %v, got %v", "sono lucà", textBlock.textBlock.GetText(false))
	}
}
