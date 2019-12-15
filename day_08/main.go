package main

import (
	"fmt"
	"math"

	"github.com/a-harvie/advent-2019/input"
)

const width int = 25
const height int = 6

type image []([]([]int))
type layer []([]int)
type row []int

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}
	fmt.Println(makeImage(test, 3, 2))
	fmt.Println(layerProduct(fewestZeros(makeImage(test, 3, 2))))

	input, err := input.StringOfInts("input/input.txt")
	if err != nil {
		fmt.Printf("Oh noes: %v", err)
	}

	img := makeImage(input[0], width, height)
	fmt.Printf("Answer pt 1: %v\n", layerProduct(fewestZeros(img)))
	fmt.Println("Answer pt 2:")
	renderImage(img)
}

func makeImage(pixelVals []int, w int, h int) image {
	image := make(image, len(pixelVals)/(w*h))

	count := 0
	for i := 0; i < len(image); i++ {
		layer := make(layer, h)
		for j := 0; j < h; j++ {
			row := make(row, w)
			for k := 0; k < w; k++ {
				row[k] = pixelVals[count]
				count++
			}
			layer[j] = row
		}
		image[i] = layer
	}

	return image
}

func fewestZeros(img image) layer {
	zeros := math.MaxInt64
	fewestIndex := 0
	for i, layer := range img {
		zeroCount := zeroCount(layer)
		if zeroCount < zeros {
			fewestIndex = i
			zeros = zeroCount
		}
	}

	return img[fewestIndex]
}

func zeroCount(l layer) int {
	count := 0
	for _, row := range l {
		for _, pixel := range row {
			if pixel == 0 {
				count++
			}
		}
	}

	return count
}

func layerProduct(l layer) int {
	ones := 0
	twos := 0
	for _, row := range l {
		for _, pixel := range row {
			switch pixel {
			case 1:
				ones++
			case 2:
				twos++
			default:
			}
		}
	}

	return ones * twos
}

func renderImage(img image) {
	w := len(img[0][0])
	h := len(img[0])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			for z := 0; z < len(img); z++ {
				if img[z][y][x] == 2 {
					continue
				} else {
					if img[z][y][x] == 1 {
						fmt.Printf("□")
					} else {
						fmt.Printf(" ")
					}
					break
				}
			}
		}
		fmt.Println()
	}
}
