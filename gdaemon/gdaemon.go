/* LICENSE: AGPL v3 
Copyright Stefan Cristian B. < stefan.cristian @ rogentos.ro > */

package main

import (
	"fmt"
	"time"
	"os"
	"math/rand"
)


func kreator(name string) bool {
	for {
		if _, err := os.Stat(name); err != nil {
		 if os.IsNotExist(err) {
			fmt.Printf("Does not work\n")
			return false
		 }
		}
		whenever := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * whenever)
		fmt.Printf("Works. It is okay\n")
	}
	return true
}

func main() {
	var name string
	fmt.Printf("gib the name: \n")
	fmt.Scanf("%s", &name)

	if IfFile(name) == true {
		go kreator(name)
	} else { os.Exit(1) }

	var input string
	fmt.Scanln(&input)
}
