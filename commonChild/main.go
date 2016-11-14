package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/common-child

func freqCount(r []rune, log bool) []int {
	var freq = make([]int, 26, 26)

	for _, r := range r {
		var letter = int(r - 'A')
		freq[letter]++
	}

	return freq
}

func convertToBits(line string) []uint32 {
	digits := make([]uint32, 0, len(line))
	for _, r := range line {
		digit := uint32(1 << uint32(r-'A'))
		digits = append(digits, digit)
	}
	return digits
}

func collateBits(line string) uint32 {
	var collatedBits uint32
	for _, r := range line {
		collatedBits |= uint32(1 << uint32(r-'A'))
	}
	return collatedBits
}

func regenerateString(line string, bitsToRemove uint32) string {
	var buffer bytes.Buffer
	for _, r := range line {
		bit := uint32(1 << uint32(r-'A'))
		if (bitsToRemove & bit) == 0 {
			buffer.WriteRune(r)
		}
	}

	return buffer.String()
}

func runTest(line1 string, line2 string, log bool) int {
	if log {
		fmt.Printf("Line 1: %s\nLine 2: %s\n", line1, line2)
	}

	var bits1 = collateBits(line1)
	var bits2 = collateBits(line2)

	var bitsInCommon = bits1 & bits2

	var bitsToRemove1 = bits1 & ^bitsInCommon
	var bitsToRemove2 = bits2 & ^bitsInCommon

	var line1Regenerated = regenerateString(line1, bitsToRemove1)
	var line2Regenerated = regenerateString(line2, bitsToRemove2)

	if log {
		fmt.Printf("b1= %026b\n", bits1)
		fmt.Printf("b2= %026b\n", bits2)
		fmt.Printf("r1= %026b\n", bitsToRemove1)
		fmt.Printf("r2= %026b\n", bitsToRemove2)

		fmt.Printf("l1= %s\n", line1Regenerated)
		fmt.Printf("l2= %s\n", line2Regenerated)
	}
	answer := lcsLength(line1Regenerated, line2Regenerated)
	//answer := lcs(line1Regenerated, line2Regenerated)
	if log {
		fmt.Printf("Common Child: %d\n", answer)
	}

	return answer
}

func main() {
	filenames := []string{"T5.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		testLine1 := nextWord(scanner)
		testLine2 := nextWord(scanner)
		fmt.Printf("%d\n", runTest(testLine1, testLine2, opened))
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

func lcs(s1 string, s2 string) string {
	var m = make([][]int, 1+len(s1))
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, 1+len(s2))
	}
	longest := 0
	longestStartsAtX := 0
	for x := 1; x < 1+len(s1); x++ {
		for y := 1; y < 1+len(s2); y++ {
			if s1[x-1] == s2[y-1] {
				m[x][y] = m[x-1][y-1] + 1
				if m[x][y] > longest {
					longest = m[x][y]
					longestStartsAtX = x
				}
			}
		}
	}

	outputLCSMatrix(s1, s2, m)
	return s1[longestStartsAtX-longest : longestStartsAtX]
}

func lcsLength(s1 string, s2 string) int {
	var m = make([][]int, 1+len(s1))
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, 1+len(s2))
	}
	for x := 1; x < 1+len(s1); x++ {
		for y := 1; y < 1+len(s2); y++ {
			if s1[x-1] == s2[y-1] {
				m[x][y] = m[x-1][y-1] + 1
			} else {
				m[x][y] = max(m[x][y-1], m[x-1][y])
			}
		}
	}
	return m[len(s1)][len(s2)]
}

func max(m1 int, m2 int) int {
	if m1 > m2 {
		return m1
	}
	return m2
}

/*function LCSLength(X[1..m], Y[1..n])
  C = array(0..m, 0..n)
  for i := 0..m
     C[i,0] = 0
  for j := 0..n
     C[0,j] = 0
  for i := 1..m
      for j := 1..n
          if X[i] = Y[j]
              C[i,j] := C[i-1,j-1] + 1
          else
              C[i,j] := max(C[i,j-1], C[i-1,j])
  return C[m,n]*/

func testLCS() {
	//	a := "ABCXYZAYAGV"
	//b := "XYZABCB"
	a := "ABCABABCD"
	b := "XYZABCABABCDXYZ"

	s := lcs(a, b)

	fmt.Printf("s1: %s, s2: %s, LCS:%s\n", a, b, s)
}
