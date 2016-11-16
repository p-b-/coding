package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// https://www.hackerrank.com/challenges/dijkstrashortreach

type edge struct {
	connectedToNode *node
	distance        int
}

type node struct {
	index int
	//edges []*edge
	edges map[int]*edge

	distance int
}

type queue struct {
	items []interface{}
}

type nodeQueue struct {
	items []*node
}

func (q *queue) Initialise() {
	q.items = make([]interface{}, 0, 1000000)
}

func (q *queue) Enqueue(a interface{}) {
	q.items = append(q.items, a)

}
func (q *queue) Dequeue() interface{} {
	var a = q.items[0]
	q.items = q.items[1:]
	return a
}

func (q *queue) Empty() bool {
	return len(q.items) == 0
}

func (q *nodeQueue) Initialise() {
	q.items = make([]*node, 0, 1000000)
}

func (q *nodeQueue) Enqueue(n *node) {
	q.items = append(q.items, n)

}
func (q *nodeQueue) Dequeue() *node {
	var n = q.items[0]
	q.items = q.items[1:]
	return n
}

func (q *nodeQueue) Empty() bool {
	return len(q.items) == 0
}

var nodes []*node

func logNodes() {
	count := len(nodes)
	var stopLogAt = count
	var startLogAt = count
	if count > 100 {
		stopLogAt = 20
		startLogAt = count - 20
	}
	for i, n := range nodes {
		if i == stopLogAt {
			fmt.Printf("...\n...\n...\n")
			continue
		}
		if i >= stopLogAt && i < startLogAt {
			continue
		}
		if len(n.edges) > 5 {
			fmt.Printf("Node %d is at %p\n %d edges [unlisted]\n", i, n, len(n.edges))
		} else {
			fmt.Printf("Node %d is at %p\n", i, n)
			for _, e := range n.edges {
				fmt.Printf(" Connected to %d (distance %d)\n", e.connectedToNode.index, e.distance)
			}
		}
	}
}

func logNodeDistances() {
	count := len(nodes)
	var stopLogAt = count
	var startLogAt = count
	if count > 100 {
		stopLogAt = 20
		startLogAt = count - 20
	}
	for i, n := range nodes {
		if i == stopLogAt {
			fmt.Printf("...\n...\n...\n")
			continue
		}
		if i >= stopLogAt && i < startLogAt {
			continue
		}
		fmt.Printf("Node:%d, distance:%d\n", i, n.distance)
	}
}

func runTest(startingNode int, log bool) {
	log = false
	if log {
		fmt.Printf("Starting node: %d\n", startingNode)
		logNodes()
	}

	currentNode := nodes[startingNode]
	currentNode.distance = 0

	enqueued := make(map[int]bool, len(nodes))
	var q nodeQueue
	q.Initialise()
	q.Enqueue(currentNode)
	enqueued[currentNode.index] = true

	var potentialNewDistance int

	for q.Empty() == false {
		currentNode = q.Dequeue() //.(*node)
		delete(enqueued, currentNode.index)
		for _, edge := range currentNode.edges {
			potentialNewDistance = currentNode.distance + edge.distance
			if /*edge.connectedToNode.distance == -1 ||*/ edge.connectedToNode.distance > potentialNewDistance {
				edge.connectedToNode.distance = potentialNewDistance
				if _, ok := enqueued[edge.connectedToNode.index]; !ok {
					q.Enqueue(edge.connectedToNode)
				}
			}
		}
	}
	// Last node used to ensure no space is printed at the end of the line
	lastNode := len(nodes) - 1
	if lastNode == startingNode {
		lastNode--
	}
	/*	count := len(nodes)
		var stopLogAt = count
		var startLogAt = count
		if log && count > 100 {
			stopLogAt = 20
			startLogAt = count - 20
		}*/

	for i, n := range nodes {
		if i != startingNode {
			/*		if log {
					if i == stopLogAt {
						fmt.Printf("...\n...\n...\n")
						continue
					}
					if i >= stopLogAt && i < startLogAt {
						continue
					}
				}*/
			if n.distance == 0x7fffffff {
				fmt.Printf("-1")
			} else {
				fmt.Printf("%d", n.distance)
			}
			if i < lastNode {
				fmt.Printf(" ")
			}
		}
	}
	fmt.Printf("\n")

	//	logNodeDistances()
}

func main() {
	filenames := []string{"T4.txt"}
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
		testCount := nextInt(scanner)
		for testIndex := 0; testIndex < testCount; testIndex++ {
			startingNode := readTestConfig(scanner)
			runTest(startingNode, opened)
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

func readTestConfig(scanner *bufio.Scanner) int {
	nodeCount := nextInt(scanner)
	edgeCount := nextInt(scanner)

	nodes = make([]*node, 0, nodeCount)
	for nodeIndex := 0; nodeIndex < nodeCount; nodeIndex++ {
		newNode := new(node)
		newNode.index = nodeIndex
		//newNode.edges = make([]*edge, 0)
		newNode.edges = make(map[int]*edge)
		newNode.distance = 0x7fffffff
		nodes = append(nodes, newNode)
	}
	//	edgeMap := make(map[int]map[int]int)

	for edgeIndex := 0; edgeIndex < edgeCount; edgeIndex++ {
		edgeIndex1 := nextInt(scanner) - 1
		edgeIndex2 := nextInt(scanner) - 1
		distance := nextInt(scanner)

		if edgeIndex1 > edgeIndex2 {
			edgeIndex1, edgeIndex2 = edgeIndex2, edgeIndex1
		}
		/*		specificEdgeMap, ok := edgeMap[edgeIndex1]

				if !ok {
					//		fmt.Printf("Creating edge map for index %d\n", edgeIndex1)

					specificEdgeMap = make(map[int]int)
					edgeMap[edgeIndex1] = specificEdgeMap
				} else if previousDistance, ok := specificEdgeMap[edgeIndex1]; ok {
					//	fmt.Printf("Already exists %d->%d %d\n", edgeIndex1, edgeIndex2, previousDistance)
					if distance < previousDistance {
						specificEdgeMap[edgeIndex1] = distance
					}
					// Already created
					continue
				}
				//	fmt.Printf("New %d->%d %d\n", edgeIndex1, edgeIndex2, distance)

				specificEdgeMap[edgeIndex2] = distance*/

		//		previousDistance, exists := specificEdgeMap[edgeIndex1]
		//	fmt.Printf("SEM[%d] = %d,%t\n", edgeIndex1, previousDistance, exists)

		//		fmt.Printf("E%d -> E%d (%d) new\n", edgeIndex1, edgeIndex2, distance)
		originatingNode := nodes[edgeIndex1]
		destNode := nodes[edgeIndex2]
		var edge1 edge
		edge1.connectedToNode = destNode
		edge1.distance = distance
		if existingEdge, ok := originatingNode.edges[edgeIndex2]; ok {
			if existingEdge.distance > distance {
				existingEdge.distance = distance
				destNode.edges[edgeIndex1].distance = distance
			}
			continue
		}
		originatingNode.edges[edgeIndex2] = &edge1
		//originatingNode.edges = append(originatingNode.edges, &edge1)
		var edge2 edge
		edge2.connectedToNode = originatingNode
		edge2.distance = distance
		//destNode.edges = append(destNode.edges, &edge2)
		destNode.edges[edgeIndex1] = &edge2
	}

	startingNode := nextInt(scanner) - 1

	return startingNode
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
