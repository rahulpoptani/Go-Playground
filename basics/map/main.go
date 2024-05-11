package main

import "fmt"

func main() {
	// Add element while map creation
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	// add new element in map
	colors["white"] = "#fffff"
	fmt.Println(colors)
	// delete element from map
	delete(colors, "red")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Color is", color, "and hex is", hex)
	}
}
