package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ratings []int
var candies []int

func ratingString(rating int) string {
	if rating > 5 {
		return fmt.Sprintf("*****%d", rating)
	}
	return strings.Repeat("*", rating) + fmt.Sprintf("%d", rating)
}

func sumCandies(log bool) int {
	total := 0
	for _, c := range candies {
		total += c
	}
	return total
}

func increaseCandiesBackwards(ci int, log bool) {
	for ; ci >= 1; ci-- {
		if ratings[ci-1] > ratings[ci] &&
			candies[ci-1] == candies[ci] {
			candies[ci-1]++
		} else {
			return
		}
	}
}

func countSmallerRatings(childCount int, ci int, log bool) int {
	var smallest = ratings[ci]
	if log {
		fmt.Printf("Counting smaller ratings from %d\n", ci)
	}
	var count = 1

	for ci++; ci < childCount; ci++ {
		if ratings[ci] < smallest {
			if log {
				fmt.Printf("Rating [%d]=%d is smaller than %d\n", ci, ratings[ci], smallest)
			}
			count++
			smallest = ratings[ci]
		} else {
			if log {
				fmt.Printf("Rating [%d]=%d is >= than %d\n", ci, ratings[ci], smallest)
				fmt.Printf("Count %d\n", count)
			}
			return count
		}
	}
	return count
}

func runTest2(childCount int, log bool) int {
	if childCount == 0 {
		return 0
	} else if childCount == 1 {
		return 1
	}
	candies[0] = 1
	for ci := 1; ci < childCount; ci++ {
		var rating = ratings[ci-1]
		var nextRating = ratings[ci]
		if nextRating > rating {
			candies[ci] = candies[ci-1] + 1
		} else if nextRating == rating {
			candies[ci] = 1
		} else if nextRating < rating {
			smallerCount := countSmallerRatings(childCount, ci-1, log)
			if candies[ci-1] < smallerCount {
				candies[ci-1] = smallerCount
			}
			smallerCount--
			for smallerCount > 0 {
				if log {
					fmt.Printf("Setting [%d]=%d\n", ci, smallerCount)
				}
				candies[ci] = smallerCount
				smallerCount--
				ci++
				if log && childCount < 500 {
					for logIndex := 0; logIndex < childCount; logIndex++ {
						fmt.Printf("[%d] %s %d ", logIndex, ratingString(ratings[logIndex]), candies[logIndex])
					}
					fmt.Println()
				}
			}
			// Next loop will come back to this point
			ci--
		}
		if log && childCount < 500 {
			for logIndex := 0; logIndex < childCount; logIndex++ {
				fmt.Printf("[%d] %s %d ", logIndex, ratingString(ratings[logIndex]), candies[logIndex])
			}
			fmt.Println()
		}
	}
	return sumCandies(log)
}

func runTest(childCount int, log bool) int {
	if childCount == 0 {
		return 0
	} else if childCount == 1 {
		return 1
	}
	candies[0] = 1
	for ci := 1; ci < childCount; ci++ {
		var rating = ratings[ci-1]
		var nextRating = ratings[ci]
		if nextRating > rating {
			candies[ci] = candies[ci-1] + 1
		} else if nextRating == rating {
			candies[ci] = 1
		} else if nextRating < rating {
			candies[ci] = 1
			if candies[ci-1] == 1 {
				// Previous (higher) rating had 1 candy, need to increase that to two
				if log {
					fmt.Printf("Need to go backwards from %d an increase candies\n", ci)
				}
				increaseCandiesBackwards(ci, log)
			}
		}
		if log && childCount < 500 {
			for logIndex := 0; logIndex < childCount; logIndex++ {
				fmt.Printf("[%d] %s %d ", logIndex, ratingString(ratings[logIndex]), candies[logIndex])
			}
			fmt.Println()
		}
	}
	return sumCandies(log)
}

func main() {
	filenames := []string{"t1.txt", "t2.txt", "t3.txt"} //, "t4.txt"}
	for fi, fn := range filenames {
		f, opened, scanner := openFile(fn)

		if opened {
			defer f.Close()
			if fi > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", fn)
		}

		childCount := readConfig(scanner, opened)

		var answer int
		if fi == 2 {
			answer = runTest2(childCount, opened)
		} else {
			answer = runTest2(childCount, false)
		}

		fmt.Printf("%d\n", answer)

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

func readConfig(scanner *bufio.Scanner, log bool) int {
	childCount := nextInt(scanner)
	ratings = make([]int, 0, childCount)
	candies = make([]int, childCount, childCount)

	for ci := 0; ci < childCount; ci++ {
		r := nextInt(scanner)
		ratings = append(ratings, r)
	}

	return childCount
}
