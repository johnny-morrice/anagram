package anagram

import (
	"reflect"
	"testing"
)

var words []string = []string{
	"god",
	"dog",
	"act",
	"cat",
	"pea",
	"ape",
}

var catAnagram *Anagram = &Anagram{
	Words: []string{"act", "cat"},
	Normal: "act",
}

var apeAnagram *Anagram = &Anagram{
	Words: []string{"ape", "pea"},
	Normal: "aep",
}

var dogAnagram *Anagram = &Anagram{
	Words: []string{"dog", "god"},
	Normal: "dgo",
}

func TestGenAnagrams(t *testing.T) {
	expect := []*Anagram{
		catAnagram,
		apeAnagram,
		dogAnagram,
	}

	actual := GenAnagrams(words)
	// Cheeky...
	SortAnagrams(actual)

	if len(expect) != len(actual) {
		panic("Mismatch in expect/actual length")
	}

	for i, a := range actual {
		if e := expect[i]; !e.Eq(a) {
			t.Error("Unexpected Anagram:", a)
		}
	}
}

func TestNormalize(t *testing.T) {
	expect := catAnagram

	actual := &Anagram{
		Words: []string{"cat", "act"},
	}
	err := actual.Normalize()

	if err != nil {
		t.Error("Unexpected error")
	}

	if !expect.Eq(actual) {
		t.Error("Unexpected Anagram:", actual)
	}
}

func TestNormalizeFailure(t *testing.T) {
	actual := []*Anagram{
		&Anagram{
			Words: []string{"duck", "cow"},
		},
		&Anagram{
			Words: []string{"hello"},
		},
	}

	for i, a := range actual {
		err := a.Normalize()

		if err == nil {
			t.Error("Expected error at", i)
		}
	}
}

func TestBruteCombinations2(t *testing.T) {
	expect := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{0, 3},
		[]int{1, 2},
		[]int{1, 3},
		[]int{2, 3},
	}

	actual := bruteCombintations2(4)

	if !reflect.DeepEqual(expect, actual) {
		t.Error("Unexpected combinations", actual)
	}
}

func TestFact(t *testing.T) {
	expect := []int{
		2,
		6,
		24,
		120,
	}

	actual := []int{
		2,
		3,
		4,
		5,
	}

	if len(expect) != len(actual) {
		panic("Mismatch in expect/actual length")
	}

	for i := 0; i < len(actual); i++ {
		e := expect[i]
		a := actual[i]
		if e != fact(a) {
			t.Error("Expected", e, "but received", a)
		}
	}
}
