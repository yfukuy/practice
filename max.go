package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x := 3
	y := 5

	fmt.Printf("x = %d,y = %d: max = %d\n", x, y, max(x, y))
}
