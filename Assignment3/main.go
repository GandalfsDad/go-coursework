package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need to provide a filepath")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("Too many filepaths provided")
		os.Exit(1)
	}

	l := os.Args[1]
	fmt.Println(l)
}
