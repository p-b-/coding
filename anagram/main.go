package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/anagram

func freqCount(r []rune, log bool) []int {
	var freq = make([]int, 26, 26)

	for _, r := range r {
		var letter = int(r - 'a')
		freq[letter]++
	}

	return freq
}

func runTest(line string, log bool) int {
	if log {
		fmt.Printf("Line: %s\n", line)
	}

	if len(line)%2 != 0 {
		return -1
	}

	var r = []rune(line)

	var a = r[0 : len(line)/2]
	var b = r[len(line)/2:]
	if log {
		fmt.Printf("A= %v\n", a)
		fmt.Printf("B= %v\n", b)
	}

	aFreq := freqCount(a, log)
	bFreq := freqCount(b, log)

	/*	if log {
		fmt.Printf("Freq A: %v\n", aFreq)
		fmt.Printf("Freq B: %v\n", bFreq)
	} */
	var changeCount int
	for index := 0; index < len(aFreq); index++ {
		if bFreq[index] > aFreq[index] {
			changeCount += bFreq[index] - aFreq[index]
		}
	}
	return changeCount
}

func main() {
	filenames := []string{"T1.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		testCount := nextInt(scanner)
		for testIndex := 0; testIndex < testCount; testIndex++ {
			line := nextWord(scanner)
			fmt.Printf("%d\n", runTest(line, opened))
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
