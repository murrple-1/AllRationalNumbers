package main

// Inspired by http://youtu.be/DpwUVExX27E

import (
	"container/list"
	"fmt"
)

const (
	MaxDepth = 6 // The code is exponential, so if this is too high, expect to wait for years before your answer comes back
)

func generate_depthFirst(numerator uint, denominator uint, currentDepth uint, maxDepth uint) {
	if currentDepth >= maxDepth {
		return
	}
	fmt.Printf("%d / %d = %g\n", numerator, denominator, float64(numerator)/float64(denominator))
	sum := numerator + denominator
	newDepth := currentDepth + 1
	generate_depthFirst(numerator, sum, newDepth, maxDepth)
	generate_depthFirst(sum, denominator, newDepth, maxDepth)
}

type Vertex struct {
	numerator   uint
	denominator uint
}

const (
	MaxVertexes = 40
)

func generate_breadthFirst(maxVertexes uint) {
	queue := list.New()
	queue.PushBack(&Vertex{1, 1})
	var step uint = 1
	for e := queue.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Vertex)
		sum := v.numerator + v.denominator
		if step < maxVertexes {
			queue.PushBack(&Vertex{v.numerator, sum})
			step++
		}
		if step < maxVertexes {
			queue.PushBack(&Vertex{sum, v.denominator})
			step++
		}
		fmt.Printf("%d / %d = %g\n", v.numerator, v.denominator, float64(v.numerator)/float64(v.denominator))
	}
}

func main() {
	fmt.Println("Depth First Generation")
	generate_depthFirst(1, 1, 0, MaxDepth)
	fmt.Println()
	fmt.Println("Breadth First Generation")
	generate_breadthFirst(MaxVertexes)
}
