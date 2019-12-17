package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/a-harvie/advent-2019/input"
)

type asteroidMap []([]string)

type coord struct {
	x int
	y int
}

func main() {
	tests, coordExpectations, countExpectations := getTests()
	for i, test := range tests {
		// if i == 0 {
		// 	continue
		// }
		// if i > 0 {
		// 	break
		// }
		fmt.Printf("Test %v expects: %v %v\n", i, coordExpectations[i], countExpectations[i])
		list := parseMap(test)
		bestLocation, count := findBestLocation(list)
		fmt.Printf("Actual: %v %v\n", bestLocation, count)
		// the last test is the only one we had an example for pt 2
		if i == 4 {
			vaporized := vaporize(bestLocation, list, 200)
			fmt.Printf("Vaporized: %v\n", vaporized)
		}
	}
	input, _ := input.StringOfStrings("input/input.txt")
	asteroidList := parseMap(input)
	best, count := findBestLocation(asteroidList)
	fmt.Println(best, count)
	fmt.Printf("Answer pt 1: %v\n", count)

	vapo := vaporize(best, asteroidList, 200)
	fmt.Println(vapo)
	fmt.Printf("Answer pt 2: %v\n", (vapo.x*100 + vapo.y))
}

func parseMap(m asteroidMap) []coord {
	asteroidList := make([]coord, 0)
	for y := 0; y < len(m); y++ {
		fmt.Println(m[y])
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] == "#" {
				asteroidList = append(asteroidList, coord{x, y})
			}
		}
	}
	return asteroidList
}

func findBestLocation(aList []coord) (coord, int) {
	mostVisible := 0
	bestLocation := coord{0, 0}

	for _, a1 := range aList {
		visibleAsteroids := make(map[float64]bool)
		visibleCount := 0
		for _, a2 := range aList {
			if a1 != a2 {
				slope := getSlope(a1, a2)
				if visibleAsteroids[slope] == false {
					visibleAsteroids[slope] = true
					visibleCount++
				}
			}
		}
		if visibleCount > mostVisible {
			mostVisible = visibleCount
			bestLocation = a1
		}
	}

	return bestLocation, mostVisible
}

type firingOrder map[float64][]target

type target struct {
	pos  coord
	dist float64
}

func vaporize(start coord, aList []coord, limit int) coord {
	// map of firing angles to a slice of targets at that angle. Will be sorted later
	fo := getFiringOrder(start, aList)

	// can't order map keys, so get a sorted slice of the keys
	keys := make([]float64, len(fo))
	i := 0
	for k := range fo {
		keys[i] = k
		i++

	}
	sort.Float64s(keys)

	var lastTarget coord
	count := 0
	// iterate through the map by sorted keys (clockwise)
	// remove the asteroids as we go
	for {
		if count >= limit || len(fo) == 0 {
			fmt.Printf("Breaking: %v %v\n", count, len(fo))
			break
		}
		nextKeys := make([]float64, 0)
		for _, slope := range keys {
			lastTarget = fo[slope][0].pos
			count++
			if len(fo[slope]) == 1 {
				// last asteroid at this angle. Remove angle and don't persist its key.
				delete(fo, slope)
			} else {
				fo[slope] = fo[slope][1:]
				nextKeys = append(nextKeys, slope)
			}
			if count >= limit {
				break
			}
		}
		keys = nextKeys

	}

	return lastTarget
}

func getFiringOrder(start coord, aList []coord) firingOrder {
	fo := make(map[float64][]target)
	for _, targ := range aList {
		if start == targ {
			continue
		}
		slope := getSlope(start, targ)
		dist := getDistance(start, targ)
		// fmt.Println(targ, slope, dist)
		fo[slope] = append(fo[slope], target{targ, dist})
	}

	// sort asteroids at the same angle by their distance
	for k, v := range fo {
		sort.Slice(v, func(a, b int) bool {
			return v[a].dist < v[b].dist
		})
		fo[k] = v
	}
	return fo
}

func getSlope(a1 coord, a2 coord) float64 {
	x := float64(a1.x - a2.x)
	y := float64(a1.y - a2.y)

	// Atan2 already starts with "up" = 0 but will go "counterclockwise"
	// Use the inversion to go "clockwise"
	angle := (-1 * math.Atan2(x, y)) * 180.0 / math.Pi
	if angle < 0 {
		angle += 360.0
	}
	if angle == -0 {
		angle = 0
	}

	return angle
}

func getDistance(a1 coord, a2 coord) float64 {
	dx := float64(a2.x - a1.x)
	dy := float64(a2.y - a1.y)
	return math.Sqrt(dx*dx + dy*dy)
}

// this version doesn't work. Problem with gcd?
// func getSlope(a1 coord, a2 coord) coord {
// 	dy := a2.y - a1.y
// 	dx := a2.x - a1.x
// 	if dx == 0 {
// 		// return math.Inf(dy / abs(dy))
// 		return coord{0, dy / abs(dy)}
// 	} else if dy == 0 {
// 		// return math.Copysign(0, float64(dx/abs(dx)))
// 		return coord{dx / abs(dx), 0}
// 	}

// 	x, y := reduce(dx, dy)
// 	return coord{x, y}
// }

// func reduce(x int, y int) (int, int) {
// 	gcd := gcd(x, y)
// 	return x / gcd, y / gcd
// }

// func gcd(x int, y int) int {
// 	for y != 0 {
// 		x, y = y, x%y
// 	}
// 	return x
// }

// func abs(x int) int {
// 	if x >= 0 {
// 		return x
// 	}
// 	return -x
// }
