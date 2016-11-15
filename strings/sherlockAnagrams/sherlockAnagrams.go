package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBits(line string) []uint32 {
	digits := make([]uint32, 0, len(line))
	for _, r := range line {
		digit := uint32(1 << uint32(r-'a'))
		digits = append(digits, digit)
	}
	return digits
}

func getSubString(bitarray []uint32, startIndex int, endIndex int) (uint32, []uint32) {
	var bits uint32
	repeatedBits := make([]uint32, 0, endIndex-startIndex+1)
	for index := startIndex; index <= endIndex; index++ {
		if bits&bitarray[index] != 0 {
			repeatedBits = append(repeatedBits, bitarray[index])
		} else {
			bits |= bitarray[index]
		}
	}
	return bits, repeatedBits
}

func duplicateComparison(duplicates []uint32, compareDuplicates []uint32) bool {
	var l = len(duplicates)
	for duplicateIndex := 0; duplicateIndex < l; duplicateIndex++ {
		d := duplicates[duplicateIndex]
		var found = false
		for compareIndex := 0; compareIndex < l; compareIndex++ {
			if d == compareDuplicates[compareIndex] {
				compareDuplicates[compareIndex] = 0
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func runTest(line string, log bool) int {
	if log {
		fmt.Printf("anagrams of: %s\n", line)
	}
	lineLength := len(line)
	lineAsBits := convertToBits(line)

	var matchCount = 0

	for subStringLength := 1; subStringLength < lineLength; subStringLength++ {
		for offsetStart := 0; offsetStart < lineLength-subStringLength; offsetStart++ {
			subString, duplicates := getSubString(lineAsBits, offsetStart, offsetStart+subStringLength-1)
			//		fmt.Printf("Compare substring %d -> %d\n", offsetStart, offsetStart+subStringLength-1)
			for offset := offsetStart + 1; offset <= lineLength-subStringLength; offset++ {
				compareToSubString, compareDuplicates := getSubString(lineAsBits, offset, offset+subStringLength-1)

				if subString == compareToSubString && len(duplicates) == len(compareDuplicates) {
					if subStringLength == 1 || len(duplicates) == 0 {
						//	fmt.Printf("%d->%d matches %d->%d\n", offsetStart, offsetStart+subStringLength-1, offset, offset+subStringLength-1)
						matchCount++
					} else if duplicateComparison(duplicates, compareDuplicates) {
						matchCount++
					}
				}
			}
		}
	}
	return matchCount
}

func main() {
	filenames := []string{"sherlock_t1.txt", "sherlock_t2.txt"}
	for fileIndex, filename := range filenames {
		f, opened, scanner := openFile(filename)
		defer f.Close()

		if opened {
			if fileIndex > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", filename)
		}

		testCount := nextNumber(scanner)

		for testIndex := 0; testIndex < testCount; testIndex++ {
			testLine := nextWord(scanner)
			matchCount := runTest(testLine, opened)
			fmt.Printf("%d\n", matchCount)
		}
	}
}

func openFile(filename string) (*os.File, bool, *bufio.Scanner) {
	var opened = true
	file, err := os.Open(filename)

	if err != nil {
		opened = false
		file = os.Stdin
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	return file, opened, scanner
}

func nextNumber(scanner *bufio.Scanner) int {
	scanner.Scan()
	word := scanner.Text()
	num, _ := strconv.ParseInt(word, 10, 32)
	return int(num)
}

func nextWord(scanner *bufio.Scanner) string {
	scanner.Scan()
	word := scanner.Text()
	return word
}
