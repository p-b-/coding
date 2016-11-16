package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/insertionsort2

func printElements(toSort []int) {
	for i := 0; i < len(toSort); i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", toSort[i])
	}
	fmt.Printf("\n")
}

func sortElement(toSort []int, indexToSort int, log bool) {
	sortElement := toSort[indexToSort]
	for i := indexToSort - 1; i >= 0; i-- {
		if toSort[i] > sortElement {
			toSort[i+1] = toSort[i]
		} else {
			toSort[i+1] = sortElement
			return
		}
	}
	if len(toSort) > 1 && sortElement < toSort[1] {
		toSort[0] = sortElement
	}
}

func findUnsortedElement(toSort []int, log bool) int {
	if len(toSort) == 0 {
		return -1
	}

	lastValue := toSort[0]
	for i, v := range toSort {
		if i > 0 {
			if v < lastValue {
				return i
			}
		}
		lastValue = v
	}

	return -1
}

func runTest(toSort []int, log bool) {
	if len(toSort) == 0 {
		if log {
			fmt.Printf("Empty array - nothing to sort\n")
		}
		return
	}
	if log {
		fmt.Printf("Starting at: %v\n", toSort)
	}

	for toSortIndex := 1; toSortIndex < len(toSort); toSortIndex++ {
		if toSortIndex == -1 {
			if log {
				fmt.Printf("Ordered - nothing to sort\n")
			}
			return
		}
		if log {
			fmt.Printf("Sort element [%d]=%d\n", toSortIndex, toSort[toSortIndex])
		}

		sortElement(toSort, toSortIndex, log)
		printElements(toSort)

	}
}

func main() {
	filenames := []string{"T1.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		elementCount := nextInt(scanner)
		testArray := nextIntArray(scanner, elementCount)
		runTest(testArray, opened)
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
