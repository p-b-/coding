package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// https://www.hackerrank.com/contests/sears-dots-arrows/challenges/the-easy-puzzle-1/

// SortByAbsSize used to sort array
type SortByAbsSize []int64

func (s SortByAbsSize) Len() int {
	return len(s)
}

func (s SortByAbsSize) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortByAbsSize) Less(i int, j int) bool {
	return absInt64(s[i]) < absInt64(s[j])
}

func runTest(K int64, elements []int64, log bool) string {
	if log {
		//	fmt.Printf("K: %d\n", K)
		//	fmt.Printf("Array: %v\n", elements)
	}
	sort.Sort(SortByAbsSize(elements))
	if log {
		//		fmt.Printf("Sorted array: %v\n", elements)
	}

	l := lcmArray64(elements)
	if log {
		fmt.Printf("LCM of array: %d\n", l)
	}

	if l%K == 0 {
		return "YES"
	}
	return "NO"
}

func lcmArray64(elements []int64) int64 {
	lenArray := len(elements)

	currentLCM := elements[0]
	//fmt.Printf("LCM: %d\n", currentLCM)
	for i := 1; i < lenArray; i++ {
		currentLCM = lcm(currentLCM, elements[i])
		//fmt.Printf("LCM: %d\n", currentLCM)
	}
	return currentLCM
}

func lcm(a int64, b int64) int64 {

	return a * b / gcd(a, b)
}

func gcd(a int64, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
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
		testCount := nextInt(scanner)
		for testIndex := 0; testIndex < testCount; testIndex++ {
			K, elements := readTestConfig(scanner, opened)

			fmt.Printf("%s\n", runTest(K, elements, opened))
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

func readTestConfig(scanner *bufio.Scanner, log bool) (int64, []int64) {
	n := nextInt(scanner)
	K := nextInt64(scanner)
	elements := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		e := nextInt64(scanner)
		elements = append(elements, e)
	}
	return K, elements
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func absInt64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
