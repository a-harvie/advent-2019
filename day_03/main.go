package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type pathStep struct {
	pos   point
	delay int
}

func main() {
	tests := []([]([]string)){
		{
			[]string{"R8", "U5", "L5", "D3"},
			[]string{"U7", "R6", "D4", "L4"},
			[]string{"6"},
			[]string{"30"},
		},
		{
			[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			[]string{"159"},
			[]string{"610"},
		},
		{
			[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			[]string{"135"},
			[]string{"410"},
		},
	}

	for _, test := range tests {
		closest, leastDelay := closestIntersection(test[0], test[1])
		fmt.Printf("Test: \n%v\n %v\nExpects: \n%v\nActual: \n%v\n", test[0], test[1], test[2], closest)
		fmt.Printf("Pt 2 expects: \n%v\nActual: \n%v\n", test[3], leastDelay)
	}

	input, _ := getInput()
	// fmt.Printf("input: %v\n%v", len(input), input)
	closest, leastDelay := closestIntersection(input[0], input[1])
	fmt.Printf("Final answer, part 1: %v\n", closest)
	fmt.Printf("Final asnwer, part 2: %v\n", leastDelay)

}

func getInput() ([]([]string), error) {
	file, err := os.Open("input/input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	lines := make([]([]string), 0)
	for {
		fileLine, err := reader.ReadString('\n')

		line := make([]string, 0)
		for _, item := range strings.Split(fileLine, ",") {
			line = append(line, strings.Trim(item, "\"\n"))
		}
		lines = append(lines, line)

		if err != nil {
			break
		}
	}

	return lines, nil
}

func closestIntersection(wire1 []string, wire2 []string) (int, int) {
	path1 := calculatePath(wire1)
	path2 := calculatePath(wire2)
	intersections := getIntersections(path1, path2)
	return getMinDistance(intersections), getMinDelay(intersections)
}

func calculatePath(wire []string) []pathStep {
	path := make([]pathStep, 0)
	x := 0
	y := 0
	delay := 0
	path = append(path, pathStep{point{x, y}, delay})

	for _, step := range wire {
		direction := step[0]
		distance, _ := strconv.Atoi(step[1:])

		switch direction {
		case 'U':
			limit := y + distance
			for yy := y + 1; yy <= limit; yy++ {
				y = yy
				pos := point{x, y}
				// delay = getDelay(delay, path, pos)
				delay++
				path = append(path, pathStep{pos, delay})
			}
		case 'D':
			limit := y - distance
			for yy := y - 1; yy >= limit; yy-- {
				y = yy
				pos := point{x, y}
				// delay = getDelay(delay, path, pos)
				delay++
				path = append(path, pathStep{pos, delay})
			}
		case 'L':
			limit := x - distance
			for xx := x - 1; xx >= limit; xx-- {
				x = xx
				pos := point{x, y}
				// delay = getDelay(delay, path, pos)
				delay++
				path = append(path, pathStep{pos, delay})
			}
		case 'R':
			limit := x + distance
			for xx := x + 1; xx <= limit; xx++ {
				x = xx
				pos := point{x, y}
				// delay = getDelay(delay, path, pos)
				delay++
				path = append(path, pathStep{pos, delay})
			}
		}

	}
	// fmt.Printf("getPath: %v\n", path)
	return path
}

// func getDelay(delay int, path []pathStep, newPos point) int {
// 	for _, s := range path {
// 		if s.pos.x == newPos.x && s.pos.y == newPos.y {
// 			return s.delay
// 		}
// 	}

// 	return delay + 1
// }

func getIntersections(path1 []pathStep, path2 []pathStep) []pathStep {
	intersections := make([]pathStep, 0)
	for _, s := range path1 {
		if s.pos.x == 0 && s.pos.y == 0 {
			continue
		}
		for _, s2 := range path2 {
			if s.pos.x == 0 && s.pos.y == 0 {
				continue
			}
			if s.pos.x == s2.pos.x && s.pos.y == s2.pos.y {
				intersections = append(intersections, pathStep{point{s.pos.x, s.pos.y}, s.delay + s2.delay})
			}
		}
	}
	// fmt.Printf("getIntersections: %v\n", intersections)
	return intersections
}

func getMinDistance(path []pathStep) int {
	min := math.MaxInt64
	for _, s := range path {
		d := manhattanDistance(s.pos, point{0, 0})
		if d < min {
			min = d
		}
	}
	return min
}

func manhattanDistance(point1 point, point2 point) int {
	return abs(point1.x-point2.x) + abs(point1.y-point2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMinDelay(path []pathStep) int {
	min := math.MaxInt64
	for _, s := range path {
		if s.delay < min {
			min = s.delay
		}
	}

	return min
}
