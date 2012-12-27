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

func main() {
	const World = "Anything is..."
	const AnythingIs = true
	fmt.Println(whatever(15, 25, 0))
	fmt.Println(enlysium, excentrium, World, AnythingIs)
	
	fmt.Println(nevoiedeInt(Mic))
	fmt.Println(nevoiedeFloat(Mic)) 	
	fmt.Println(nevoiedeFloat(MaiMare)) 
	
	fmt.Printf("And the light %g ... went down", World, AnythingIs)
}
