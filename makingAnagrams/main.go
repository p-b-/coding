package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/making-anagrams

func freqCount(r []rune, log bool) []int {
	var freq = make([]int, 26, 26)

	for _, r := range r {
		var letter = int(r - 'a')
		freq[letter]++
	}

	return freq
}

func runTest(line1 string, line2 string, log bool) int {
	if log {
		fmt.Printf("Line 1: %s\nLine 2: %s\n", line1, line2)
	}

	var r1 = []rune(line1)
	var r2 = []rune(line2)

	if log {
		fmt.Printf("r1= %v\n", r1)
		fmt.Printf("r2= %x\n", r2)
	}

	freq1 := freqCount(r1, log)
	freq2 := freqCount(r2, log)

	if log {
		fmt.Printf("Freq1: %v\n", freq1)
		fmt.Printf("Freq2: %v\n", freq2)
	}

	// Count amount of odd frequencies - there can only be one (but only if there is an odd
	//  number of charactesr)
	var changeCount int
	for index := 0; index < len(freq1); index++ {
		changeCount += absInt(freq1[index] - freq2[index])
	}
	return changeCount
}

func main() {
	filenames := []string{"T2.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		testLine1 := nextWord(scanner)
		testLine2 := nextWord(scanner)
		fmt.Printf("%d\n", runTest(testLine1, testLine2, opened))
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

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
