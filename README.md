# Anagram

A fun anagram ranker for golang, inspired by this post by Mark Dominus:

[Anagram Scoring](http://blog.plover.com/2017/02/21/#anagram-scoring)

## Usage

Anagram toolkit to generate and rank anagrams according to
	string distance algorithms.

Usage:
  anagram [command]

Available Commands:
  find        Find anagrams
  help        Help about any command
  rank        Rank anagrams

Flags:
      --config string   config file (default is $HOME/.anagram.yaml)
      --friendly        chomp and uniq the input (default true)
  -t, --toggle          Help message for toggle
      --words string    dictionary file (one word per line)

## Examples

Using wordlists downloaded from [KeithV's Big English Word Lists](http://www.keithv.com/software/wlist/)

### Find anagrams

```
$ anagram find --words wlist_match10.txt
$ cat wlist_match{8..10}.txt | anagram find
```

### Rank anagrams

```
$ anagram rank --words wlist_match10.txt
$ cat ~/workspace/words/wlist_match{8..10}.txt | anagram rank
```