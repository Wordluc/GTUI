package Component

import "testing"

func Test_CanMoveFunctionStayStill(t *testing.T) { 
	text := StreamCharacter{}
	b := CreateTextBox(0, 0, 10, 10, text)//the interactable area is in x+1,y+1
	if b.CanMoveXCursor(1) != 0 {
		t.Errorf("CanMoveXCursor not working")
	}
	if b.CanMoveYCursor(1) != 0 {
		t.Errorf("CanMoveYCursor not working")
	}
}

func Test_CanMoveFunctionWithCharacter(t *testing.T) {
	stream:=make(chan string)
	text := StreamCharacter{
		Get: func() chan string {
			return stream
		},
	}
	b := CreateTextBox(0, 0, 10, 10, text)
	b.OnClick()
	stream<-"a"
	if b.CanMoveXCursor(1) != 1 {
		t.Errorf("CanMoveXCursor not working")
	}
	if b.CanMoveXCursor(2) != 0 {
		t.Errorf("CanMoveXCursor not working")
	}
	if b.CanMoveXCursor(3) != -1 {
		t.Errorf("CanMoveXCursor not working")
	}
	if b.CanMoveYCursor(1) != 0 {
		t.Errorf("CanMoveYCursor not working")
	}

}
