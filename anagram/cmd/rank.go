// Copyright © 2017 Johnny Morrice <john@functorama.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/johnny-morrice/anagram"
)

var rankFunc string

// rankCmd represents the rank command
var rankCmd = &cobra.Command{
	Use:   "rank",
	Short: "Rank anagrams",
	Long: `Find and rank anagrams according to a string distance function`,
	Run: func(cmd *cobra.Command, args []string) {
		anas := find()
		var ranker anagram.Ranker
		switch rankFunc {
		case LEVEN:
			ranker = anagram.DefaultLevenshteinRanker()
		case HAMMING:
			ranker = anagram.HammingRanker()
		default:
			die(fmt.Errorf("Unsupported string distance function: %v", rankFunc))
		}
		ranks := anagram.RankAll(anas, ranker)
		sort.Sort(sort.Reverse(anagram.ByRank(ranks)))

		for _, r := range ranks {
			fmt.Printf("%v\t%v\t%v\n", r.Rank, r.A, r.B)
		}
	},
}

func init() {
	RootCmd.AddCommand(rankCmd)

	rankCmd.Flags().StringVar(&rankFunc, "distance", LEVEN, "String distance function")
}

const LEVEN = "levenshtein"
const HAMMING = "hamming"
