package main

import (
        "fmt"
        "runtime"
	"go4/go4func.go"
)

func main() {
        vv := oSYS("")
        fmt.Printf("%s", vv)

        qq := suma(3,2)
        fmt.Printf("\n%d\n", qq)
	
	z := adauga()
	for i := 0; i<10; i++ {
	fmt.Printf("%d\n", z(i,i+1))
	}
}
