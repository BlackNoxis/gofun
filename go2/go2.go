package main

import (
	"fmt"
	"math"
)

func mov(a,b,c float64) float64 {
	if q := math.Pow(a,b); q < c {
		return c
	} else {
		fmt.Printf("Numarul %g este mai mare ca %g\n", q, c)
	}
	return c		
}

func main() {
	fmt.Println(
		mov(3,3,20), "\n",
	)
}
