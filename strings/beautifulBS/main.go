package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/beautiful-binary-string
// Remove 010, count minimum steps

func runTest(line string, log bool) int {
	if log {
		fmt.Printf("Line: %s\n", line)
	}

	count := 0

	asBytes := []byte(line)

	loop := true

	var bytes01010 = []byte{'0', '1', '0', '1', '0'}
	var bytes010 = []byte{'0', '1', '0'}
	for loop {
		indexOf01010 := bytes.Index(asBytes, bytes01010)
		if indexOf01010 != -1 {
			asBytes[indexOf01010+2] = '1'
			count++
		} else {
			loop = false
		}
	}
	loop = true
	for loop {
		indexOf010 := bytes.Index(asBytes, bytes010)
		if indexOf010 != -1 {
			asBytes[indexOf010+2] = '1'
			count++
		} else {
			loop = false
		}
	}

	return count
}

func main() {
	//	testLCS()
	filenames := []string{"T1.txt", "T2.txt", "T3.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		nextWord(scanner)
		testLine := nextWord(scanner)

		fmt.Printf("%d\n", runTest(testLine, opened))
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

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
