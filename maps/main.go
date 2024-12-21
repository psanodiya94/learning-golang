package main

import (
	"fmt"
	"image/color"
)

func printMap(c map[int]color.RGBA) {
	for key, value := range c {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Println("---------------------------")
}

func main() {
	// colors := map[int]color.RGBA{
	// 	1: color.RGBA{255, 0, 0, 255},
	// 	2: color.RGBA{0, 255, 0, 255},
	// 	3: color.RGBA{0, 0, 255, 255},
	// }

	// var colors map[int]color.RGBA

	colors := make(map[int]color.RGBA)

	colors[1] = color.RGBA{255, 0, 0, 255}
	colors[2] = color.RGBA{0, 255, 0, 255}
	colors[3] = color.RGBA{0, 0, 255, 255}

	printMap(colors)

	// Add a new color
	colors[4] = color.RGBA{255, 255, 0, 0}

	printMap(colors)

	// Remove a color
	delete(colors, 4)

	printMap(colors)
}
