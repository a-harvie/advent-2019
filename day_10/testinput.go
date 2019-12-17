package main

func getTests() ([]asteroidMap, []coord, []int) {
	return []asteroidMap{
			{
				[]string{".", "#", ".", ".", "#"},
				[]string{".", ".", ".", ".", "."},
				[]string{"#", "#", "#", "#", "#"},
				[]string{".", ".", ".", ".", "#"},
				[]string{".", ".", ".", "#", "#"},
			},
			{
				[]string{".", ".", ".", ".", ".", ".", "#", ".", "#", "."},
				[]string{"#", ".", ".", "#", ".", "#", ".", ".", ".", "."},
				[]string{".", ".", "#", "#", "#", "#", "#", "#", "#", "."},
				[]string{".", "#", ".", "#", ".", "#", "#", "#", ".", "."},
				[]string{".", "#", ".", ".", "#", ".", ".", ".", ".", "."},
				[]string{".", ".", "#", ".", ".", ".", ".", "#", ".", "#"},
				[]string{"#", ".", ".", "#", ".", ".", ".", ".", "#", "."},
				[]string{".", "#", "#", ".", "#", ".", ".", "#", "#", "#"},
				[]string{"#", "#", ".", ".", ".", "#", ".", ".", "#", "."},
				[]string{".", "#", ".", ".", ".", ".", "#", "#", "#", "#"},
			},
			{
				[]string{"#", ".", "#", ".", ".", ".", "#", ".", "#", "."},
				[]string{".", "#", "#", "#", ".", ".", ".", ".", "#", "."},
				[]string{".", "#", ".", ".", ".", ".", "#", ".", ".", "."},
				[]string{"#", "#", ".", "#", ".", "#", ".", "#", ".", "#"},
				[]string{".", ".", ".", ".", "#", ".", "#", ".", "#", "."},
				[]string{".", "#", "#", ".", ".", "#", "#", "#", ".", "#"},
				[]string{".", ".", "#", ".", ".", ".", "#", "#", ".", "."},
				[]string{".", ".", "#", "#", ".", ".", ".", ".", "#", "#"},
				[]string{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
				[]string{".", "#", "#", "#", "#", ".", "#", "#", "#", "."},
			},
			{
				[]string{".", "#", ".", ".", "#", ".", ".", "#", "#", "#"},
				[]string{"#", "#", "#", "#", ".", "#", "#", "#", ".", "#"},
				[]string{".", ".", ".", ".", "#", "#", "#", ".", "#", "."},
				[]string{".", ".", "#", "#", "#", ".", "#", "#", ".", "#"},
				[]string{"#", "#", ".", "#", "#", ".", "#", ".", "#", "."},
				[]string{".", ".", ".", ".", "#", "#", "#", ".", ".", "#"},
				[]string{".", ".", "#", ".", "#", ".", ".", "#", ".", "#"},
				[]string{"#", ".", ".", "#", ".", "#", ".", "#", "#", "#"},
				[]string{".", "#", "#", ".", ".", ".", "#", "#", ".", "#"},
				[]string{".", ".", ".", ".", ".", "#", ".", "#", ".", "."},
			},
			{
				[]string{".", "#", ".", ".", "#", "#", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", "#", "#", "#", "#"},
				[]string{"#", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", ".", ".", "#", "#", "."},
				[]string{".", "#", ".", "#", "#", "#", "#", "#", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", ".", "#"},
				[]string{".", "#", "#", "#", ".", "#", "#", "#", "#", "#", "#", "#", ".", "#", "#", "#", "#", ".", "#", "."},
				[]string{"#", "#", "#", "#", "#", ".", "#", "#", ".", "#", ".", "#", "#", ".", "#", "#", "#", ".", "#", "#"},
				[]string{".", ".", "#", "#", "#", "#", "#", ".", ".", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
				[]string{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
				[]string{"#", ".", "#", "#", "#", "#", ".", ".", ".", ".", "#", "#", "#", ".", "#", ".", "#", ".", "#", "#"},
				[]string{"#", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
				[]string{"#", "#", "#", "#", "#", ".", "#", "#", ".", "#", "#", "#", ".", ".", "#", "#", "#", "#", ".", "."},
				[]string{".", ".", "#", "#", "#", "#", "#", "#", ".", ".", "#", "#", ".", "#", "#", "#", "#", "#", "#", "#"},
				[]string{"#", "#", "#", "#", ".", "#", "#", ".", "#", "#", "#", "#", ".", ".", ".", "#", "#", ".", ".", "#"},
				[]string{".", "#", "#", "#", "#", "#", ".", ".", "#", ".", "#", "#", "#", "#", "#", "#", ".", "#", "#", "#"},
				[]string{"#", "#", ".", ".", ".", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", ".", ".", "."},
				[]string{"#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", ".", "#", "#", "#", "#", "#", "#", "#"},
				[]string{".", "#", "#", "#", "#", ".", "#", ".", "#", "#", "#", ".", "#", "#", "#", ".", "#", ".", "#", "#"},
				[]string{".", ".", ".", ".", "#", "#", ".", "#", "#", ".", "#", "#", "#", ".", ".", "#", "#", "#", "#", "#"},
				[]string{".", "#", ".", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", ".", "#", "#", "#"},
				[]string{"#", ".", "#", ".", "#", ".", "#", "#", "#", "#", "#", ".", "#", "#", "#", "#", ".", "#", "#", "#"},
				[]string{"#", "#", "#", ".", "#", "#", ".", "#", "#", "#", "#", ".", "#", "#", ".", "#", ".", ".", "#", "#"},
			},
		},
		[]coord{
			coord{3, 4},
			coord{5, 8},
			coord{1, 2},
			coord{6, 3},
			coord{11, 13},
		},
		[]int{
			8,
			33,
			35,
			41,
			210,
		}
}
