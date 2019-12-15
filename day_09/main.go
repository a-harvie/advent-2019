package main

import (
	"fmt"
	"os"

	"github.com/a-harvie/advent-2019/input"
	"github.com/a-harvie/advent-2019/intcode"
)

func main() {
	testDay09pt1 := []([]([]int)){
		{
			[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			[]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
		},
		{
			[]int{104, 1125899906842624, 99},
			[]int{1125899906842624},
		},
	}

	fmt.Println("Test 01")
	fmt.Printf("Expected: %v\n", testDay09pt1[0][0])
	_, output, _ := intcode.Compute(testDay09pt1[0][0], []int{}, make([]int, 0))
	fmt.Printf("Actual:   %v\n", output)

	fmt.Println("Test 02")
	fmt.Println("Expected: ################")
	_, output, _ = intcode.Compute(testDay09pt1[1][0], []int{}, make([]int, 0))
	fmt.Printf("Actual:   %v\n", output[0])

	fmt.Println("Test 03")
	fmt.Printf("Expected: %v\n", testDay09pt1[2][1])
	_, output, _ = intcode.Compute(testDay09pt1[2][0], []int{}, make([]int, 0))
	fmt.Printf("Actual:   %v\n", output)

	input, err := input.CommaSeparatedInts("input/input.txt")
	if err != nil {
		fmt.Printf("Oh noes: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(input[0])

	_, output, _ = intcode.Compute(input[0], []int{1}, []int{})
	fmt.Printf("Answer Day 09 pt 1: %v\n", output)
	_, output, _ = intcode.Compute(input[0], []int{2}, []int{})
	fmt.Printf("Answer Day 09 pt 2: %v\n", output)
}
