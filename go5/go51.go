package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func whatever(a string) string {
        cmd := exec.Command("tr", "a-z", "A-Z")
        cmd.Stdin = strings.NewReader("mother")
        var out bytes.Buffer
        cmd.Stdout = &out
        err := cmd.Run()
        if err != nil {
                log.Fatal(err)
        }
	a = out.String()
        //fmt.Printf("in all caps: %q\n", out.String())
	//fmt.Printf("in all caps: %q\n", a)
	return a
}

func main() {
	//cmd := exec.Command("tr", "a-z", "A-Z")
	//cmd.Stdin = strings.NewReader("mother")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("in all caps: %q\n", out.String())
	fmt.Println("The requested main function shows this:", whatever("any string here") )
}
