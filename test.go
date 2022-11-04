package main

import (
	"fmt"
)

func avrg(xs []float64) float64 {
	total := 0.0
	for _, accum := range xs {
		total += accum
	}
	return total / float64(len(xs))
}

func main() {
	test := make(map[string]int)

//	fmt.Println(avrg(xs))

	test["token"] = 20
	fmt.Println(test["token"])
}

