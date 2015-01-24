package main

import (
	"os"
)

func IfDirectory(dirname string) bool {
        f, err := os.Open(dirname)
        if err != nil {
                //fmt.Println(err) //erro check
                return false
        }
        defer f.Close()
        fi, err := f.Stat()
        if err != nil {
                //fmt.Println(err) //err check
                return false
        }
        switch mode := fi.Mode(); {
        case mode.IsDir():
                //fmt.Printf("it's a dir\n") //some shows
                return true
        case mode.IsRegular():
                //fmt.Printf("It's a file\n") //some shows part 2
                return false
        }
        return true
}
