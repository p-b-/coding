package main

 // https://www.hackerrank.com/contests/w26/challenges/best-divisor
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func runTest(n int, debug bool) int {
	if debug {
		fmt.Printf("n: %d\n", n)
	}
  if n==1 {
    return 1
  } else if n==2 {
    return 2
  }
	if m == 0 || n == 0 {
		return 0
	}
	x := m / 2
	y := n / 2

	if m%2 != 0 {
		x++
	}
	if n%2 != 0 {
		y++
	}

	return x * y
}

func main() {
	filenames := []string{"T1.txt" }
	var start time.Time
	for fi, fn := range filenames {
		f, fileOpened, scanner := openFile(fn)

		if fileOpened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		if fileOpened {
			start = time.Now()
		}

		n := nextInt(scanner)
		fmt.Printf("%d\n", runTest n, fileOpened))
		if fileOpened {
			elapsed := time.Since(start)
			log.Printf("Test took: %s\n", elapsed)
		}
		if !fileOpened {
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

func nextIntArray(scanner *bufio.Scanner, elementCount int) []int {
	elements := make([]int, 0, elementCount)
	for elementCount > 0 {
		elements = append(elements, nextInt(scanner))
		elementCount--
	}

	return elements
}

func nextUInt64Array(scanner *bufio.Scanner, elementCount int) []uint64 {
	elements := make([]uint64, 0, elementCount)
	for elementCount > 0 {
		elements = append(elements, nextUInt64(scanner))
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
