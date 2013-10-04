package main

import (
    "fmt"
    "os"
    //"path/filepath"
    "path"
)

func main() {
    vv := "/whatever"
    if path.IsAbs(vv) == true {
    	fmt.Print("It's not there, joe\n")
	}

    fmt.Printf("Name a directory which you want to scan for files: \n")
    var i string
    fmt.Scanf("%s",&i)
    d, err := os.Open(i)
    if err != nil {
        fmt.Println(err)
	fmt.Printf("Do you want to create it?\n")
	var w string
	fmt.Scanf("%s",&w)
	if w == "Yes" || w == "YES" {
		fmt.Println(os.Mkdir(i, 22))
	}
        os.Exit(1)
    }
    fi, err := d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, fi := range fi {
        if fi.Mode().IsRegular() {
            fmt.Println(fi.Name(), fi.Size(), "bytes")
        }
    }
}
