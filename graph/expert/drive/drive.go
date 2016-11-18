package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/drive

type busStation struct {
	index int

	earliestLeaveTime int
	distanceFromLast  int

	peopleStarting       []*person
	peopleEnding         []*person
	peoplePassingThrough []*person
}

func (bs *busStation) personStartsHere(p *person) {
	fmt.Printf("Person %d starts at station %d\n", p.index, bs.index)
	if p.startTime > bs.earliestLeaveTime {
		// Can't leave station until all passengers have arrived at station
		bs.earliestLeaveTime = p.startTime
	}
	bs.peopleStarting = append(bs.peopleStarting, p)
}

func (bs *busStation) personEndsHere(p *person) {
	fmt.Printf("Person %d ends at station %d\n", p.index, bs.index)
	bs.peopleEnding = append(bs.peopleEnding, p)
}

func (bs *busStation) personPassesThroughHere(p *person) {
	bs.peoplePassingThrough = append(bs.peoplePassingThrough, p)

}

func (bs *busStation) basicCost() int {
	var cost int
	for _, p := range bs.peopleStarting {
		var timeDiff = bs.earliestLeaveTime - p.startTime
		cost += timeDiff
	}
	return cost
}

type person struct {
	index             int
	startTime         int
	startStationIndex int
	destStationIndex  int
}

var stations []*busStation

func initStations(scanner *bufio.Scanner, totalStations int) {
	stations = make([]*busStation, 0, totalStations+1)
	blankStation := new(busStation)
	stations = append(stations, blankStation)
	for stationIndex := 0; stationIndex < totalStations; stationIndex++ {
		newStation := new(busStation)
		newStation.index = stationIndex
		newStation.peopleStarting = make([]*person, 0, 0)
		if stationIndex > 0 {
			distance := nextNumber(scanner)
			newStation.distanceFromLast = distance
		}
		stations = append(stations, newStation)
	}
}

func initPeople(scanner *bufio.Scanner, totalPeople int) {
	for personIndex := 0; personIndex < totalPeople; personIndex++ {
		var newPerson person
		newPerson.index = personIndex
		newPerson.startTime = nextNumber(scanner)
		newPerson.startStationIndex = nextNumber(scanner)
		newPerson.destStationIndex = nextNumber(scanner)

		fmt.Printf("Person %d starttime %d, station %d->%d\n", newPerson.index, newPerson.startTime, newPerson.startStationIndex, newPerson.destStationIndex)

		startStation := stations[newPerson.startStationIndex]
		destStation := stations[newPerson.destStationIndex]

		startStation.personStartsHere(&newPerson)
		destStation.personEndsHere(&newPerson)

		for stationIndex := newPerson.startStationIndex + 1; stationIndex < newPerson.destStationIndex; stationIndex++ {
			journeyStation := stations[stationIndex]
			journeyStation.personPassesThroughHere(&newPerson)
		}
	}
}

func runTest(scanner *bufio.Scanner, log bool) {
	totalStations := nextNumber(scanner)
	totalPeople := nextNumber(scanner)
	totalNitro := nextNumber(scanner)

	if log {
		fmt.Printf("Total stations:%d\nTotal people: %d\nTotal nitro: %d\n", totalStations, totalPeople, totalNitro)
	}

	initStations(scanner, totalStations)
	if log {
		fmt.Printf("Stations initialised: %d\n", len(stations))
	}
	initPeople(scanner, totalPeople)
	if log {
		fmt.Printf("People initialised\n")
	}
	for stationIndex, s := range stations {
		for _, p := range s.peopleStarting {
			fmt.Printf("%03d: Person arrives at station %d\n", p.startTime, stationIndex)
		}
	}
}

func main() {
	var filenames = []string{"drive_t1.txt"}

	for fileIndex, filename := range filenames {
		f, scanner, opened := openFile(filename)

		if opened {
			defer f.Close()
			if fileIndex > 0 {
				fmt.Println()
			}
			fmt.Printf("File: %s\n", filename)
		}

		runTest(scanner, opened)
	}
}

func openFile(filename string) (*os.File, *bufio.Scanner, bool) {
	var opened = true
	f, err := os.Open(filename)

	if err != nil {
		opened = false
		f = os.Stdin
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	return f, scanner, opened
}

func nextNumber(scanner *bufio.Scanner) int {
	scanner.Scan()
	word := scanner.Text()
	num, _ := strconv.ParseInt(word, 10, 32)

	return int(num)
}
