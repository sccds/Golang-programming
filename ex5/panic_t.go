package main

import (
	"fmt"
	"math"
)

func ConvertIntToInt16(x int) int16 {
	if math.MinInt16 <= x && x <= math.MaxInt16 {
		return int16(x)
	}
	panic(fmt.Sprintf("%d is out of int16 range", x))
}

func main() {
	i := ConvertIntToInt16(655567)
	fmt.Printf("%d", i)
}
