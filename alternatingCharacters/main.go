package main

// https://www.hackerrank.com/challenges/alternating-characters

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func runTest(line string, log bool) int {
	var lineLen = len(line)
	if log {
		if lineLen < 200 {
			fmt.Printf("Word: %s\n", line)
		} else {
			fmt.Printf("Word (len %d): %s\n", lineLen, line[0:200])
		}
	}

	// Contest specification states that there will be at least one character
	// Lines of 1 character are already alternating
	if lineLen < 2 {
		return 0
	}
	answer := 0

	ch1, w1 := utf8.DecodeRuneInString(line)
	line = line[w1:]
	for len(line) > 0 {
		ch2, w2 := utf8.DecodeRuneInString(line)

		line = line[w2:]
		if log {
			fmt.Printf("Character 1 %c\n", ch1)
			fmt.Printf("Character 2 %c\n\n", ch2)
		}

		if ch1 == ch2 {
			answer++
		} else {
			ch1 = ch2
		}
	}

	return answer
}

func main() {
	filenames := []string{"t1.txt"}
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
		for ti := 0; ti < testCount; ti++ {
			line := nextWord(scanner)
			answer := runTest(line, opened)

			fmt.Printf("%d\n", answer)
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
