package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/string-similarity

func runTest(testLine string, log bool) int64 {
	log = false

	//	r := createLCSRow(testLine)
	fmt.Printf("Line: %s\n", testLine)

	answer := sumLCSRowFaster(testLine)
	fmt.Printf("Answer: %d\n", answer)
	return answer
	//fmt.Printf("Row sum %d\n", sum)
	//	return sum
	//	return sum

	/*	testLineLen := int64(len(testLine))
		sum = testLineLen

		//fmt.Printf("Row %v\n", r)

		m := createLCSMatrix(testLine, testLine)
		if log {
			outputLCSMatrix(testLine, testLine, m)
		}
		for i := int64(2); i <= testLineLen; i++ {
			if log {
				fmt.Printf("Consider from first row, col %d\n", i)
				fmt.Printf("First element %d\n", m[i][1])
			}
			for l := int64(0); l <= testLineLen-i; l++ {
				if m[l+1][i+l] != 0 {
					//				fmt.Printf("Element [%d][%d] = %d\n", l+1, i+l, m[l+1][i+l])
					sum++
				} else {
					break
				}
			}
		}
		fmt.Printf("Matrix sum %d\n", sum)

		return sum*/
}

func main() {
	//	testLCS()
	filenames := []string{"T1.txt"} //, "T2.txt"}
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
			testLine := nextWord(scanner)
			fmt.Printf("%d\n", runTest(testLine, opened))
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

func sumLCSRowFaster(s string) int64 {
	var lenS = len(s)

	var row = make([]int, lenS)

	// Slice of slices, each nested slice having a count to jump forward to next reoccurence of that letter
	jumpRows := make([][]int, 26)
	jumpRowsComplete := make([]bool, 26)
	jumpRowsFirstIndex := make([]int, 26)
	for ji := 0; ji < 26; ji++ {
		jumpRows[ji] = make([]int, lenS)
	}

	var charToConsider = s[0]
	var jumpRowIndex = charToConsider - 'a'
	var sum int64
	row[0] = 1 // First column always matches first letter
	sum = 1
	lastCharPos := 0
	var jumpRowForChar = jumpRows[jumpRowIndex]
	for i := 1; i < lenS; i++ {
		if s[i] == charToConsider {
			jumpRowForChar[lastCharPos] = i
			lastCharPos = i

			row[i] = 1
			sum++
		}
	}
	jumpRowsComplete[jumpRowIndex] = true
	jumpRowsFirstIndex[jumpRowIndex] = 0

	fmt.Printf("JR %c start %d %v\n", charToConsider, jumpRowsFirstIndex[charToConsider-'a'], jumpRows[charToConsider-'a'])
	//	fmt.Printf("R1 %v\n", row)

	for i := 1; i < lenS; i++ {
		charToConsider = s[i]
		jumpRowIndex = charToConsider - 'a'
		jumpRowForChar = jumpRows[jumpRowIndex]

		var loopEnd = lenS - i
		if !jumpRowsComplete[jumpRowIndex] {
			jumpRowsFirstIndex[jumpRowIndex] = i
			lastCharPos := -1
			for j := 0; j < loopEnd; j++ {
				if charToConsider == s[j+i] {
					if lastCharPos != -1 {
						jumpRowForChar[lastCharPos] = j + i
					}
					lastCharPos = j + i
				}
				if row[j] == 0 {
					// Only consider comparisons if matched against first letter (first loop)
					// A negative element indicates that there has been a break in the match
					continue
				}
				if charToConsider == s[j+i] {
					sum++
				} else {
					row[j] = 0
				}
			}
			jumpRowsComplete[jumpRowIndex] = true
			fmt.Printf("JR %c start %d %v\n", charToConsider, jumpRowsFirstIndex[charToConsider-'a'], jumpRows[charToConsider-'a'])

		} else {
			var startJ = jumpRowsFirstIndex[jumpRowIndex] - i
			fmt.Printf("i=%d, start J[%d] =%d-%d=%d\n", i, jumpRowIndex, jumpRowsFirstIndex[jumpRowIndex], i, startJ)
			var stopLooping bool
			for j := startJ; j < loopEnd && !stopLooping; j++ {
				nextOffset := jumpRowForChar[j+i]
				if nextOffset == 0 {
					stopLooping = true
				}
				if j < 0 {
					j = nextOffset - i
					continue
				}
				fmt.Printf("c = %c, j=%d\n", charToConsider, j)
				if row[j] == 0 {
					// Only consider comparisons if matched against first letter (first loop)
					// A negative element indicates that there has been a break in the match
					j = nextOffset - i
					continue
				}
				if charToConsider == s[j+i] {
					sum++
				} else {
					row[j] = 0
				}
				j = nextOffset - i
			}
		}
	}

	return sum
}

func sumLCSRow(s string) int64 {
	var lenS = len(s)
	var row = make([]int, lenS)

	var firstLetter = s[0]
	row[0] = 1 // First column always matches first letter
	for i := 1; i < lenS; i++ {
		if s[i] == firstLetter {
			row[i] = 1
		}
	}
	//	fmt.Printf("R1 %v\n", row)

	for i := 1; i < lenS; i++ {
		var charToConsider = s[i]
		//	fmt.Printf("   ")
		for j := 0; j < lenS; j++ {
			/*if i+j < lenS {
				fmt.Printf(" %c", s[j+i])
			}*/
			if i+j >= lenS || row[j] <= 0 {
				// Only consider comparisons if matched against first letter (first loop)
				// A negative element indicates that there has been a break in the match
				continue
			}
			//		fmt.Printf("Considering i=%d, char=%c, j=%d\n", i, charToConsider, j)
			if charToConsider == s[j+i] {
				row[j]++
			} else {
				row[j] = -row[j]
			}
		}
		//	fmt.Println()
		//	fmt.Printf("R%d %v\n", i, row)
	}

	//fmt.Printf("Row %v\n", row)

	var sum int64
	for _, e := range row {
		if e > 0 {
			sum += int64(e)
		} else if e < 0 {
			sum -= int64(e)
		}
	}

	return sum
}
