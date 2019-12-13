package main

import (
	"fmt"
	"math"
	"os"

	"github.com/a-harvie/advent-2019/input"
	"github.com/a-harvie/advent-2019/intcode"
)

type param struct {
	value int
	mode  int
}

func main() {
	testDay07pt1 := []([]([]int)){
		{
			[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			[]int{43210},
			[]int{4, 3, 2, 1, 0},
		},
		{
			[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			[]int{54321},
			[]int{0, 1, 2, 3, 4},
		},
		{
			[]int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0},
			[]int{65210},
			[]int{1,0,4,3,2},
		},
	}

	for _, test := range testDay07pt1 {
		fmt.Printf("Program: %v\n", test[0])
		fmt.Printf("Expects: %v %v\n", test[1], test[2])
		output, phases, err := iterateAmplifiers(test[0])
		if err != nil {
			fmt.Printf("Oh noes: \n%v\n", err)
			return
		}
		fmt.Printf("Actual:  %v %v\n", output, phases)
	}

	input, err := input.CommaSeparatedInts("input/input.txt")
	if err != nil {
		fmt.Printf("Oh noes: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(input[0])
	// output := make([]int, 0)
	output, phases, err := iterateAmplifiers(input[0])
	if err != nil {
		fmt.Printf("Oh noes: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Final answer, pt 1: %v %v\n", output, phases)
	// output = make([]int, 0)
	// _, output, err = intcode.Compute(input[0], 5, output)
	// fmt.Printf("Final answer, pt 1: %v\n", output)
}

func iterateAmplifiers(program []int ) (int, []int, error) {
	phases := []int{0,1,2,3,4}
	maxOutput := math.MinInt64
	maxPhases := []int{0,1,2,3,4}
	for permutation := make([]int, len(phases)); permutation[0] < len(permutation); nextPermutation(permutation) {
		phasePermutation := getPermutation(phases, permutation)
		outputs := make([]int, 0)
		programOutput := 0
		for _, phaseInput := range phasePermutation {
			programInput := 0
			if len(outputs) > 0 {
				programInput = outputs[len(outputs)-1]
			}
			_, output, err := intcode.Compute(program, []int{phaseInput,programInput}, make([]int, 0))
			if err != nil {
				return 0, nil, err
			}
			programOutput = output[0]
			outputs = append(outputs,programOutput)
		}
		if programOutput > maxOutput {
			maxOutput = programOutput
			maxPhases = phasePermutation
		}
	}
	return maxOutput, maxPhases, nil
}

func nextPermutation(p []int) {
    for i := len(p) - 1; i >= 0; i-- {
        if i == 0 || p[i] < len(p)-i-1 {
            p[i]++
            return
        }
        p[i] = 0
    }
}

func getPermutation(orig, p []int) []int {
    result := append([]int{}, orig...)
    for i, v := range p {
        result[i], result[i+v] = result[i+v], result[i]
    }
    return result
}