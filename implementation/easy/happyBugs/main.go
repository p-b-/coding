package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// https://www.hackerrank.com/challenges/happy-ladybugs

const (
	yesStr = "YES"
	noStr  = "NO"
)

func freqCount(r []rune, log bool) []int {
	var freq = make([]int, 27, 27)

	for _, r := range r {
		if r == '_' {
			freq[26]++
		} else {
			var letter = int(r - 'A')
			freq[letter]++
		}
	}
	return freq
}

// happyBugString checks a completely full string - no space for bugs to move
func happyBugString(r []rune, log bool) bool {
	if len(r) == 1 {
		// only one bug is unhappy
		return false
	}
	lastElementIndex := len(r) - 1
	if r[0] != r[1] {
		// First bug must match next bug to be happy
		return false
	}
	for index := 1; index <= lastElementIndex; index++ {
		if r[index] == r[index-1] {
			continue
		}
		if index == lastElementIndex {
			return false
		}
		if r[index] != r[index+1] {
			return false
		}
	}
	return true
}

func runTest(placeCount int, layout string, log bool) string {
	if log {
		fmt.Printf("%d %s\n", placeCount, layout)
	}
	if placeCount == 0 {
		// No unhappy ladybugs
		return yesStr
	}

	fc := freqCount([]rune(layout), log)

	for index, c := range fc {
		// 26 is for _
		if index < 26 && c == 1 {
			if log {
				fmt.Printf("Char %d count %d - returning no\n", index, c)
			}
			// unhappy bug
			return noStr
		}
	}

	if fc[26] == 0 {
		if log {
			fmt.Printf("No space to move, check string\n")
		}
		// Ensure layout is happy
		if happyBugString([]rune(layout), log) {
			return yesStr
		}
		return noStr
	}
	return yesStr
}

func main() {
	filenames := []string{"T1.txt"}
	var start time.Time
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		if opened {
			start = time.Now()
		}
		testCount := nextInt(scanner)
		for testIndex := 0; testIndex < testCount; testIndex++ {
			placeCount := nextInt(scanner)
			layout := nextWord(scanner)
			fmt.Printf("%s\n", runTest(placeCount, layout, opened))
		}
		//}
		if opened {
			elapsed := time.Since(start)
			log.Printf("Test took: %s\n", elapsed)
		}
		if !opened {
			return
		}
	}
}

func openFile(filename string) (*os.File, bool, *bufio.Scanner) {
	file, err := os.Open(filename)
	var opened = true
	if err != nil {
		file = os.Stdin
		opened = false
	}

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	scanner.Split(bufio.ScanWords)

	return file, opened, scanner
}

func nextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	word := scanner.Text()
	num, _ := strconv.ParseInt(word, 10, 32)
	return int(num)
}

func nextInt64(scanner *bufio.Scanner) int64 {
	scanner.Scan()
	word := scanner.Text()
	num, _ := strconv.ParseInt(word, 10, 64)
	return int64(num)
}

func nextUInt64(scanner *bufio.Scanner) uint64 {
	scanner.Scan()
	word := scanner.Text()
	num, _ := strconv.ParseUint(word, 10, 64)
	return num
}

func nextWord(scanner *bufio.Scanner) string {
	scanner.Scan()
	word := scanner.Text()
	return word
}

func nextIntArray(scanner *bufio.Scanner, elementCount int) []int {
	elements := make([]int, 0, elementCount)
	for elementCount > 0 {
		elements = append(elements, nextInt(scanner))
		elementCount--
	}

	return elements
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
