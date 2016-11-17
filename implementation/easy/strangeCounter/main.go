package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// https://www.hackerrank.com/challenges/strange-code
func runTest(t int64, debug bool) int64 {

	// timeBlock represents how many seconds in the current block, leading up to
	//  the block we are interested int
	timeBlockLength := int64(3)

	// if in first block [1=3,2=2,3=1] , we don't have to take further blocks into account
	// However, if the time is larger than this, iterate taking off each block length
	//  as we go
	for t > timeBlockLength {
		// Take block length off.  If the time instance we are interested in is in the second block
		//  [4=6,5=5, 6=4,7=3,8=2,9=1] then we wouldn't have to take off any more
		t -= timeBlockLength
		// Double the length of the block, to remove this block if necessary in next iteration
		timeBlockLength *= 2
	}
	// If in first block [1=3,2=2,3=1] then we need to reverse it by subtracting from block length
	//  ie 3-t.  Need to add one as we aren't returning 0-based counter  (return 1,2 or 3, not 0,1,or 2)
	//   timeBlockLength would still be three if in first block
	// If in second block, [4=6,5=5, 6=4,7=3,8=2,9=1], then we reverse it by subtracting from block length
	//  ie 6-t, still needing to add 1.

	return timeBlockLength - t + 1
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
		t := nextInt64(scanner)
		fmt.Printf("%d\n", runTest(t, opened))
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
