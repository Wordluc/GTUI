package EventManager

import (
	"testing"
	"time"
)

func testTime(t time.Time, timeToWait int) (int, bool) {
	res := (time.Now().UnixNano() - t.UnixNano()) / int64(time.Millisecond)
	return int(res), int(res) == timeToWait
}
func _TestEventManager(t *testing.T) {
	Setup()
	var isok bool
	var res int
	timer := time.Now()

	Subscribe(Refresh, 2, func(comp []any) {
		res, isok = testTime(timer, 2)
	})
	Call(Refresh, nil)
	time.Sleep(100 * time.Millisecond)
	if !isok {
		t.Errorf("expected to wait 2, got %v", res)
	}
}
type callerForTest struct {
	name string
}
func _TestEventManagerWithMultipeCall(t *testing.T) {
	delete(eventManager.subscribers,Refresh)
	var timer time.Time
	var isok bool
	var ncalls int
	var res int
	var caller string
	timer = time.Now()
	e:=Subscribe(Refresh, 500, func(comp []any) {
		ncalls++
		res, isok = testTime(timer, 500)
		for _,c:=range comp{
			caller+=c.(callerForTest).name
		}
		caller+="|"
	})
	if e != nil {
		t.Errorf("expected no error,got %v", e)
	}
	callerA:=callerForTest{"a"}
	callerB:=callerForTest{"b"}
	if e:=Call(Refresh, []any{callerA});e!=nil{
		t.Errorf("expected no error,got %v", e)
	}
	time.Sleep(100 * time.Millisecond)
	Call(Refresh, []any{callerB})
	time.Sleep(500 * time.Millisecond)
	if !isok {
		t.Errorf("expected to wait 2,got %d",res)
	}
	if ncalls != 1 {
		t.Errorf("expected 1 calls,got %d", ncalls)
	}
	if caller!="ab|"{
		t.Errorf("expected ab|,got %s", caller)
	}
}
func _TestEventManagerWithDistantMultipeCall(t *testing.T) {
	delete(eventManager.subscribers,Refresh)
	var ncalls int
	var caller string
	e:=Subscribe(Refresh, 300, func(comp []any) {
		ncalls++
		for _,c:=range comp{
			caller+=c.(callerForTest).name
		}
		caller+="|"
	})
	if e != nil {
		t.Errorf("expected no error,got %v", e)
	}
	callerA:=callerForTest{"a"}
	callerB:=callerForTest{"b"}
	if e:=Call(Refresh, []any{callerA});e!=nil{
		t.Errorf("expected no error,got %v", e)
	}
	time.Sleep(550 * time.Millisecond)
	Call(Refresh, []any{callerB})
	time.Sleep(500 * time.Millisecond)
	if ncalls != 2 {
		t.Errorf("expected 2 calls,got %d", ncalls)
	}
	if caller!="a|b|"{
		t.Errorf("expected a|b|,got %s", caller)
	}

}
func TestEventManagerSequentially(t *testing.T) {
	_TestEventManager(t)
	_TestEventManagerWithMultipeCall(t)
	_TestEventManagerWithDistantMultipeCall(t)
}
