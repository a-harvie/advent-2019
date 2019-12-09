package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type param struct {
	value int
	mode  int
}

func main() {
	// tests := []([]([]int)){
	// 	{
	// 		[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
	// 		[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
	// 	},
	// 	{
	// 		[]int{1, 0, 0, 0, 99},
	// 		[]int{2, 0, 0, 0, 99},
	// 	},
	// 	{
	// 		[]int{2, 3, 0, 3, 99},
	// 		[]int{2, 3, 0, 6, 99},
	// 	},
	// 	{
	// 		[]int{2, 4, 4, 5, 99, 0},
	// 		[]int{2, 4, 4, 5, 99, 9801},
	// 	},
	// 	{
	// 		[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
	// 		[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
	// 	},
	// }

	// testDay05pt1 := []([]([]int)){
	// 	{
	// 		[]int{3, 0, 4, 0, 99},
	// 		[]int{1},
	// 	},
	// }

	testDay05pt2 := []([]([]int)){
		{
			[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			[]int{-1, 0, 1, 2},
			[]int{1, 0, 1, 1},
		},
		{
			[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			[]int{-1, 0, 1, 2},
			[]int{1, 0, 1, 1},
		},
		// {
		// 	[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		// 		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		// 		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
		// 	[]int{1, 7, 8, 9, 10000},
		// 	[]int{999, 999, 1000, 1001, 1001},
		// },
	}

	// for _, test := range tests {
	// 	output := make([]int, 0)
	// 	fmt.Printf("Test: \n%v\nExpects: \n%v\n", test[0], test[1])
	// 	program, output, err := intcodeCompute(test[0], 0, output)
	// 	if err != nil {
	// 		fmt.Printf("Oh noes: \n%v\n", err)
	// 		return
	// 	}
	// 	fmt.Printf("Actual: \n%v\n", program)
	// }
	// for _, test := range testDay05pt1 {
	// 	output := make([]int, 0)
	// 	fmt.Printf("Test: \n%v\nExpects: \n%v\n", test[0], test[1])
	// 	_, output, err := intcodeCompute(test[0], 0, output)
	// 	if err != nil {
	// 		fmt.Printf("Oh noes: \n%v\n", err)
	// 		return
	// 	}
	// 	fmt.Printf("Output: \n%v\n\n", output)
	// }

	for _, test := range testDay05pt2 {
		fmt.Printf("Program: %v\n", test[0])
		for i, input := range test[1] {
			output := make([]int, 0)
			testOutput := test[2][i]
			fmt.Printf("Test input: \n%v\nExpects: \n%v\n", input, testOutput)
			_, output, err := intcodeCompute(test[0], input, output)
			if err != nil {
				fmt.Printf("Oh noes: \n%v\n", err)
				return
			}
			fmt.Printf("Output: \n%v\n\n", output)
		}

	}

	input, err := getInput()
	if err != nil {
		fmt.Printf("Oh noes: %v\n", err)
		os.Exit(1)
	}
	output := make([]int, 0)
	_, output, err = intcodeCompute(input, 1, output)
	fmt.Printf("Final answer, pt 1: %v\n", output)
	output = make([]int, 0)
	_, output, err = intcodeCompute(input, 5, output)
	fmt.Printf("Final answer, pt 1: %v\n", output)
}

func getInput() ([]int, error) {
	file, err := os.Open("input/input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	input := make([]int, 0)
	for {
		fileLine, err := reader.ReadString('\n')
		for _, item := range strings.Split(fileLine, ",") {
			i, err := strconv.Atoi(strings.Trim(item, "\"\n"))
			if err != nil {
				return nil, err
			}
			input = append(input, i)
		}

		if err != nil {
			break
		}
	}

	return input, nil
}

func intcodeCompute(start []int, input int, output []int) ([]int, []int, error) {

	program := make([]int, len(start))
	_ = copy(program, start)

	i := 0

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
			params = append(params, input)
			opInput(&program, params)
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
	fmt.Printf("Input called, inputting: %v\n", params[1])
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
