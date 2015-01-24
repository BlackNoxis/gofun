package main

import (
	"fmt"
	"os"
)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "whatever" {
			fmt.Printf("You have accessed the secret function. Tread carefully.\n")
		} else {
			fmt.Printf("Error: You cannot add additional options or path names\n")
			os.Exit(1)
		}
	}
}
func main() {
	go omega()
}

