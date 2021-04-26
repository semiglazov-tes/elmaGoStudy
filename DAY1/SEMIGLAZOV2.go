package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	maxZeroCount := 0
	binPresArr := SolutionBinaryGap(4)
	for _, value := range binPresArr {
		if strings.Contains(value, "0") {
			if strings.Count(value, "0") > maxZeroCount {
				maxZeroCount = strings.Count(value, "0")
			}
		}
	}
	fmt.Println(maxZeroCount)
}
func SolutionBinaryGap(N int) []string {
	binPresentation := strconv.FormatInt(int64(N), 2)
	binPresentationSplit := strings.Split(binPresentation, "1")
	return binPresentationSplit
}
