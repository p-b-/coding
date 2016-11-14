package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

// https://www.hackerrank.com/challenges/gem-stones

func countBits(bitArray uint32) uint32 {
	var mask = uint32(1)
	var count = 0

	for bitArray > 0 {
		if (bitArray & mask) != 0 {
			count++
		}
		bitArray >>= 1
	}
	return uint32(count)
}

func generateBitArray(line string) uint32 {
	var bitArray uint32
	for len(line) > 0 {
		ch, w := utf8.DecodeRuneInString(line)
		line = line[w:]

		digit := uint32((ch - 'a'))

		bitArray |= (1 << digit)
	}

	return bitArray
}

func runTest(scanner *bufio.Scanner, log bool) int {
	rockCount := nextInt(scanner)
	var overallBitArray uint32

	var firstRock = true

	for ri := 0; ri < rockCount; ri++ {
		rock := nextWord(scanner)
		bitArray := generateBitArray(rock)
		if log {
			fmt.Printf("Rock is %s %b\n", rock, bitArray)
		}

		if firstRock {
			overallBitArray = bitArray
			firstRock = false
		} else {
			overallBitArray &= bitArray
		}
	}
	gemCount := countBits(overallBitArray)
	if log {
		fmt.Printf("Gems: %b\n", overallBitArray)
		fmt.Printf("Gem count: %d\n", gemCount)
	}
	return int(gemCount)
}

func main() {
	filenames := []string{"T2.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}

		fmt.Printf("%d\n", runTest(scanner, opened))

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
