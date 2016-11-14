package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func runTest(scanner *bufio.Scanner) {
	var runesToOutput []rune
	str := readString(scanner)

	//var buffer bytes.Buffer

	count := 0
	counting := false
	var runeToCount rune
	for i, w := 0, 0; i < len(str); i += w {

		r, width := utf8.DecodeRuneInString(str[i:])
		w = width

		if counting {
			if runeToCount == r {
				count++
			} else {
				count++
				if count%2 == 1 {
					runesToOutput = append(runesToOutput, runeToCount)
					//					buffer.WriteRune(runeToCount)
				} else {
					// Consecutive runes completely removed, the current rune being consider may
					//  be consecutive to rune before this consecutive runesToOutput
					// aaabbbbbac -> abbbac -> aac -> c
					//                ^^^  removing these means the 'a's are now consecutive
					if lastRuneIs(runesToOutput, r) {
						runesToOutput = removeLastRune(runesToOutput)
						counting = false
					}
				}
				count = 0
				runeToCount = r
			}
		} else {
			counting = true
			count = 0
			runeToCount = r
			if lastRuneIs(runesToOutput, r) {
				runesToOutput = removeLastRune(runesToOutput)
				count = 1
			}
		}
	}
	if counting {
		count++
		if count%2 == 1 {
			runesToOutput = append(runesToOutput, runeToCount)
			//			buffer.WriteRune(runeToCount)
		} else {
			if lastRuneIs(runesToOutput, runeToCount) {
				runesToOutput = removeLastRune(runesToOutput)
			}
		}
	}

	outputString := string(runesToOutput)

	//outputString := buffer.String()
	if len(outputString) == 0 {
		fmt.Printf("Empty String\n")
	} else {

		fmt.Printf("%s\n", outputString)
	}
}

func removeLastRune(runes []rune) []rune {
	if len(runes) == 0 {
		return runes
	}
	return runes[:len(runes)-1]
}

func lastRuneIs(runes []rune, compareTo rune) bool {
	if len(runes) == 0 {
		return false
	}
	if runes[len(runes)-1] == compareTo {
		return true
	}
	return false
}

func main() {
	var filenames = [...]string{"testCases/srs_1.txt",
		"testCases/srs_2.txt",
		"testCases/srs_3.txt"}
	for fileIndex, filename := range filenames {
		if fileIndex > 0 {
			fmt.Println()
		}
		f, scanner, success := openFile(filename)

		if success {
			fmt.Printf("Opened file: %s\n", filename)
			defer f.Close()
		}

		runTest(scanner)

		if !success {
			// reading from stdin - do not loop
			break
		}
	}
}

func openFile(filename string) (*os.File, *bufio.Scanner, bool) {
	openedFile := true
	f, err := os.Open(filename)

	if err != nil {
		f = os.Stdin
		openedFile = false
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	return f, scanner, openedFile
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
