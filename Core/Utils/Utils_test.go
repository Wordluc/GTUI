package Utils

import (
	"slices"
	"testing"
)

func TestDiffOk(t *testing.T) {
	from := []int{1, 2, 3}
	to := []int{1, 2, 3, 4, 5}
	result := []int{4, 5}
	if !slices.Equal(GetDiff(from, to), result) {
		t.Errorf("Diff failed")
	}
}
func TestDiffEmptyResult(t *testing.T) {
	from := []int{1, 2, 3}
	to := []int{1, 2,}
	result := []int{}
	if !slices.Equal(GetDiff(from, to), result) {
		t.Errorf("Diff failed")
	}
}
