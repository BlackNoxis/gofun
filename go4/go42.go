package main

import (
	"fmt"
)

func prod(q float64) float64 {
	return q*q	
}

func supra(q, e float64) float64 {
	x := prod(q)
	diff := x - e
	q = diff / 2.0*x
	fmt.Println(x,e,diff,q)
	return q
}

//func Sqrt(x float64) float64 {
//	fmt.Println(q)
//}

func main() {
	fmt.Println(supra(2,3))
}
