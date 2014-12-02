package main

// Inspired by http://youtu.be/DpwUVExX27E

import (
	"fmt"
)

const (
	MaxDepth = 6 // The code is exponential, so if this is too high, expect to wait for years before your answer comes back
)

func step(numerator int, denominator int, currentDepth uint, maxDepth uint) {
	if currentDepth >= maxDepth {
		return
	}
	value := float64(numerator) / float64(denominator)
	fmt.Printf("%d/%d = %f\n", numerator, denominator, value)
	sum := numerator + denominator
	newDepth := currentDepth + 1
	step(numerator, sum, newDepth, maxDepth)
	step(sum, denominator, newDepth, maxDepth)
}

func main() {
	step(1, 1, 0, MaxDepth)
}
