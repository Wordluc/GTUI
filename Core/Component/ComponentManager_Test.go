package Component

import (
	"github.com/Wordluc/GTUI/Core/Utils"
	"slices"
	"testing"
)

func TestDiffWithComponent(t *testing.T) {
	b:=Button{}
	from := []IComponent{&b}
	to := []IComponent{&b}
	result := []IComponent{}
	if !slices.Equal(Utils.GetDiff(from, to), result) {
		t.Errorf("Diff failed")
	}
}
