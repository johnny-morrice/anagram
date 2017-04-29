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
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/johnny-morrice/anagram"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find anagrams",
	Long: `Find anagrams and print to stdout.
	If --word not provided, read from stdin.`,
	Run: func(cmd *cobra.Command, args []string) {
		var input io.Reader
		if wordFile == "" {
			input = os.Stdin
		} else {
			file, err := os.Open(wordFile)

			if err != nil {
				die(err)
			}

			defer file.Close()

			input = file
		}

		words := slurpLines(input)

		if friendly {
			words = cleanLines(words)
		}

		anas := anagram.GenAnagrams(words)

		sort.Sort(anagram.ByNumber(anas))

		for _, a := range anas {
			fmt.Println(strings.Join(a.Words, " "))
		}
	},
}

func init() {
	RootCmd.AddCommand(findCmd)
}

func slurpLines(r io.Reader) []string {
	lines := []string{}
	sc := bufio.NewScanner(r)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

// Being nice to users.
func cleanLines(lines []string) []string {
	trimmed := make([]string, len(lines))

	for i, l := range lines {
		trimmed[i] = strings.Trim(l, " \t\n")
	}

	uniq := map[string]struct{}{}

	for _, s := range trimmed {
		uniq[s] = struct{}{}
	}

	out := []string{}
	for s, _ := range uniq {
		out = append(out, s)
	}

	return out
}
