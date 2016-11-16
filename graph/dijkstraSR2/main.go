package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/dijkstrashortreach

type edge struct {
	connectedToNode *node
	distance        int
}

type node struct {
	index int
	edges []*edge

	distance int
}

type queue struct {
	items []interface{}
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

var nodes []*node

func logNodes() {
	for i, n := range nodes {
		fmt.Printf("Node %d is at %p\n", i, n)
		for _, e := range n.edges {
			fmt.Printf(" Connected to %d (distance %d)\n", e.connectedToNode.index, e.connectedToNode.distance)
		}
	}
}

func logNodeDistances() {
	for i, n := range nodes {
		fmt.Printf("Node:%d, distance:%d\n", i, n.distance)
	}
}

func runTest(startingNode int, log bool) {
	if log {
		fmt.Printf("Starting node: %d\n", startingNode)
		logNodes()
	}

	currentNode := nodes[startingNode]
	currentNode.distance = 0

	var q queue
	q.Enqueue(currentNode)

	for q.Empty() == false {
		currentNode = q.Dequeue().(*node)
		for _, edge := range currentNode.edges {
			if edge.connectedToNode.distance == -1 || edge.connectedToNode.distance > currentNode.distance+edge.distance {
				edge.connectedToNode.distance = currentNode.distance + edge.distance
				q.Enqueue(edge.connectedToNode)
			}
		}
	}
	// Last node used to ensure no space is printed at the end of the line
	lastNode := len(nodes) - 1
	if lastNode == startingNode {
		lastNode--
	}
	for i, n := range nodes {
		if i != startingNode {
			fmt.Printf("%d", n.distance)
			if i < lastNode {
				fmt.Printf(" ")
			}
		}
	}
	fmt.Printf("\n")

	//	logNodeDistances()
}

func main() {
	filenames := []string{"T1.txt"}
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
			startingNode := readTestConfig(scanner)
			runTest(startingNode, opened)
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
		newNode.edges = make([]*edge, 0)
		newNode.distance = -1
		nodes = append(nodes, newNode)
	}
	edgeMap := make(map[int]map[int]bool)

	for edgeIndex := 0; edgeIndex < edgeCount; edgeIndex++ {
		edgeIndex1 := nextInt(scanner) - 1
		edgeIndex2 := nextInt(scanner) - 1
		distance := nextInt(scanner)

		if edgeIndex1 > edgeIndex2 {
			edgeIndex1, edgeIndex2 = edgeIndex2, edgeIndex1
		}
		specificEdgeMap, ok := edgeMap[edgeIndex1]
		if !ok {
			specificEdgeMap = make(map[int]bool)
			edgeMap[edgeIndex1] = specificEdgeMap
		} else if _, ok := specificEdgeMap[edgeIndex1]; ok {
			// Already created
			continue
		}
		originatingNode := nodes[edgeIndex1]
		destNode := nodes[edgeIndex2]
		var edge1 edge
		edge1.connectedToNode = destNode
		edge1.distance = distance
		originatingNode.edges = append(originatingNode.edges, &edge1)
		var edge2 edge
		edge2.connectedToNode = originatingNode
		edge2.distance = distance
		destNode.edges = append(destNode.edges, &edge2)
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
