package main

import "fmt"

func main() {
	tests := []([]([]int)){
		{
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
		},
		{
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
		},
		{
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
		},
		{
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, test := range tests {
		fmt.Printf("Test: \n%v\nExpects: \n%v\n", test[0], test[1])
		program, err := intcodeCompute(test[0])
		if err != nil {
			fmt.Printf("Oh noes: \n%v\n", err)
			return
		}
		fmt.Printf("Actual: \n%v\n\n", program)
	}
}

func intcodeComputeInput(start []int, noun int, verb int) ([]int, error) {
	start[1] = noun
	start[2] = verb
	return intcodeCompute(start)
}

func intcodeCompute(start []int) ([]int, error) {

	program := make([]int, len(start))
	_ = copy(program, start)

	i := 0

	for {
		cmd := program[i]
		if cmd == 99 {
			break
		}

		indexA := program[i+1]
		indexB := program[i+2]
		indexO := program[i+3]
		opA := program[indexA]
		opB := program[indexB]
		var result int

		switch cmd {
		case 1:
			result = opA + opB
		case 2:
			result = opA * opB
		default:
			return nil, fmt.Errorf("oh noes, can't do %v", cmd)
		}

		program[indexO] = result

		i += 4
	}

	return program, nil
}
