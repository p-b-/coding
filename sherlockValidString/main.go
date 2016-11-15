package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/sherlock-and-valid-string

func freqCount(r []rune, log bool) []int {
	var freq = make([]int, 26, 26)

	for _, r := range r {
		var letter = int(r - 'a')
		freq[letter]++
	}

	if log {
		fmt.Printf("Letter Freq: %v\n", freq)
	}
	return freq
}

func createFreqMap(freq []int, log bool) map[int]int {
	retFreq := make(map[int]int)
	for _, n := range freq {
		if n == 0 {
			continue
		}
		retFreq[n]++
	}

	if log {
		for k, v := range retFreq {
			fmt.Printf("Letter frequency %d occurs %d times\n", k, v)
		}
	}

	return retFreq
}

func runTest(line string, log bool) string {
	if log {
		fmt.Printf("Line: %s\n", line)
	}
	freq := freqCount([]rune(line), log)

	freqMap := createFreqMap(freq, log)

	if len(freqMap) == 1 {
		if log {
			fmt.Printf("Only one frequency\n")
		}
		return "YES"
	}
	if len(freqMap) > 2 {
		if log {
			fmt.Printf("More than 2 frequencies, freqCount: %d\n", len(freqMap))
		}
		return "NO"
	}

	var overOneOccurence bool
	var frequencies []int
	// Can only have one count that has more than one frequency
	for k, v := range freqMap {
		if v > 1 {
			if overOneOccurence {
				if log {
					fmt.Printf("Already got a frequency with more than one occurence\n")
				}
				return "NO"
			}
			overOneOccurence = true
		}
		frequencies = append(frequencies, k)
	}
	if log {
		for _, f := range frequencies {
			fmt.Printf("Freq %d\n", f)
		}
	}

	if frequencies[0] == 1 || frequencies[1] == 1 {
		if log {
			fmt.Printf("Decrementing frequency count 1 leads to only one frequency\n")
		}
		return "YES"
	}

	if absInt(frequencies[1]-frequencies[0]) > 1 {
		if log {
			fmt.Printf("Difference between counts is larger than one: %d\n", absInt(frequencies[1]-frequencies[0]))
		}
		return "NO"
	}

	return "YES"
}

func main() {
	filenames := []string{"T1.txt", "T2.txt", "T3.txt", "T4.txt", "T5.txt", "T6.txt", "T7.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		testLine := nextWord(scanner)
		fmt.Printf("%s\n", runTest(testLine, opened))
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
