package main

// https://www.hackerrank.com/challenges/richie-rich
import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func recreateString(digits []byte, digitCount int) string {
	var offset int

	var offsetDirection = 2
	var halfwayCount = digitCount / 2

	var buffer bytes.Buffer
	for recreatePassIndex := 0; recreatePassIndex < 2; recreatePassIndex++ {
		for digitIndex := 0; digitIndex < halfwayCount; digitIndex++ {
			buffer.WriteRune(rune(digits[offset] + '0'))
			offset += offsetDirection
		}

		if recreatePassIndex == 0 && digitCount%2 != 0 {
			buffer.WriteRune(rune(digits[offset] + '0'))
		}
		offsetDirection = -2
		offset--
	}

	return buffer.String()
}

func countNonMatches(digits []byte) []int {
	nonMatchingIndices := make([]int, 0, 0)
	for digitIndex := 0; digitIndex < len(digits); digitIndex += 2 {
		if digits[digitIndex] != digits[digitIndex+1] {
			nonMatchingIndices = append(nonMatchingIndices, digitIndex)
		}
	}
	return nonMatchingIndices
}

func runTest(reader *bufio.Reader, log bool) (string, bool) {
	digitCount := nextNumber(reader)
	maxChanges := nextNumber(reader)
	if log {
		fmt.Printf("Digit count: %d\nMax changes: %d\n", digitCount, maxChanges)
	}

	digitsSize := digitCount + digitCount%2
	digits := make([]byte, digitsSize, digitsSize)

	lengthRead := readDigits(reader, digitCount, digits)
	if lengthRead < digitCount {
		return "", false
	}
	if log {
		input := recreateString(digits, digitCount)
		fmt.Printf("Digits: %s %v\n", input, digits)
	}
	nonMatchingIndices := countNonMatches(digits)
	if len(nonMatchingIndices) > maxChanges {
		return "", false
	}

	// Changes that can be made that aren't necessary to make string palendromic
	spareChanges := maxChanges - len(nonMatchingIndices)

	for digitsIndex := 0; digitsIndex < digitCount; digitsIndex += 2 {
		if spareChanges > 1 && digits[digitsIndex] < 9 &&
			digits[digitsIndex] == digits[digitsIndex+1] {
			// The digits earlier in the string that are palendromic, but less than 9, can
			//  be swapped for 9
			digits[digitsIndex] = 9
			digits[digitsIndex+1] = 9
			spareChanges -= 2
		} else if digits[digitsIndex] != digits[digitsIndex+1] {
			// Need to make a change here. Set the palendromic pair to the highest
			if digits[digitsIndex] < digits[digitsIndex+1] {
				digits[digitsIndex] = digits[digitsIndex+1]
			} else {
				digits[digitsIndex+1] = digits[digitsIndex]
			}
			// Now changes, but could it be improved?  If this isn't set 9, and there is a spare
			//  change, can change both to 9
			if spareChanges > 0 && digits[digitsIndex] < 9 {
				digits[digitsIndex] = 9
				digits[digitsIndex+1] = 9
				spareChanges--
			}
		}
	}

	if digitCount%2 == 1 && spareChanges > 0 {
		if log {
			fmt.Printf("Digit count: %d\nDivided by 2: %d\n", digitCount, digitCount/2)
		}
		digits[(digitCount/2)*2] = 9
	}

	output := recreateString(digits, digitCount)

	return output, true
}

func main() {
	filenames := []string{"richie_t1.txt", "richie_t2.txt",
		"richie_t3.txt", "richie_t4.txt", "richie_t5.txt",
		"richie_t6.txt", "richie_t7.txt"}

	for fileIndex, filename := range filenames {
		f, reader, opened := openFile(filename)
		defer f.Close()

		if fileIndex > 0 {
			fmt.Printf("\n")
		}

		if opened {
			fmt.Printf("File: %s\n", filename)
		}

		output, success := runTest(reader, opened)

		if !success {
			fmt.Printf("-1\n")
		} else {
			fmt.Printf("%s\n", output)
		}
	}
}

func openFile(filename string) (*os.File, *bufio.Reader, bool) {
	f, err := os.Open(filename)
	var opened = true
	if err != nil {
		f = os.Stdin
		opened = false
	}
	reader := bufio.NewReader(f)

	return f, reader, opened
}

func nextWord(reader *bufio.Reader) string {
	var buff bytes.Buffer

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			return buff.String()
		}
		if r == ' ' || r == '\n' || r == '\t' {
			return buff.String()
		}
		buff.WriteRune(r)
	}
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

func nextNumber(reader *bufio.Reader) int {
	word := nextWord(reader)
	num, _ := strconv.ParseInt(word, 10, 32)

	return int(num)
}
