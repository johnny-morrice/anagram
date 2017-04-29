package anagram

import (
	"fmt"
	"sort"
	"strings"

	leven "github.com/texttheater/golang-levenshtein/levenshtein"
)

// Anagram represents a collection of words which are anagrams of one another.
type Anagram struct {
	// Words are each anagrams of one another.
	// Words in Anagrams returned by library functions are sorted alphabetically.
	Words []string
	// Normal is the normalised representation of the anagram.
	Normal string
}

// ByNormal Anagram sorting.
type ByNormal []*Anagram

func (a ByNormal) Len() int { return len(a)}
func (a ByNormal) Swap(i, j int) { a[i], a[j] = a[j], a[i]}

func (a ByNormal) Less(i, j int) bool {
	return a[i].Normal < a[j].Normal
}

// ByNumber sorts anagrams by the number of Words.
type ByNumber []*Anagram

func (a ByNumber) Len() int { return len(a)}
func (a ByNumber) Swap(i, j int) { a[i], a[j] = a[j], a[i]}

func (a ByNumber) Less(i, j int) bool {
	return len(a[i].Words) < len(a[j].Words)
}

// Normalize an Anogram.  Returns an error if the provided words are not an anogram.
func (ar *Anagram) Normalize() error {
	if l := len(ar.Words); l < 2 {
		return fmt.Errorf("Cannot create anagram with %v words", l)
	}

	normals := make([]string, len(ar.Words))

	for i, w := range ar.Words {
		n := SortNormal(w)
		normals[i] = n
	}

	first := normals[0]
	for _, n := range normals[1:] {
		if n != first {
			return fmt.Errorf("Not anagrams: '%v', '%v'", first, n)
		}
	}

	ar.Normal = first
	sort.Strings(ar.Words)
	return nil
}

// Eq returns true if the other Anagram is equal.
func (ar *Anagram) Eq(other *Anagram) bool {
	if ar.Normal != other.Normal {
		return false
	}

	for i, w := range ar.Words {
		if w != other.Words[i] {
			return false
		}
	}

	return true
}

// Rank each combination of words in the anagram.
func (ar *Anagram) Rank(ranker Ranker) []Ranking {
	indices := bruteCombintations2(len(ar.Words))
	ranks := make([]Ranking, len(indices))

	for i, wij := range indices {
		r := Ranking{
			A: ar.Words[wij[0]],
			B: ar.Words[wij[1]],
		}
		r.Rank = ranker.Rank(r.A, r.B)
		ranks[i] = r
	}

	return ranks
}

// TODO Could use arrays.
func bruteCombintations2(n int) [][]int {
	len := fact(n) / (fact(2) * fact(n - 2))
	out := make([][]int, len)

	mark := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			out[mark] = []int{i, j}
			mark++
		}
	}

	return out
}

func fact(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

// Ranking represents the interestingness of an anagram.
type Ranking struct {
	A string
	B string
	Rank int
}

// ByRank sorts Rankings by their Rank.
type ByRank []Ranking

func (r ByRank) Len() int { return len(r)}
func (r ByRank) Swap(i, j int) { r[i], r[j] = r[j], r[i]}

func (r ByRank) Less(i, j int) bool {
	return r[i].Rank < r[j].Rank
}

// Ranker represents a string distance function.
type Ranker interface {
	Rank(a, b string) int
}

// LevenshteinRanker provides Levenshtein string distance ranking.
type LevenshteinRanker struct {
	Options leven.Options
}

func DefaultLevenshteinRanker() Ranker {
	return LevenshteinRanker{Options: leven.DefaultOptions}
}

type hammingRanker struct {}

func (hr hammingRanker) Rank(a, b string) int {
	if len(a) != len(b) {
		panic("Computed Hamming distance on strings of unequal length")
	}

	arunes := []rune(a)
	brunes := []rune(b)

	score := 0
	for i, ar := range arunes {
		if br := brunes[i]; br != ar {
			score++
		}
	}

	return score
}

// HammingRanker provides Hamming string distance ranking.
func HammingRanker() Ranker {
	return hammingRanker{}
}

func (lr LevenshteinRanker) Rank(a, b string) int {
	ar := []rune(a)
	br := []rune(b)
	return leven.DistanceForStrings(ar, br, lr.Options)
}

// Find all anagrams among the input words.
func Find(words []string) []*Anagram{
	buckets := map[string][]string{}

	for _, w := range words {
		normal := SortNormal(w)
		buckets[normal] = append(buckets[normal], w)
	}

	anas := []*Anagram{}
	for normal, ws := range buckets {
		if len(ws) == 1 {
			continue
		}

		a := &Anagram{
			Words: ws,
			Normal: normal,
		}
		sort.Strings(a.Words)
		anas = append(anas, a)
	}

	return anas
}

// SortNormal provides a normalised representation of a word.
func SortNormal(word string) string {
	parts := strings.Split(word, "")
	sort.Strings(parts)
	return strings.Join(parts, "")
}
