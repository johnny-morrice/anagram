package anagram

import (
	"fmt"
	"sort"
	"strings"
)

type Anagram struct {
	Words []string
	Normal string
}

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
	len := fact(n) / fact(n - 2)
	out := make([][]int, len)

	mark := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			out[mark] = []int{i, j}
		}
	}

	return out
}

// Based on algorithm T from TAOCP Volume 4A "Generating All Combinations"
// func knuthCombinations(n, t int) [][]int {
// 	if t <= 0 || t >= n || n < 3 {
// 		panic(fmt.Sprintf("invalid arguments: %v %v", n, t)
// 	}
//
// 	len := fact(n) / fact(n - t)
// 	out := make([][]int, len)
// 	mark := 0
//
// 	c := make([]int, t + 2)
// 	for j := 0; j < t; j++ {
// 		c[j] = i
// 	}
//
// 	c[t + 1] = n
// 	c[t + 2] = 0
//
// 	var x int
// 	for j := t;; {
// 		copy(out, c, t, mark)
// 		mark++
//
// 		if j > -1 {
// 			x = j
// 			goto T6
// 		}
//
// 		if c[0] + 1 < c[1] {
// 			c[0] = c[0] + 1
// 			continue
// 		} else {
// 			j := 1
// 		}
//
// 		for {
// 			c[j - 1] = j - 2
// 			x = c[j] + 1
// 			if x == c[j + 1] {
// 				j++
// 			} else {
// 				break
// 			}
// 		}
//
// 		if j <= t {
// 			break
// 		}
//
// T6:
// 		c[j] = x
// 		j--
// 	}
//
// 	return out
// }

func copy(out [][]int, comb []int, t, mark int) {
	cpy := make([]int, t)
	for i := 0; i < t; i++ {
		cpy[i] = comb[i]
	}
	out[mark] = cpy
}

func fact(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

type Ranking struct {
	A string
	B string
	Rank int
}

type Ranker interface {
	Rank(a, b string) int
}

func GenAnagrams(words []string) []*Anagram{
	buckets := map[string][]string{}

	for _, w := range words {
		normal := SortNormal(w)
		buckets[normal] = append(buckets[normal], w)
	}

	anas := []*Anagram{}
	for _, ws := range buckets {
		// TODO could optimize
		a := &Anagram{Words: ws}
		err := a.Normalize()

		if err != nil {
			panic(err)
		}

		anas = append(anas, a)
	}

	return anas
}

func SortAnagrams(anas []*Anagram) {
	sort.Sort(byAnagram(anas))
}

type byAnagram []*Anagram

func (a byAnagram) Len() int { return len(a)}
func (a byAnagram) Swap(i, j int) { a[i], a[j] = a[j], a[i]}

func (a byAnagram) Less(i, j int) bool {
	return a[i].Normal < a[j].Normal
}

// TODO can optimize surely.
func SortNormal(word string) string {
	parts := strings.Split(word, "")
	sort.Strings(parts)
	return strings.Join(parts, "")
}
