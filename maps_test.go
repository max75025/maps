package maps

import (
	"sort"
	"testing"
)

var m1 = map[int]int{1: 2, 2: 4, 4: 8, 8: 16}
var m2 = map[int]string{1: "2", 2: "4", 4: "8", 8: "16"}

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func sliceEqual[E comparable](s1, s2 []E) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestKeys(t *testing.T) {
	want := []int{1, 2, 4, 8}

	got1 := Keys(m1)
	sort.Ints(got1)
	if !sliceEqual(got1, want) {
		t.Errorf("Keys(%v) = %v, want %v", m1, got1, want)
	}

	got2 := Keys(m2)
	sort.Ints(got2)
	if !sliceEqual(got2, want) {
		t.Errorf("Keys(%v) = %v, want %v", m2, got2, want)
	}
}

func TestValues(t *testing.T) {
	got1 := Values(m1)
	want1 := []int{2, 4, 8, 16}
	sort.Ints(got1)
	if !sliceEqual(got1, want1) {
		t.Errorf("Values(%v) = %v, want %v", m1, got1, want1)
	}

	got2 := Values(m2)
	want2 := []string{"16", "2", "4", "8"}
	sort.Strings(got2)
	if !sliceEqual(got2, want2) {
		t.Errorf("Values(%v) = %v, want %v", m2, got2, want2)
	}
}
