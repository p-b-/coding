package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// https://www.hackerrank.com/challenges/floyd-city-of-blinding-lights

type query struct {
	node1 int
	node2 int
}

var queries []*query

func createMatrix(count int) [][]int {
	m := make([][]int, count, count)

	for i := 0; i < count; i++ {
		m[i] = make([]int, count, count)
		for j := 0; j < count; j++ {
			if i != j {
				m[i][j] = 0x7fffffff
			}
		}
	}

	return m
}

func displayMatrix(m [][]int, nodeCount int) {
	fmt.Printf("         to --->\n       ")
	for j := 0; j < nodeCount; j++ {
		fmt.Printf("%03d ", j+1)
	}
	fmt.Printf("\n")
	for i := 0; i < nodeCount; i++ {
		fmt.Printf("%03d    ", i+1)
		for j := 0; j < nodeCount; j++ {
			if m[i][j] == 0x7fffffff {
				fmt.Printf("x   ")

			} else {
				fmt.Printf("%03d ", m[i][j])
			}
		}
		fmt.Printf("\n")
	}
}

func runTest(m [][]int, nodeCount int, log bool) {
	//	displayMatrix(m, nodeCount)
	for k := 0; k < nodeCount; k++ {
		for i := 0; i < nodeCount; i++ {
			for j := 0; j < nodeCount; j++ {
				a := m[i][k] + m[k][j]
				//		fmt.Printf("from %d to %d, current %d a=%d\n", i+1, j+1, m[i][j], a)
				if m[i][j] > a {
					m[i][j] = a
				}
			}
		}
		//		fmt.Printf("\nk=%d\n", k+1)
		//		displayMatrix(m, nodeCount)
	}
}

func main() {
	filenames := []string{"T2.txt"}
	var start time.Time
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}
		if opened {
			start = time.Now()
		}
		distanceMatrix, nodeCount := readTestConfig(scanner)
		runTest(distanceMatrix, nodeCount, opened)
		for _, q := range queries {
			d := distanceMatrix[q.node1][q.node2]
			if d == 0x7fffffff {
				d = -1
			}
			fmt.Printf("%d\n", d)
		}
		if opened {
			elapsed := time.Since(start)
			log.Printf("Test took: %s\n", elapsed)
		}
		if !opened {
			return
		}
	}
}

func readTestConfig(scanner *bufio.Scanner) ([][]int, int) {
	nodeCount := nextInt(scanner)
	edgeCount := nextInt(scanner)

	m := createMatrix(nodeCount)

	for edgeIndex := 0; edgeIndex < edgeCount; edgeIndex++ {
		edgeIndex1 := nextInt(scanner) - 1
		edgeIndex2 := nextInt(scanner) - 1
		distance := nextInt(scanner)

		m[edgeIndex1][edgeIndex2] = distance
	}

	queryCount := nextInt(scanner)

	queries = make([]*query, 0, queryCount)
	for queryIndex := 0; queryIndex < queryCount; queryIndex++ {
		n1 := nextInt(scanner) - 1
		n2 := nextInt(scanner) - 1

		var q query
		q.node1 = n1
		q.node2 = n2

		queries = append(queries, &q)
	}

	return m, nodeCount
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

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
