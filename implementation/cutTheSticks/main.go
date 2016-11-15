package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// https://www.hackerrank.com/challenges/cut-the-sticks

func runTest(stickSizes []int, log bool) {
	if len(stickSizes) == 0 {
		fmt.Printf("0\n")
		return
	}
	if log {
		fmt.Printf("%v\n", stickSizes)
	}

	sort.Ints(stickSizes)
	if log {
		fmt.Printf("Sorted: %v\n", stickSizes)
	}

	for len(stickSizes) > 0 {
		if log {
			fmt.Printf("%v\n", stickSizes)
		}
		fmt.Printf("%d\n", len(stickSizes))
		lowestSize := stickSizes[0]

		for i := 0; i < len(stickSizes); i++ {
			stickSizes[i] -= lowestSize
		}

		newStart := 1
		for i := 0; i < len(stickSizes); i++ {
			if stickSizes[i] == 0 {
				newStart = i + 1
			} else {
				break
			}
		}
		if newStart > len(stickSizes) {
			return
		}
		if log {
			fmt.Printf("New start = %d\n", newStart)
		}
		stickSizes = stickSizes[newStart:]
	}
}

func main() {
	var filenames = []string{"T5.txt"} //"T1.txt", "T2.txt", "T3.txt", "T4.txt", "T5.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		stickCount := nextInt(scanner)
		stickSizes := nextIntArray(scanner, stickCount)

		runTest(stickSizes, opened)
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
