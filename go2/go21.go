package main

import "fmt"

func min(x []int) int {
    		men := x[0]
    		for i := 1; i < len(x); i++ {
    			if x[i] < men {
    				men = x[i]
    			}
    		}
    		return men
}

func main() {
	p := []int{322, 44, 55, 66, 77, 88, 312}
	fmt.Println("minimum p ==", min(p))
}
