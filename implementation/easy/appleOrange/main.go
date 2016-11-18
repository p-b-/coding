package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// https://www.hackerrank.com/challenges/apple-and-orange

func hitsHouse(x int, houseStartX int, houseEndX int) bool {
	if x < houseStartX || x > houseEndX {
		return false
	}
	return true
}

func runTest(houseStartX int, houseEndX int, appleTreeX int, orangeTreeX int, apples []int, oranges []int, opened bool) {
	var applesHitHouse int
	var orangesHitHouse int
	for _, ax := range apples {
		if hitsHouse(appleTreeX+ax, houseStartX, houseEndX) {
			applesHitHouse++
		}
	}
	for _, ax := range oranges {
		if hitsHouse(orangeTreeX+ax, houseStartX, houseEndX) {
			orangesHitHouse++
		}
	}
	fmt.Printf("%d\n%d\n", applesHitHouse, orangesHitHouse)
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
		houseStartX := nextInt(scanner)
		houseEndX := nextInt(scanner)
		appleTreeX := nextInt(scanner)
		orangeTreeX := nextInt(scanner)
		appleCount := nextInt(scanner)
		orangeCount := nextInt(scanner)
		apples := nextIntArray(scanner, appleCount)
		oranges := nextIntArray(scanner, orangeCount)
		runTest(houseStartX, houseEndX, appleTreeX, orangeTreeX, apples, oranges, opened)
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
