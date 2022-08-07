package main

import "fmt"

func main() {
	vals := makeRange(0, 10)

	for _, v := range vals {
		if v%2 == 0 {
			fmt.Println(v, " is ", "even")
		} else {
			fmt.Println(v, " is ", "odd")
		}
	}
}

func makeRange(start int, end int) []int {
	v := make([]int, end-start+1)

	for i := range v {
		v[i] = start + i
	}

	return v
}
