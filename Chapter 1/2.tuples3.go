package main

import (
	"fmt"
)

func powerSeries(a int) (int, int, error) {
	square := a * a
	cube := square * a
	return square, cube, nil
}

func main() {
	fmt.Println(powerSeries(3))
}
