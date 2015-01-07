package main

import (
	"fmt"
	"time"
	"os"
	"math/rand"
)

func thename(name chan string) {
    for i := 0; ; i++ {
        name <- "/home"
    }
}

func kreator(name chan string) bool {
	for {
		if _, err := os.Stat(<-name); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Does not work\n")
			return false
	   }
		}
		whenever := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * whenever)
	fmt.Printf("Works\n")
	}
	return true
	fmt.Print("I am alive, as always\n"
}

func main() {
	var name chan string = make(chan string)

	go thename(name)
	go kreator(name)

    var input string
    fmt.Scanln(&input)
}
