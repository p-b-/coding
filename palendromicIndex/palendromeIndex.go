package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"
)

// https://www.hackerrank.com/challenges/palindrome-index

func readPalendromeFromString(source string, digitCount int, buffer []byte) int {
	var count int

	var offset int

	var halfCount = digitCount / 2
	var offsetDirection = 2

	var runes = []rune(source)
	var sourceIndex int

	for readPass := 0; readPass < 2; readPass++ {
		for digitIndex := 0; digitIndex < halfCount; digitIndex++ {
			runeDigit := runes[sourceIndex]
			sourceIndex++
			var digit = byte(runeDigit - 'a')

			buffer[offset] = digit
			offset += offsetDirection
			count++
		}
		if readPass == 0 && digitCount%2 != 0 {
			runeDigit := runes[sourceIndex]
			sourceIndex++
			var digit = byte(runeDigit - 'a')

			buffer[offset] = digit
			buffer[offset+1] = digit
			count++
		}
		offset--
		offsetDirection = -2
	}

	return count
}

func stringIsPalendromic(line string) bool {
	runeCount := utf8.RuneCountInString(line)
	lRuneIndex := 0
	rRuneIndex := runeCount - 1
	for len(line) > 0 {
		if lRuneIndex == rRuneIndex {
			return true
		}
		rLeft, widthLeft := utf8.DecodeRuneInString(line)
		rRight, widthRight := utf8.DecodeLastRuneInString(line)
		if rLeft != rRight {
			return false
		}
		lRuneIndex++
		rRuneIndex--

		if lRuneIndex > rRuneIndex {
			return true
		}

		line = line[widthLeft : len(line)-widthRight]
	}
	return true
}

func runTest(scanner *bufio.Scanner, log bool) (int, bool) {
	scanner.Scan()
	line := scanner.Text()

	return testString(line, log)
}

func testString(line string, log bool) (int, bool) {
	digitCount := utf8.RuneCountInString(line)
	if digitCount == 0 {
		return 0, false
	}
	digitsSize := digitCount + digitCount%2
	digits := make([]byte, digitsSize, digitsSize)

	readPalendromeFromString(line, digitCount, digits)
	if log {
		fmt.Printf("%v\n", digits)
	}

	var leftIndex = 0
	var rightIndex = digitCount - 1

	for index := 0; index < digitsSize; index += 2 {
		if digits[index] == digits[index+1] {
			leftIndex++
			rightIndex--
			continue
		}
		if log {
			fmt.Printf("Mismatch at index %d (l%d r%d), %d!=%d\n", index, leftIndex, rightIndex, digits[index], digits[index+1])
		}
		if leftIndex+1 == rightIndex {
			if log {
				fmt.Printf("Middle character, remove the middle at %d\n", leftIndex)
			}
			return leftIndex, true
		}
		if digits[index+2] == digits[index+1] &&
			digits[index] == digits[index+3] {

			//		fmt.Printf("Cannot decide easily\n")

			var newLine string
			if leftIndex == 0 {
				newLine = line[1:]
			} else {
				newLine = line[0:leftIndex] + line[leftIndex+1:]
			}
			if stringIsPalendromic(newLine) {
				return leftIndex, true
			}
			return rightIndex, true
		} else if digits[index+2] == digits[index+1] {
			if log {
				fmt.Printf("Removing character at %d\n", leftIndex)
			}
			return leftIndex, true
		} else if digits[index] == digits[index+3] {
			if log {
				fmt.Printf("Removing character at %d\n", rightIndex)
			}
			return rightIndex, true
		}

		leftIndex++
		rightIndex--
	}
	return 0, false
}

func main() {
	filenames := []string{"palindex_t1.txt"}
	for fileIndex, filename := range filenames {
		f, scanner, opened := openFile(filename)
		defer f.Close()

		if fileIndex > 0 {
			fmt.Printf("\n")
		}

		if opened {
			fmt.Printf("File: %s\n", filename)
		}

		testCount := nextNumber(scanner)
		for testIndex := 0; testIndex < testCount; testIndex++ {
			output, success := runTest(scanner, opened)
			if !success {
				fmt.Printf("-1\n")
			} else {
				fmt.Printf("%d\n", output)
			}
		}
	}
}

func openFile(filename string) (*os.File, *bufio.Scanner, bool) {
	f, err := os.Open(filename)
	var opened = true
	if err != nil {
		f = os.Stdin
		opened = false
	}
	buf := make([]byte, 0, 64*1024)
	scanner := bufio.NewScanner(f)
	scanner.Buffer(buf, 1024*1024)
	scanner.Split(bufio.ScanWords)

	return f, scanner, opened
}

func readInputDigit(reader *bufio.Reader) (byte, bool) {
	r, _, err := reader.ReadRune()
	if err == io.EOF {
		return 0, false
	}
	if r == ' ' || r == '\n' || r == '\t' {
		return 0, false
	}
	var digit = byte(r - '0')

	return digit, true
}

func readDigits(reader *bufio.Reader, digitCount int, buffer []byte) int {
	var count int

	var offset int

	var halfCount = digitCount / 2
	var offsetDirection = 2

	for readPass := 0; readPass < 2; readPass++ {
		for digitIndex := 0; digitIndex < halfCount; digitIndex++ {
			digit, ok := readInputDigit(reader)
			if !ok {
				return count
			}
			buffer[offset] = digit
			offset += offsetDirection
			count++
		}
		if readPass == 0 && digitCount%2 != 0 {
			digit, ok := readInputDigit(reader)
			if !ok {
				return count
			}
			buffer[offset] = digit
			buffer[offset+1] = digit
			count++
		}
		offset--
		offsetDirection = -2
	}

	return count
}

func nextNumber(scanner *bufio.Scanner) int {
	scanner.Scan()
	word := scanner.Text()
	num, _ := strconv.ParseInt(word, 10, 32)

	return int(num)
}
