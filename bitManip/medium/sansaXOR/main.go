package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// https://www.hackerrank.com/challenges/sansa-and-xor

func countFreq(elements []uint64) map[uint64]int {
	f := make(map[uint64]int)
	for _, e := range elements {
		f[e]++
	}
	return f
}

func runTest(elements []uint64, debug bool) uint64 {
	elementCount := len(elements)
	if debug {
		fmt.Printf("Count: %d\n", elementCount)
		fmt.Printf("Elements: %v\n", elements)
	}

	f := countFreq(elements)
	for k, v := range f {
		fmt.Printf("%d occurs %d\n", k, v)
	}

	if elementCount == 0 {
		return 0
	} else if elementCount%2 == 0 {
		// Everything cancels out
		return 0
	}
	return elements[0] ^ elements[elementCount-1]
}

func main() {
	filenames := []string{"T1.txt", "T2.txt"}
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
			elementCount := nextInt(scanner)
			elements := nextUInt64Array(scanner, elementCount)
			fmt.Printf("%d\n", runTest(elements, opened))
		}
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

func nextUInt64Array(scanner *bufio.Scanner, elementCount int) []uint64 {
	elements := make([]uint64, 0, elementCount)
	for elementCount > 0 {
		elements = append(elements, nextUInt64(scanner))
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
