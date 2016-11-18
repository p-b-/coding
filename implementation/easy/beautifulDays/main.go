package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// https://www.hackerrank.com/challenges/beautiful-days-at-the-movies

func reverseNumber(toReverse int) int {
	n := fmt.Sprintf("%d", toReverse)
	var buffer bytes.Buffer
	for i := len(n) - 1; i >= 0; i-- {
		buffer.WriteByte(n[i])
	}
	num, _ := strconv.ParseInt(buffer.String(), 10, 32)

	return int(num)
}

func runTest(dateStart int, dateEnd int, divisor int, log bool) int {
	count := 0

	for d := dateStart; d < dateEnd; d++ {
		reversed := reverseNumber(d)

		diff := absInt(reversed - d)

		if diff%divisor == 0 {
			count++
			if log {
				fmt.Printf("%d - %d = %d, mod %d = %d\n", d, reversed, diff, divisor, diff%divisor)
			}
		}
	}
	return count
}

func main() {
	filenames := []string{"T1.txt"}
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
		dateStart, dateEnd, divisor := readTestConfig(scanner)
		fmt.Printf("%d\n", runTest(dateStart, dateEnd, divisor, opened))
		if opened {
			elapsed := time.Since(start)
			log.Printf("Test took: %s\n", elapsed)
		}
		if !opened {
			return
		}
	}
}

func readTestConfig(scanner *bufio.Scanner) (int, int, int) {
	i := nextInt(scanner)
	j := nextInt(scanner)
	k := nextInt(scanner)
	return i, j, k
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
