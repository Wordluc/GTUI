package Component

import (
	"testing"
)
func Test_InsertComponent(t *testing.T) {
  man := Create(15, 15, 5)
	comp := CreateButton(0, 0, 10, 10, "prova")
	err := comp.insertingToMap(man)
	if err != nil {
		t.Errorf("Error inserting component: %v", err)
	}
  if (len((*man.Map)[2]) != 3) {
		t.Errorf("Size not correct")
	}
	if c, e := man.Search(0, 0); e != nil || c != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(5, 0); e != nil || c != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(0,5); e != nil || c != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(5,5); e != nil || c != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(10,0); e == nil || c == comp {
		t.Errorf("Component inserted in wrong place")
	}
	if c, e := man.Search(0,10); e == nil || c == comp {
		t.Errorf("Component inserted in wrong place")
	}
}
func Test_InsertComponent2(t *testing.T) {
  man := Create(15, 15, 5)
	comp := CreateButton(0, 3, 10, 10, "prova")
	err := comp.insertingToMap(man)
	if err != nil {
		t.Errorf("Error inserting component: %v", err)
	}
  if (len((*man.Map)[2]) != 3) {
		t.Errorf("Size not correct")
	}
	if c, e := man.Search(0,0); e == nil || c == comp {
		t.Errorf("Component inserted in wrong place")
	}
	if c, e := man.Search(0,2); e == nil || c == comp {
		t.Errorf("Component inserted in wrong place")
	}
	if c, e := man.Search(13,3); e == nil || c == comp {
		t.Errorf("Component inserted in wrong place")
	}
	if c, e := man.Search(5,4); e != nil || c != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(3,3); e != nil || c != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(7,11); e != nil || c != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(7,12); e != nil || c != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(7,13); e == nil || c == comp {
		t.Errorf("Component inserted in wrong place")
	}
}
