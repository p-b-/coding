package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var seq map[int]*big.Int

func calculateSequenceAt(n int) *big.Int {
	// t(i+2) = t(i) + t(i+1)^2

	if value, ok := seq[n]; ok {
		//		fmt.Printf("Already know T%d, %s\n", n, seq[n])
		return value
	}
	valueNMinus1 := calculateSequenceAt(n - 1)
	valueNMinus2 := calculateSequenceAt(n - 2)
	calcValue := new(big.Int).Set(valueNMinus1)
	calcValue = calcValue.Mul(calcValue, calcValue)
	calcValue = calcValue.Add(calcValue, valueNMinus2)

	//	fmt.Printf("Setting T%d to %s\n", n, calcValue)
	seq[n] = calcValue

	return seq[n]
}

func runTest(t1 int, t2 int, n int, log bool) *big.Int {

	seq = make(map[int]*big.Int)

	var T1, T2 big.Int
	T1.SetInt64(int64(t1))
	T2.SetInt64(int64(t2))
	seq[0] = &T1
	seq[1] = &T2

	answer := calculateSequenceAt(n - 1)
	if log {
		for i := 0; i < n-1; i++ {
			fmt.Printf("T%d is %s\n", i, seq[i])
		}
	}

	return answer
}

func main() {
	filenames := []string{"fib_t1.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)
		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}

		t1, t2, n := readConfig(scanner)
		answer := runTest(t1, t2, n, opened)
		fmt.Printf("%s\n", answer)
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
	scanner.Split(bufio.ScanWords)

	return file, opened, scanner
}

func nextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	word := scanner.Text()
	num, _ := strconv.ParseInt(word, 10, 32)
	return int(num)
}

func readConfig(scanner *bufio.Scanner) (int, int, int) {
	t1 := nextInt(scanner)
	t2 := nextInt(scanner)
	n := nextInt(scanner)

	return t1, t2, n
}
