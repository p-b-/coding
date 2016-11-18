package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

// https://www.hackerrank.com/challenges/kangaroo

func runTest(x1 int, v1 int, x2 int, v2 int, log bool) string {
	if x1 == x2 {
		return "YES"
	} else if v1 == v2 {
		return "NO"
	}

	if log {
		fmt.Printf("Kanga one %d, %d\n", x1, v1)
		fmt.Printf("Kanga two %d, %d\n", x2, v2)
	}
	// k1 = n1*v1 + x1
	// k2 = n2*v2 + x2
	// to coincide timewise, n1=n2 as well as k1=k2
	// n1=n2=n

	// n*v1+x1 = n*v2+x2
	// n=  (x2-x1)/(v1-v2)

	// x2-x1 !=0 (checked at entry to this function)

	n := float64(x2-x1) / float64(v1-v2)
	if log {
		fmt.Printf("n = %f\n", n)
	}
	if n < 0 {
		return "NO"
	}
	if math.Abs(n-math.Floor(n)) < 0.0001 {
		return "YES"
	}

	return "NO"
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
		x1 := nextInt(scanner)
		v1 := nextInt(scanner)
		x2 := nextInt(scanner)
		v2 := nextInt(scanner)
		fmt.Printf("%s\n", runTest(x1, v1, x2, v2, opened))
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
