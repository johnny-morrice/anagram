package anagram

import (
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
		t.Error("Mismatch in expect/actual length")
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
