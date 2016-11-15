package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/jumping-on-the-clouds

func runTest(clouds []int, log bool) int {
	steps := 0

	cloudCount := len(clouds)
	lastCloud := cloudCount - 1
	pos := 0

	for pos < lastCloud {
		if pos+1 == lastCloud {
			pos++
			steps++
		} else if pos+2 == lastCloud {
			pos += 2
			steps++
		} else if clouds[pos+2] == 1 {
			// Avoid thundercloud
			steps++
			pos++
		} else {
			steps++
			pos += 2
		}
	}

	return steps
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
		cloudCount := nextInt(scanner)
		clouds := nextIntArray(scanner, cloudCount)
		fmt.Printf("%d\n", runTest(clouds, opened))
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
