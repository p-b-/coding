package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type mandragora struct {
	health          uint64
	healthSummation uint64
}

var mandragoras []mandragora

// ByHealth is a type used to sort the mandragoras
type ByHealth []mandragora

// Len used to sort mandragoras
func (h ByHealth) Len() int {
	return len(h)
}

// Swap used to sort mandragoras
func (h ByHealth) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

// Less used to sort mandragoras
func (h ByHealth) Less(i int, j int) bool {
	return h[i].health < h[j].health
}

func readMandragoras(scanner *bufio.Scanner, mandragoraCount int) {
	mandragoras = make([]mandragora, 0, mandragoraCount)
	for mi := 0; mi < mandragoraCount; mi++ {
		h := nextUInt64(scanner)
		var m mandragora
		m.health = h
		mandragoras = append(mandragoras, m)
	}
}

func calcSummedHealth(log bool) {
	var mLen = len(mandragoras)
	var sum = mandragoras[mLen-1].health
	mandragoras[mLen-1].healthSummation = sum

	for mi := mLen - 2; mi >= 0; mi-- {
		m := mandragoras[mi]
		sum += m.health
		mandragoras[mi].healthSummation = sum
	}
	if log {
		fmt.Println("Health sum:")
		for mi, m := range mandragoras {
			fmt.Printf("Mandragora: %d, Health: %d, Sum: %d\n", mi, m.health, m.healthSummation)
		}
	}
}

func runTest(scanner *bufio.Scanner, mandragoraCount int, log bool) uint64 {
	if mandragoraCount == 0 {
		return 0
	}
	readMandragoras(scanner, mandragoraCount)
	if log {
		fmt.Println("Presort:")
		for mi, m := range mandragoras {
			fmt.Printf("Mandragora: %d, Health: %d\n", mi, m.health)
		}
	}
	sort.Sort(ByHealth(mandragoras))

	if log {
		fmt.Println("Sorted:")
		for mi, m := range mandragoras {
			fmt.Printf("Mandragora: %d Health: %d\n", mi, m.health)
		}
	}

	calcSummedHealth(log)

	var strength = uint64(1)
	var experience = strength * mandragoras[0].healthSummation

	for mi := 0; mi < mandragoraCount; mi++ {
		var newExperience uint64

		if mi < mandragoraCount-1 {
			newExperience = (strength + 1) * mandragoras[mi+1].healthSummation
		} else {
			newExperience = strength + 1
		}

		if newExperience > experience {
			if log {
				fmt.Printf("Incrementing strength, new experience would be %d, currently %d\n", newExperience, experience)
			}
			experience = newExperience
			strength++
		} else if log {
			fmt.Printf("Not incrementing strength, new experience would be %d, currently %d\n", newExperience, experience)
		}
		if log {
			fmt.Printf("Strength: %d, Experience: %d\n", strength, newExperience)
		}
	}

	return experience
}

func main() {
	filenames := []string{"t1.txt"}
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
			mandragoraCount := nextInt(scanner)

			answer := runTest(scanner, mandragoraCount, opened)
			fmt.Printf("%d\n", answer)
		}

		//		answer := runTest(t1, t2, n, opened)
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

func nextUInt64(scanner *bufio.Scanner) uint64 {
	scanner.Scan()
	word := scanner.Text()
	num, _ := strconv.ParseUint(word, 10, 64)
	return num
}

func readConfig(scanner *bufio.Scanner) (int, int, int) {
	t1 := nextInt(scanner)
	t2 := nextInt(scanner)
	n := nextInt(scanner)

	return t1, t2, n
}
