package anagram

type Anagram struct {
	Words []string
	Normal string
}

func (ar *Anagram) Eq(other *Anagram) bool {
	return false
}

func GenAnagrams(words []string) []*Anagram{
	return []*Anagram{}
}
