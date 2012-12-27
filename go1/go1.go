package main

import "fmt"

var enlysium bool = true
var excentrium int = 777

func whatever(sum, scad, inmu int) (a, b, c, d, e, f, z int) {
	a = sum + 25
	b = sum + a
	c = sum + b
	
	d = sum - 25
	e = sum - 30
	f = sum - 50
	
	z = inmu + a + b + 3
	return
}

const (
	Mic = 4<<7
	MaiMare = Mic>>5
)

func nevoiedeInt(g int) int { return g*1 + 1 }
func nevoiedeFloat(g float64) float64 {
	return g*2
}

func anotherSum (sumq int) (sume int) {
	sumq = sumq
	sume = sumq + 3
	for u := 0; u < 1; u++ {
		sume += u
	}
	return
}


func main() {
	sumz := 0
	for i := 0; i < 10; i++ {
		sumz += i
	}	
	
	const World = "Anything is..."
	const AnythingIs = true
	fmt.Println(whatever(15, 25, 0))
	fmt.Println(enlysium, excentrium, World, AnythingIs)
	
	fmt.Println(nevoiedeInt(Mic))
	fmt.Println(nevoiedeFloat(Mic)) 	
	fmt.Println(nevoiedeFloat(MaiMare)) 
	
	fmt.Printf("And the light %g ... went down", World, AnythingIs) 

	fmt.Println(" ", sumz)

		
	fmt.Println("The sum for all first 3 numbers is: ", anotherSum(3))
}
