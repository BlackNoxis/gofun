package main

import (
	"fmt"
)

func main() {
	var i float64
	var b float64
	fmt.Printf("State number first number: ")
	fmt.Scanf("%E", &i)
	fmt.Printf("Good. State the maximum number of which you want percentage of: ")
	fmt.Scanf("%E", &b)

	var x float64
	max := 100.00

	x = max / ( b / i )
	porcent := "%"

	fmt.Printf("here`s your percentage: %v%s \n", x, porcent)
}
