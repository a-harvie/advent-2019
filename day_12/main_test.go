package main

import (
	"fmt"
	"testing"

	"github.com/a-harvie/advent-2019/input"
)

func BenchmarkSimulated(b *testing.B) {
	input, err := input.Vectors("input/input.txt")
	if err != nil {
		fmt.Printf("Oh noes: %v\n", err)
	}
	moons := makeMoons(input)

	for n := 0; n < b.N; n++ {
		moons.simulate()
	}
}
