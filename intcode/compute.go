package intcode

import (
	"fmt"
)

const (
	memoryMultiplier int = 100 // "much" larger ?!?
)

type param struct {
	value int
	mode  int
}

// Compute takes an input program as a slice of integers, and an integer input value,
// and computes to the provided output slice
func Compute(start []int, input []int, output []int) ([]int, []int, error) {

	program := make([]int, len(start)*memoryMultiplier)
	_ = copy(program, start)

	i := 0
	r := 0 // relative base
	inputPointer := 0

	for {
		cmd, paramModes := parseCmd(program[i])
		// fmt.Printf("Processing: %v %v %v", cmd, r, paramModes)
		if cmd == 99 {
			break
		}

		params := parseParams(&program, i, r, cmd, paramModes)

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
			output = append(output, params[0])
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
		case 9:
			r += opAdjustBase(&program, params)
			i += 2
		default:
			return nil, nil, fmt.Errorf("oh noes, can't do %v", cmd)
		}
	}
	//
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

func parseParams(program *[]int, index int, relativeBase int, cmd int, paramModes []int) []int {
	var params []int
	// fmt.Println("parseParams", index, cmd)
	switch cmd {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 7:
		fallthrough
	case 8:
		params = make([]int, 3)
		params[0] = parseParam(program, index+1, relativeBase, paramModes[0], false)
		params[1] = parseParam(program, index+2, relativeBase, paramModes[1], false)
		params[2] = parseParam(program, index+3, relativeBase, paramModes[2], true)
	case 3:
		params = make([]int, 1)
		params[0] = parseParam(program, index+1, relativeBase, paramModes[0], true)
	case 4:
		fallthrough
	case 9:
		params = make([]int, 1)
		params[0] = parseParam(program, index+1, relativeBase, paramModes[0], false)
	case 5:
		fallthrough
	case 6:
		params = make([]int, 2)
		params[0] = parseParam(program, index+1, relativeBase, paramModes[0], false)
		params[1] = parseParam(program, index+2, relativeBase, paramModes[1], false)
	}

	return params
}

func parseParam(program *[]int, index int, relativeBase int, paramMode int, writeToMemory bool) int {
	var param int
	switch paramMode {
	case 0:
		if writeToMemory {
			param = (*program)[index]
		} else {
			param = (*program)[(*program)[index]]
		}
	case 1:
		param = (*program)[index]
	case 2:
		if writeToMemory {
			param = (*program)[index] + relativeBase
		} else {
			param = (*program)[(*program)[index]+relativeBase]
		}
	}
	return param
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

func opAdjustBase(program *[]int, params []int) int {
	return params[0]
}
