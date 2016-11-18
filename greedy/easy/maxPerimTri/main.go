package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

// https://www.hackerrank.com/challenges/maximum-perimeter-triangle

func runTest(sticks []int, debug bool) {
	stickCount := len(sticks)

	sort.Sort(sort.Reverse(sort.IntSlice(sticks)))
	if debug {
		fmt.Printf("%v\n", sticks)
	}

	for stickIndex1 := 0; stickIndex1 < stickCount-2; stickIndex1++ {
		s1 := sticks[stickIndex1]
		for stickIndex2 := stickIndex1 + 1; stickIndex2 < stickCount-1; stickIndex2++ {
			s2 := sticks[stickIndex2]
			for stickIndex3 := stickIndex2 + 1; stickIndex3 < stickCount; stickIndex3++ {
				if s1 < sticks[stickIndex3]+s2 {
					fmt.Printf("%d %d %d\n", sticks[stickIndex3], s2, s1)
					return
				}
			}
		}
	}
	fmt.Printf("-1\n")
}

func main() {
	filenames := []string{"T1.txt", "T2.txt", "T3.txt"}
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

		stickCount := nextInt(scanner)
		sticks := nextIntArray(scanner, stickCount)
		runTest(sticks, opened)

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

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
