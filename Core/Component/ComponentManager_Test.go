package Component

import (
	"GTUI/Core/Utils"
	"slices"
	"testing"
)

func Test_InsertComponent(t *testing.T) {
	man := Create(15, 15, 5)
	comp := CreateButton(0, 0, 10, 10, "prova")
	err:=man.Add(comp)
	if err != nil {
		t.Errorf("Error inserting component: %v", err)
	}
	if len((*man.Map)[2]) != 3 {
		t.Errorf("Size not correct")
	}
	if c, e := man.Search(1, 1); e != nil || c[0] != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(5, 1); e != nil || c[0] != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(1, 5); e != nil || c[0] != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(5, 5); e != nil || c[0] != comp {
		t.Errorf("Component not inserted")
	}
	if c, _ := man.Search(1, 10); len(c) != 0 {
		t.Errorf("Component inserted in wrong place")
	}
	if c, e := man.Search(0, 0);e!=nil || len(c) != 0 {
		t.Errorf("Component inserted in wrong place")
	}
}

func Test_InsertComponent2(t *testing.T) {
	man := Create(15, 15, 5)
	comp := CreateButton(0, 4, 10, 10, "prova")
	err:=man.Add(comp)
	if err != nil {
		t.Errorf("Error inserting component: %v", err)
	}
	if len((*man.Map)[2]) != 3 {
		t.Errorf("Size not correct")
	}
	if c, _ := man.Search(1, 1); len(c) != 0 {
		t.Errorf("Component inserted in wrong place")
	}
	if c, _ := man.Search(1, 2); len(c) != 0 {
		t.Errorf("Component inserted in wrong place")
	}
	if c, _ := man.Search(13, 3); len(c) != 0 {
		t.Errorf("Component inserted in wrong place")
	}
	if c, _ := man.Search(7, 15); len(c) != 0 {
		t.Errorf("Component inserted in wrong place")
	}
	if c, e := man.Search(5, 5); e != nil || c[0] != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(5, 5); e != nil || c[0] != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(7, 12); e != nil || c[0] != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(7, 11); e != nil || c[0] != comp {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(7, 12); e != nil || c[0] != comp {
		t.Errorf("Component not inserted")
	}
}

func Test_InsertComponent3(t *testing.T) {
	man := Create(15, 20, 5)
	man.Add(CreateButton(2, 0, 4, 3, "prova"))
	if c, _ := man.Search(0, 0); len(c) != 0{
		t.Errorf("Component inserted in wrong place")
	}
	if c, _ := man.Search(3, 1); len(c) != 1{
		t.Errorf("Component not inserted")
	}
	if c, _ := man.Search(5, 1); len(c) != 1{
		t.Errorf("Component not inserted")
	}
}
func Test_InsertComponentOverlapping(t *testing.T) {
	man := Create(15, 20, 5)
	man.Add(CreateButton(0, 0, 10, 10, "prova"))
	man.Add(CreateButton(0, 3, 10, 10, "prova"))
	if c, e := man.Search(1, 1); e != nil || c[0] == nil {
		t.Errorf("Component not inserted")
	}
	if c, e := man.Search(1, 4); e != nil || len(c) != 2 {
		t.Errorf("Component not inserted")
	}
}
func TestDiffWithComponent(t *testing.T) {
	b:=Button{}
	from := []IComponent{&b}
	to := []IComponent{&b}
	result := []IComponent{}
	if !slices.Equal(Utils.GetDiff(from, to), result) {
		t.Errorf("Diff failed")
	}
}
