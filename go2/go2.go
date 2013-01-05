package main

import (
	"fmt"
	"math"
)
	
func swap(x, y string) (string, string) {
	return y, x
}

func declared(q, w, o float64) (float64, float64, float64) {
	return w, q, o
}

func pow(q, w, o float64) float64 {
	return math.Pow(q,math.Pow(w,o))	
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	
	q, w, o := declared(math.Pow(pow(2,2,2),3), pow(2,2,2), 10)
	fmt.Println(q, o, w)
}
