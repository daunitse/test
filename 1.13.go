package main

import (
	"fmt"
	"math"
)

func main() {
	var a, g int
	_, err := fmt.Scanln(&a)
	if err != nil {
		panic(err)
	}
	for i := 1; i < a; i++ {
		g = math.Pow(float64(i), 2.0)
	}
}
