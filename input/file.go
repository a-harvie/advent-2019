package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// CommaSeparatedInts takes path to a file with lines of comma separated ints,
// and returns a slice of int slices, one per line in the input file
func CommaSeparatedInts(filePath string) ([]([]int), error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	input := make([]([]int), 0)
	for {
		fileLine, err := reader.ReadString('\n')
		inputLine := make([]int, 0)
		for _, item := range strings.Split(fileLine, ",") {
			i, err := strconv.Atoi(strings.Trim(item, "\"\n"))
			if err != nil {
				return nil, err
			}
			inputLine = append(inputLine, i)
		}

		input = append(input, inputLine)

		if err != nil {
			break
		}
	}

	return input, nil
}

// StringOfInts takes a path to a file with lines of non-separated ints
// and retuns a slice of int slices, one per line in the input file
func StringOfInts(filePath string) ([]([]int), error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	input := make([]([]int), 0)
	for {
		fileLine, err := reader.ReadString('\n')
		inputLine := make([]int, 0)
		for i := 0; i < len(fileLine); i++ {
			i, err := strconv.Atoi(strings.Trim(string(fileLine[i]), "\"\n"))
			if err != nil {
				return nil, err
			}
			inputLine = append(inputLine, i)
		}

		input = append(input, inputLine)

		if err != nil {
			break
		}
	}

	return input, nil
}

// StringOfStrings takes a path to a file with lines of non-separated strings
// and returns a slice of string slices, one per line in the input file
func StringOfStrings(filePath string) ([]([]string), error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	input := make([]([]string), 0)
	for {
		fileLine, err := reader.ReadString('\n')

		input = append(input, strings.Split(strings.TrimSpace(fileLine), ""))

		if err != nil {
			break
		}
	}

	return input, nil
}

// Vectors parses AoC vectors of the format <x=0, y=1, z=2>, one per line
func Vectors(filePath string) ([]([]int), error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	input := make([]([]int), 0)
	for {
		fileLine, err := reader.ReadString('\n')
		fileLine = strings.Trim(fileLine, "<>\n")
		parts := strings.Split(fileLine, ", ")
		inputLine := make([]int, 3)
		for i, element := range parts {
			splitElem := strings.Split(element, "=")
			value, _ := strconv.Atoi(splitElem[1])
			inputLine[i] = value
		}

		input = append(input, inputLine)

		if err != nil {
			break
		}
	}

	return input, nil
}
