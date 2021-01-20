package main

import "fmt"

func powerSeries(a int) (square int, cube int) {
	square = a * a
	cube = square * a
	return
}

func main() {

	fmt.Println(powerSeries(3))
}
