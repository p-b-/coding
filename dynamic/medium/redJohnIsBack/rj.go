package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var factorials map[int]uint64
var distinctPerm map[int]map[int]uint64

var primeCount []uint64

func initialisePrimes() {
	var seive [220000]bool //default is false, which we will take to mean it is prime
	seive[0] = true
	seive[1] = true

	for offset := 2; offset < len(seive); offset++ {
		if seive[offset] == true {
			continue
		}
		// if offset < 50 {
		// 	fmt.Printf("Count up from offset %d skipping %d\n", offset*2, offset)
		// }
		for n := offset * 2; n < len(seive); n += offset {
			// if n < 50 {
			// 	fmt.Printf("Setting %d to be composite\n", n)
			// }
			seive[n] = true
		}
		// for ; offset < len(seive); offset++ {
		// 	if offset < 50 {
		// 		fmt.Printf("Consider offset: %d\n", offset)
		// 	}
		// 	if seive[offset] == false {
		// 		offset--
		// 		break
		// 	}
		// }
	}
	primeCount = make([]uint64, len(seive), len(seive))
	primeCount[0] = 0
	primeCount[1] = 0

	var countOfPrimes = 0
	for n := 2; n < len(seive); n++ {
		if seive[n] == false {
			// if n < 346 {
			// 	fmt.Printf("Counting %d as prime, making %d so far\n", n, countOfPrimes+1)
			// }
			countOfPrimes++
			primeCount[n] = uint64(countOfPrimes)
		} else {
			// if n < 346 {
			// 	fmt.Printf("%d is composite, making %d so far\n", n, countOfPrimes)
			// }
			primeCount[n] = uint64(countOfPrimes)
		}
		// if n < 500 {
		// 	fmt.Printf("Under %d there are %d primes\n", n, primeCount[n])
		// }
		// if n < 10000 {
		// 	if seive[n] {
		// 		//	fmt.Printf("%d composite\n",n)
		// 	} else {
		// 		//fmt.Printf("%d prime\n", n)
		// 		fmt.Printf("%d ", n)
		// 	}
		// }
	}
	//fmt.Println()
}

func factorial(n int) uint64 {
	var v uint64
	var ok bool
	if v, ok = factorials[n]; !ok {
		v = factorial(n-1) * uint64(n)
		factorials[n] = v
	}
	return v
}

func factorialWithoutLower(n int, without int) uint64 {
	if n == without {
		return uint64(1)
	}
	var v = uint64(n)
	for i := n - 1; i > without; i-- {
		v *= uint64(i)
	}
	return v
}

func distinctPermutations(vertCount int, horizCount int) uint64 {
	if horizCount < vertCount {
		vertCount, horizCount = horizCount, vertCount
	}

	//fmt.Printf("Distinct permutations for %d,%d\n", vertCount, horizCount)
	var m map[int]uint64
	var ok bool
	var perms uint64
	if m, ok = distinctPerm[vertCount]; !ok {
		m = make(map[int]uint64)
		distinctPerm[vertCount] = m
	} else if perms, ok = m[horizCount]; ok {
		//	fmt.Printf("Precalculated permutations for %d+%d %d = %d!\n", horizCount, vertCount, horizCount+vertCount, perms)
		return perms
	}

	//fmt.Printf("Getting factorial for %d+%d %d!\n", horizCount, vertCount, horizCount+vertCount)

	totalPermutations := factorialWithoutLower(horizCount+vertCount, horizCount)

	//fmt.Printf("Got factorial = %d\n", totalPermutations)
	var divisor = uint64(1)
	//fmt.Printf("Divisor = %d\n", divisor)
	// if horizCount > 1 {
	// 	divisor *= factorial(horizCount)
	// 	fmt.Printf(" Divisor . hc %d = %d\n", factorial(horizCount), divisor)
	// }
	if vertCount > 1 {
		divisor *= factorial(vertCount)
		//	fmt.Printf(" Divisor * vc %d = %d\n", factorial(vertCount), divisor)
	}
	//divisor := factorial(horizCount) * factorial(vertCount)

	perms = totalPermutations / divisor

	m[horizCount] = perms

	return perms
}

func combinations(n int, log bool) uint64 {
	if n == 0 {
		return 0
	}
	// All n>0 can have a row of verticals, so start at one
	count := uint64(1)
	if n < 4 {
		return count
	} else if n == 4 {
		return 2
	}

	maxHorizCount := n / 4
	vertCount := n % 4

	if maxHorizCount == 0 {
		return 1
	}

	for horizCount := maxHorizCount; horizCount > 0; horizCount-- {
		p := distinctPermutations(n-horizCount*4, horizCount)
		count += p
		if log {
			fmt.Printf("v=%d, h=%d, p=%d\n", vertCount, horizCount, p)
		}
		vertCount += 4
	}

	return count
}

func runTest(n int, log bool) uint64 {
	if log {
		fmt.Printf("n=%d\n", n)
	}
	factorials = make(map[int]uint64)
	factorials[1] = 1
	//	factorials[0] = 1
	distinctPerm = make(map[int]map[int]uint64)
	c := combinations(n, log)

	if log {
		fmt.Printf("For n=%d, combinations=%d\n", n, c)
	}

	if c >= uint64(len(primeCount)) {
		return 0
	}
	answer := primeCount[c]
	//	return answer
	return answer
}

func main() {
	initialisePrimes()
	filenames := []string{"t1.txt", "t2.txt"}
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
		for ti := 0; ti < testCount; ti++ {
			n := nextInt(scanner)
			answer := runTest(n, opened)

			fmt.Printf("%d\n", answer)
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
