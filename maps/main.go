package main

import "fmt"

func main() {

	//Alternative Creates
	//var colours map[string]string
	//colours := make(map[string]string)

	colours := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}

	//Add something
	colours["white"] = "#ffffff"

	//Remove Something
	delete(colours, "white")

	//fmt.Println(colours)
	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}
