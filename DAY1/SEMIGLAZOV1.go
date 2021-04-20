package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SolutionSquareGenerator(4, 10))
	fmt.Println(len(SolutionSquareGenerator(4, 10)))
}

func SolutionSquareGenerator(start int, n int) []int {
	var naturalNumberSlice []int
	naturalNumber := math.Abs(float64(start))
	for i := 0; i < n; i++ {
		naturalNumberSlice = append(naturalNumberSlice, int(math.Pow(naturalNumber, 2)))
		naturalNumber += 1
	}
	return naturalNumberSlice
}
