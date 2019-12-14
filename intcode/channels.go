package intcode

import (
	"fmt"
)

// ChannelCompute takes an input program as a slice of integers, and an integer input channel,
// and computes to the provided integer output channel.
func ChannelCompute(start []int, input chan int, output chan int, errors chan error) {
	program := make([]int, len(start))
	_ = copy(program, start)

	i := 0

	for {
		cmd, paramModes := parseCmd(program[i])
		// fmt.Printf("Processing: %v %v", cmd, paramModes)
		if cmd == 99 {
			close(input)
			close(output)
			close(errors)
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
			inputParam := <-input
			opInput(&program, append(params, inputParam))
			i += 2
		case 4:
			output <- opOutput(&program, params)
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
			errors <- fmt.Errorf("oh noes, can't do %v", cmd)
			// todo: keep going?
			close(input)
			close(output)
			close(errors)
			break
		}
	}
}
