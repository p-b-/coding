package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/the-love-letter-mystery

func runTest(line string, log bool) int {
	if log {
		fmt.Printf("Line: %s\n", line)
	}

	var r = []rune(line)

	var lineLen = len(r)
	var lastIndex = lineLen - 1

	var changeCount = 0
	for ri := 0; ri < lineLen/2; ri++ {
		if log {
			fmt.Printf("%c compare to %c\n", r[ri], r[lastIndex-ri])
		}
		var r1 = int(r[ri] - 'a')
		var r2 = int(r[lastIndex-ri] - 'a')

		var diff = r1 - r2
		if diff < 0 {
			diff = -diff
		}
		changeCount += diff
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
