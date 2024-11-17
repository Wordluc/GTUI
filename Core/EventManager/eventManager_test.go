package EventManager

import (
	"testing"
	"time"
)
const (
	test1 EventType=9
	test2 EventType=10
	test3 EventType = iota
)
type callerForTest struct {
	name string
}
func testTime(t time.Time, timeToWait int) (int, bool) {
	res := (time.Now().UnixNano() - t.UnixNano()) / int64(time.Millisecond)
	return int(res), int(res) == timeToWait
}
func TestEventManager(t *testing.T) {
	Setup()
	var isok bool
	var res int
	timer := time.Now()

	Subscribe(test1, 2, func(comp []any) {
		res, isok = testTime(timer, 2)
	})
	Call(test1, nil)
	time.Sleep(100 * time.Millisecond)
	if !isok {
		t.Errorf("expected to wait 2, got %v", res)
	}
}
func TestEventManagerWithMultipeCall(t *testing.T) {
	var timer time.Time
	var isok bool
	var ncalls int
	var res int
	var caller string
	timer = time.Now()
	e:=Subscribe(test2, 500, func(comp []any) {
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
	if e:=Call(test2, []any{callerA});e!=nil{
		t.Errorf("expected no error,got %v", e)
	}
	time.Sleep(100 * time.Millisecond)
	Call(test2, []any{callerB})
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
func TestEventManagerWithDistantMultipeCall(t *testing.T) {
	var ncalls int
	var caller string
	e:=Subscribe(test3, 300, func(comp []any) {
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
	if e:=Call(test3, []any{callerA});e!=nil{
		t.Errorf("expected no error,got %v", e)
	}
	time.Sleep(550 * time.Millisecond)
	Call(test3, []any{callerB})
	time.Sleep(400 * time.Millisecond)
	if ncalls != 2 {
		t.Errorf("expected 2 calls,got %d", ncalls)
	}
	if caller!="a|b|"{
		t.Errorf("expected a|b|,got %s", caller)
	}
}

