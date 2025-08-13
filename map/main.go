package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red": "#FF0000",
	// 	"green": "#008000",
	// }

	// var colors map[string]string

	// colors := make(map[int]string) // built in function to make a map
	// colors[10] = "#ffffff" // add a new item to the map
	// delete(colors, 10) // remove an item from them map


	colors := map[string]string{
		"red": "#FF0000",
		"green": "#008000",
		"white": "#ffffff",
	}
 
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hexcode for ", color, "is", hex)
	}
}