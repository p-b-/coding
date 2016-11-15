package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/sock-merchant

func getFreqMap(intArray []int) map[int]int {
	freqMap := make(map[int]int)

	for _, n := range intArray {
		freqMap[n]++
	}

	return freqMap
}

func runTest(sockColours []int, log bool) int {
	if len(sockColours) == 0 {
		return 0
	}
	if log {
		fmt.Printf("%v\n", sockColours)
	}

	freq := getFreqMap(sockColours)
	if log {
		fmt.Printf("%v\n", freq)
	}

	pairCount := 0

	for _, v := range freq {
		pairCount += v / 2
	}
	return pairCount
}

func main() {
	var filenames = []string{"T1.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		sockCount := nextInt(scanner)
		sockColours := nextIntArray(scanner, sockCount)

		fmt.Printf("%d\n", runTest(sockColours, opened))
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
