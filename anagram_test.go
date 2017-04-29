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

func TestGenAnagrams(t *testing.T) {
	expect := []*Anagram{
		&Anagram{
			Words: []string{"act", "cat"},
			Normal: "act",
		},
		&Anagram{
			Words: []string{"ape", "pea"},
			Normal: "aep",
		},
		&Anagram{
			Words: []string{"dog", "god"},
			Normal: "dgo",
		},
	}

	results := GenAnagrams(words)

	if len(expect) != len(results) {
		t.Error("Mismatch in expect/result length")
	}

	for i, e := range results {
		if r := results[i]; !e.Eq(r) {
			t.Error("Unexpected result:", r)
		}
	}
}
