package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printColors(colors)
}

func printColors(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex of color ", color, " is ", hex)
	}
}
