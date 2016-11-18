package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/funny-string

func runeDifference(runeSlice []rune, index1 int, index2 int) int {
	diff := int(runeSlice[index1] - runeSlice[index2])
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func runTest(line string, log bool) string {
	if log {
		fmt.Printf("Line: %s\n", line)
	}

	var r = []rune(line)
	lineLen := len(r)
	lastIndex := lineLen - 1

	for index := 0; index < lastIndex; index++ {
		diff1 := runeDifference(r, index+1, index)
		diff2 := runeDifference(r, lastIndex-index-1, lastIndex-index)

		if diff1 != diff2 {
			return "Not Funny"
		}
	}
	return "Funny"
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

			fmt.Printf("%s\n", answer)
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
