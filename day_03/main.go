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

func main() {
	tests := []([]([]string)){
		{
			[]string{"R8", "U5", "L5", "D3"},
			[]string{"U7", "R6", "D4", "L4"},
			[]string{"6"},
		},
		{
			[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			[]string{"159"},
		},
		{
			[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			[]string{"135"},
		},
	}

	for _, test := range tests {
		fmt.Printf("Test: \n%v\n %v\nExpects: \n%v\nActual: \n%v\n", test[0], test[1], test[2], closestIntersection(test[0], test[1]))
	}

	input, _ := getInput()
	// fmt.Printf("input: %v\n%v", len(input), input)
	fmt.Printf("Final answer, part 1: %v\n", closestIntersection(input[0], input[1]))

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

func closestIntersection(wire1 []string, wire2 []string) int {
	path1 := calculatePath(wire1)
	path2 := calculatePath(wire2)
	intersections := getIntersections(path1, path2)
	return getMinDistance(intersections)
}

func calculatePath(wire []string) []point {
	path := make([]point, 0)
	x := 0
	y := 0
	path = append(path, point{0, 0})

	for _, step := range wire {
		direction := step[0]
		distance, _ := strconv.Atoi(step[1:])

		switch direction {
		case 'U':
			limit := y + distance
			for y = y + 1; y < limit; y++ {
				path = append(path, point{x, y})
			}
		case 'D':
			limit := y - distance
			for y = y - 1; y > limit; y-- {
				path = append(path, point{x, y})
			}
		case 'L':
			limit := x - distance
			for x = x - 1; x > limit; x-- {
				path = append(path, point{x, y})
			}
		case 'R':
			limit := x + distance
			for x = x + 1; x < limit; x++ {
				path = append(path, point{x, y})
			}
		}

	}
	// fmt.Printf("getPath: %v\n", path)
	return path
}

func getIntersections(path1 []point, path2 []point) []point {
	intersections := make([]point, 0)
	for _, p := range path1 {
		if p.x == 0 && p.y == 0 {
			continue
		}
		for _, p2 := range path2 {
			if p.x == 0 && p.y == 0 {
				continue
			}
			if p.x == p2.x && p.y == p2.y {
				intersections = append(intersections, point{p.x, p.y})
			}
		}
	}
	// fmt.Printf("getIntersections: %v\n", intersections)
	return intersections
}

func getMinDistance(path []point) int {
	min := math.MaxInt64
	for _, p := range path {
		d := manhattanDistance(p, point{0, 0})
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
