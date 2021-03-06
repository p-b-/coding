package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/repeated-string

func runTest(wordToRepeat string, n uint64, log bool) uint64 {
	lenOfWord := uint64(len(wordToRepeat))
	countOfA := uint64(strings.Count(wordToRepeat, "a"))

	completeRepeats := n / lenOfWord
	if log {
		fmt.Printf("Complete repeats %d\n", completeRepeats)
	}

	var totalCount uint64
	totalCount = completeRepeats * countOfA

	incompleteRepeatLength := n % lenOfWord
	if incompleteRepeatLength != 0 {
		wordToRepeat = wordToRepeat[:incompleteRepeatLength]
		totalCount += uint64(strings.Count(wordToRepeat, "a"))
	}

	return totalCount
}

func main() {
	var filenames = []string{"T1.txt", "T2.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		wordToRepeat := nextWord(scanner)
		n := nextUInt64(scanner)
		fmt.Printf("%d\n", runTest(wordToRepeat, n, opened))
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
