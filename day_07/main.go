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
			[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			[]int{65210},
			[]int{1, 0, 4, 3, 2},
		},
	}

	testDay07pt2 := []([]([]int)){
		{
			[]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
			[]int{139629729},
			[]int{9, 8, 7, 6, 5},
		},
		{
			[]int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
			[]int{18216},
			[]int{9, 7, 8, 5, 6},
		},
	}

	fmt.Println("TESTS PT 1")
	for _, test := range testDay07pt1 {
		fmt.Printf("Program: %v\n", test[0])
		fmt.Printf("Expects: %v %v\n", test[1][0], test[2])
		output, phases, err := iterateAmplifiers(test[0])
		if err != nil {
			fmt.Printf("Oh noes: \n%v\n", err)
			return
		}
		fmt.Printf("Actual:  %v %v\n", output, phases)
	}
	fmt.Println("TESTS PT 2")
	for _, test := range testDay07pt2 {
		fmt.Printf("Program: %v\n", test[0])
		fmt.Printf("Expects: %v %v\n", test[1][0], test[2])
		output, phases, err := concurrentAmplifiers(test[0])
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

	output, phases, err := iterateAmplifiers(input[0])
	if err != nil {
		fmt.Printf("Oh noes: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Final answer, pt 1: %v %v\n", output, phases)
	output, phases, err = concurrentAmplifiers(input[0])
	fmt.Printf("Final answer, pt 2: %v %v\n", output, phases)

}

func iterateAmplifiers(program []int) (int, []int, error) {
	phases := []int{0, 1, 2, 3, 4}
	maxOutput := math.MinInt64
	maxPhases := []int{0, 1, 2, 3, 4}
	for permutation := make([]int, len(phases)); permutation[0] < len(permutation); nextPermutation(permutation) {
		phasePermutation := getPermutation(phases, permutation)
		outputs := make([]int, 0)
		programOutput := 0
		for _, phaseInput := range phasePermutation {
			programInput := 0
			if len(outputs) > 0 {
				programInput = outputs[len(outputs)-1]
			}
			_, output, err := intcode.Compute(program, []int{phaseInput, programInput}, make([]int, 0))
			if err != nil {
				return 0, nil, err
			}
			programOutput = output[0]
			outputs = append(outputs, programOutput)
		}
		if programOutput > maxOutput {
			maxOutput = programOutput
			maxPhases = phasePermutation
		}
	}
	return maxOutput, maxPhases, nil
}

func concurrentAmplifiers(program []int) (int, []int, error) {
	phases := []int{5, 6, 7, 8, 9}
	maxOutput := math.MinInt64
	maxPhases := []int{5, 6, 7, 8, 9}
	for permutation := make([]int, len(phases)); permutation[0] < len(permutation); nextPermutation(permutation) {
		phasePermutation := getPermutation(phases, permutation)
		inputChannels := make([]chan int, len(phases))
		outputChannels := make([]chan int, len(phases))
		errorChannels := make([]chan error, len(phases))
		running := make([]bool, len(phases))

		programOutput := 0

		// Set up a computation for each amp, giving it the phase input
		for i, phaseInput := range phasePermutation {
			inputChannels[i] = make(chan int)
			outputChannels[i] = make(chan int)
			errorChannels[i] = make(chan error)

			go intcode.ChannelCompute(program, inputChannels[i], outputChannels[i], errorChannels[i])
			running[i] = true
			inputChannels[i] <- phaseInput
		}
		// Initial input to amp A
		inputChannels[0] <- 0

		done := make(chan bool)
		go func() {
			for {
				// Pass each amp's output to the next amp's input
				select {
				case x, open := <-outputChannels[0]:
					if open {
						inputChannels[1] <- x
					} else {
						running[0] = false
					}
				case x, open := <-outputChannels[1]:
					if open {
						inputChannels[2] <- x
					} else {
						running[1] = false
					}
				case x, open := <-outputChannels[2]:
					if open {
						inputChannels[3] <- x
					} else {
						running[2] = false
					}
				case x, open := <-outputChannels[3]:
					if open {
						inputChannels[4] <- x
					} else {
						running[3] = false
					}
				case x, open := <-outputChannels[4]:
					if open {
						// Check amp A is still running before trying to pass back to it
						if running[0] {
							inputChannels[0] <- x
						}
						programOutput = x
					} else {
						running[4] = false
					}
				}

				stillRunning := false
				for _, r := range running {
					stillRunning = stillRunning || r
				}
				if !stillRunning {
					done <- true
					return
				}
			}
		}()
		<-done
		if programOutput > maxOutput {
			maxOutput = programOutput
			maxPhases = phasePermutation
		}

	}
	return maxOutput, maxPhases, nil
}

// shamelessly stolen from https://stackoverflow.com/a/30230552
func nextPermutation(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

// shamelessly stolen from https://stackoverflow.com/a/30230552
func getPermutation(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}
