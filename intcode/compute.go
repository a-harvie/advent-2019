package intcode

import (
	"fmt"
)

type param struct {
	value int
	mode  int
}

// Compute takes an input program as a slice of integers, and an integer input value,
// and computes to the provided output slice
func Compute(start []int, input []int, output []int) ([]int, []int, error) {

	program := make([]int, len(start))
	_ = copy(program, start)

	i := 0
	inputPointer := 0

	for {
		cmd, paramModes := parseCmd(program[i])
		// fmt.Printf("Processing: %v %v", cmd, paramModes)
		if cmd == 99 {
			break
		}

		params := parseParams(&program, i, cmd, paramModes)

		// fmt.Printf(" parsed params: %v\n", params)

		switch cmd {
		case 1:
			opAdd(&program, params)
			i += 4
		case 2:
			opMult(&program, params)
			i += 4
		case 3:
			params = append(params, input[inputPointer])
			opInput(&program, params)
			inputPointer++
			i += 2
		case 4:
			output = append(output, opOutput(&program, params))
			i += 2
		case 5:
			i = opJumpTrue(&program, i, params)
		case 6:
			i = opJumpFalse(&program, i, params)
		case 7:
			opLessThan(&program, params)
			i += 4
		case 8:
			opEquals(&program, params)
			i += 4
		default:
			return nil, nil, fmt.Errorf("oh noes, can't do %v", cmd)
		}
	}

	return program, output, nil
}

func parseCmd(input int) (int, []int) {
	paramModes := make([]int, 3)
	if input < 10 {
		return input, paramModes
	}

	cmd := input % 100
	params := input / 100
	for i := 0; i < len(paramModes); i++ {
		paramModes[i] = params % 10
		params /= 10
	}

	// this should work for arbitrarily long param lists
	// counter := 0
	// for params := input / 100; params < 0; params /= 10 {
	// 	paramModes[counter] = params % 10
	// 	counter++
	// }

	return cmd, paramModes
}

func parseParams(program *[]int, index int, cmd int, paramModes []int) []int {
	var params []int

	switch cmd {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 7:
		fallthrough
	case 8:
		params = make([]int, 3)
		if paramModes[0] == 0 {
			params[0] = (*program)[(*program)[index+1]]
		} else {
			params[0] = (*program)[index+1]
		}

		if paramModes[1] == 0 {
			params[1] = (*program)[(*program)[index+2]]
		} else {
			params[1] = (*program)[index+2]
		}

		params[2] = (*program)[index+3]
	case 3:
		fallthrough
	case 4:
		params = make([]int, 1)
		params[0] = (*program)[index+1]
	case 5:
		fallthrough
	case 6:
		params = make([]int, 2)
		if paramModes[0] == 0 {
			params[0] = (*program)[(*program)[index+1]]
		} else {
			params[0] = (*program)[index+1]
		}

		if paramModes[1] == 0 {
			params[1] = (*program)[(*program)[index+2]]
		} else {
			params[1] = (*program)[index+2]
		}
	}

	return params
}

func opAdd(program *[]int, params []int) {
	(*program)[params[2]] = params[0] + params[1]
}

func opMult(program *[]int, params []int) {
	(*program)[params[2]] = params[0] * params[1]
}

func opInput(program *[]int, params []int) {
	// fmt.Printf("Input called, inputting: %v\n", params[1])
	(*program)[params[0]] = params[1]
}

func opOutput(program *[]int, params []int) int {
	return (*program)[params[0]]
}

func opJumpTrue(program *[]int, index int, params []int) int {
	if params[0] != 0 {
		return params[1]
	}

	return index + 3
}

func opJumpFalse(program *[]int, index int, params []int) int {
	if params[0] == 0 {
		return params[1]
	}

	return index + 3
}

func opLessThan(program *[]int, params []int) {
	output := 0
	if params[0] < params[1] {
		output = 1
	}
	(*program)[params[2]] = output
}

func opEquals(program *[]int, params []int) {
	output := 0
	if params[0] == params[1] {
		output = 1
	}
	(*program)[params[2]] = output
}
