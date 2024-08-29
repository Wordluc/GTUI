package Component

import "testing"

func Test_CanMoveFunctionStayStill(t *testing.T) { 
	text := StreamCharacter{}
	b := CreateTextBox(0, 0, 10, 10, text)//the interactable area is in x+1,y+1
	x,y:=b.DiffTotalToXY(1,1)
	if x != 0 {
		t.Errorf("CanMoveXCursor not working")
	}
	if y != 0 {
		t.Errorf("CanMoveYCursor not working")
	}
}

func Test_CanMoveFunctionWithCharacter(t *testing.T) {
	stream:=make(chan string)
	text := StreamCharacter{
		Get: func() chan string {
			return stream
		},
		Delete: func() {},
	}
	b := CreateTextBox(0, 0, 10, 10, text)
	b.StartTyping()
	stream<-"a"
	x,_:=b.DiffTotalToXY(1,1)
	if x != 1 {
		t.Errorf("CanMoveXCursor not working")
	}
	x,_=b.DiffTotalToXY(2,1)
	if x != 0 {
		t.Errorf("CanMoveXCursor not working")
	}
	x,_=b.DiffTotalToXY(3,1)
	if x != -1 {
		t.Errorf("CanMoveXCursor not working")
	}
	_,y:=b.DiffTotalToXY(3,1)
	if y != 0 {
		t.Errorf("CanMoveYCursor not working")
	}
}
