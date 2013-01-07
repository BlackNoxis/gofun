package main

import (
        "fmt"
        // "math" // soon to be used
	// "strings" // soon to be used
)

type Vertex struct {
        X int
        Y string
}

type Structure struct {
        Q int
        W int
}

func putiere (q int) int {
        z := Structure{2, 3}
	b := &z
        return b.Q - b.W
}

var (
	um = Vertex{1, "Like this"}
	im = &Structure{1, 7}
	km = Structure{Q: 7}
	lm = Structure{}
)

type HartaCJ struct {
	Lat, Long float64
}

type HartaBN struct {
	Lat, Long float64
}

type HartaN struct {
	Lat, Long float64
}

var coordCJ map[string]HartaCJ
var coordBN map[string]HartaBN
var coordN map[string]HartaN

var hBN = map[string]HartaBN{
	"Bistrita": 	{47.0759, -24.2859},
	"Nasaud":	{47.1659, -24.2400},
}

var hN = map[string]HartaN{
	"Nasaud":	{47.1659, -24.2400},
}

var liberiBytes = []byte("+-.,/0123456789=_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:")

func main() {
        v := Vertex{1, "whatever"}
        v.X = 4
	em := new(Structure)
        fmt.Println(v.X, v.X, v.Y, putiere(5), um, im, km, lm, em)

	coordCJ = make(map[string]HartaCJ) // coordCJ([]HartaCJ)
	coordCJ["Cluj-Napoca"] = HartaCJ{
		46.46, - 23.36,
	}
	fmt.Println(coordCJ["Cluj-Napoca"])
	fmt.Println(hBN)
	
	nm := make(map[string]float64)
	
	nm["whatever"] = 50
	fmt.Println("This is the value:", nm["whatever"])
	
	nm["whatever"] = 60
	fmt.Println("This is the value:", nm["whatever"])
	
	e, ok := nm["whatever"]
	fmt.Println("this is the variabile:", e, "but it is", ok)	
	
	delete(nm, "whatever")
	fmt.Println("This is the value:", nm["Whatever"])
	
	re, rok := nm["whatever"]
	fmt.Println("this is the variabile:", re, "but it is", rok)

	if re == 0 {
		fmt.Println(hBN)
	} else {
		fmt.Println(coordCJ["Cluj-Napoca"])
	}
	
	if e != 0 {
		fmt.Println(coordCJ["Cluj-Napoca"])
	} else {

		fmt.Println(hBN)
		// this is wrong > fmt.Println(strings.EqualFold(hBN, hN))		
	}
	
	const F = "Bistrita-Nasaud is not Bistrita!"
	
	aq := []string{"Cluj-Napoca\n", "Bistrita\n", "Nasaud"}	
	fmt.Println(aq)
	for zq := 0; zq < len(aq); zq++ {
		fmt.Printf("Location %d is %s",
			zq, aq[zq])
				if zq == 2 && aq[2] == "Nasaud" {
					fmt.Printf("\nSo that's why %s is different from %s", aq[2], aq[1])
					fmt.Println("Because", hBN, "is something else than", hN)
				}		
		if aq[1] == "Bistrita\n" && aq[2] == "Nasaud" {
			if zq == 1 {
			fmt.Printf("If Location %d is not %s, then %s\n",zq, aq[2], F)
			}
		}	
	}
	fmt.Println(hNN)
}
