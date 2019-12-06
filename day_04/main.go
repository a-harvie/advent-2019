package main

import (
	"fmt"
	"strconv"
)

func main() {
	tests := make(map[string]bool)
	tests["111111"] = true
	tests["223450"] = false
	tests["123789"] = false

	for k, v := range tests {
		fmt.Printf("Test: \n%v\nExpects: \n%v\nActual: \n%v\n", k, v, isValid(k, false))
	}

	low := 367479
	high := 893698
	fmt.Printf("Final pt 1: %v\n", numValid(low, high, false))
	fmt.Printf("Final pt 2: %v\n", numValid(low, high, true))
}

func numValid(low int, high int, strict bool) int {
	valid := 0

	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)

		if isValid(s, strict) {
			valid++
		}
	}
	return valid
}

func isValid(password string, strict bool) bool {
	pw := convertNumeric(password)
	// fmt.Println(pw)
	dnd := doesNotDecrease(pw)
	hd := hasDouble(pw, strict)
	return dnd && hd
}

func convertNumeric(s string) []int {
	n := make([]int, 0)

	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(string(s[i]))
		n = append(n, num)
	}

	return n
}

func doesNotDecrease(seq []int) bool {
	for i := 0; i < len(seq)-1; i++ {
		if seq[i] > seq[i+1] {
			return false
		}
	}
	return true
}

func hasDouble(seq []int, strict bool) bool {
	if len(seq) < 2 {
		return false
	}

	for i := 1; i < len(seq); i++ {
		if strict {
			if i == 1 {
				if seq[i] == seq[i-1] && seq[i] != seq[i+1] {
					return true
				}
			} else if i == len(seq)-1 {
				if seq[i] == seq[i-1] && seq[i-1] != seq[i-2] {
					return true
				}
			} else {
				if seq[i] == seq[i-1] && seq[i-1] != seq[i-2] && seq[i] != seq[i+1] {
					return true
				}
			}
		} else {
			if seq[i] == seq[i-1] {
				return true
			}
		}
	}

	return false
}
