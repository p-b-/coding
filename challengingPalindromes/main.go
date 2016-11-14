package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/challenging-palindromes

func freqCount(r []rune, log bool) []int {
	var freq = make([]int, 26, 26)

	for _, r := range r {
		var letter = int(r - 'a')
		freq[letter]++
	}

	return freq
}

func runTest(line1 string, line2 string, log bool) string {
	if log {
		fmt.Printf("Line 1: %s\nLine 2: %s\n", line1, line2)
	}

	var r1 = []rune(line1)
	var r2 = []rune(line2)

	if log {
		fmt.Printf("r1= %v\n", r1)
		fmt.Printf("r2= %x\n", r2)
	}

	freq1 := freqCount(r1, log)
	freq2 := freqCount(r2, log)

	if log {
		fmt.Printf("Freq1: %v\n", freq1)
		fmt.Printf("Freq2: %v\n", freq2)
	}

	// Count amount of odd frequencies - there can only be one (but only if there is an odd
	//  number of charactesr)
	var changeCount int
	for index := 0; index < len(freq1); index++ {
		changeCount += absInt(freq1[index] - freq2[index])
	}
	return ""
}

func main() {
	testLCS()
	/*	filenames := []string{"T1.txt"}
		for fi, fn := range filenames {
			f, opened, scanner := openFile(fn)

			if opened {
				defer f.Close()
				if fi > 0 {
					fmt.Println()
				}
				fmt.Printf("File: %s\n", fn)
			}
			testCount := nextInt(scanner)
			for testIndex := 0; testIndex < testCount; testIndex++ {
				testLine1 := nextWord(scanner)
				testLine2 := nextWord(scanner)
				fmt.Printf("%s\n", runTest(testLine1, testLine2, opened))
			}
			if !opened {
				return
			}
		}*/
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

func outputLCSMatrix(s1 string, s2 string, matrix [][]int) {
	fmt.Printf("  y ")
	for _, r := range s2 {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
	for i := 0; i < len(matrix); i++ {
		if i > 0 {
			fmt.Printf("%c ", s1[i-1])
		} else {
			fmt.Printf("x ")
		}
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func LCS(s1 string, s2 string) string {
	var m = make([][]int, 1+len(s1))
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, 1+len(s2))
	}
	longest := 0
	x_longest := 0
	for x := 1; x < 1+len(s1); x++ {
		for y := 1; y < 1+len(s2); y++ {
			if s1[x-1] == s2[y-1] {
				m[x][y] = m[x-1][y-1] + 1
				if m[x][y] > longest {
					longest = m[x][y]
					x_longest = x
				}
			}
		}
	}

	outputLCSMatrix(s1, s2, m)
	return s1[x_longest-longest : x_longest]
}

func testLCS() {
	//	a := "ABCXYZAYAGV"
	//b := "XYZABCB"
	a := "ABCABABCD"
	b := "XYZABCABABCDXYZ"

	s := LCS(a, b)

	fmt.Printf("s1: %s, s2: %s, LCS:%s\n", a, b, s)
}
