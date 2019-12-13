package input

import (
	"bufio"
	"os"
	"strings"
	"strconv"
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

		input = append(input,inputLine)

		if err != nil {
			break
		}
	}

	return input, nil
}