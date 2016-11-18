package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/game-of-thrones

func freqCount(r []rune, log bool) []int {
	var freq = make([]int, 26, 26)

	for _, r := range r {
		var letter = int(r - 'a')
		freq[letter]++
	}

	return freq
}

func runTest(line string, log bool) string {
	if log {
		fmt.Printf("Line: %s\n", line)
	}

	var r = []rune(line)

	var lineIsOdd = (len(r) % 2) == 1

	if log {
		fmt.Printf("r= %v\n", r)
	}

	freq := freqCount(r, log)

	if log {
		fmt.Printf("Freq: %v\n", freq)
	}

	// Count amount of odd frequencies - there can only be one (but only if there is an odd
	//  number of charactesr)
	var oddCount int
	for _, f := range freq {
		if f%2 == 1 {
			if !lineIsOdd || oddCount == 1 {
				return "NO"
			}
			oddCount++
		}
	}
	return "YES"
}

func main() {
	filenames := []string{"T1.txt", "T2.txt", "T3.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		testLine := nextWord(scanner)
		fmt.Printf("%s\n", runTest(testLine, opened))
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
