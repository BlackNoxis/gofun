package main

import (
        "fmt"
        //"strings"
        "os"
        //"errors"
)

func main() {

	if Exists("dude") {
		fmt.Printf("yay. it's there. something\n")
		//IfDirectory("dude")
		} else {
		var i string
		fmt.Printf("meh. you want to create it? Yes/No\n")
		fmt.Scanf("%s",&i)
		//i_yes := []string{"Y","yes","Yes","Ye","YES","YeS","Ye","y"}
		//i_no := []string{"No","N","NO","nO"}
		if i == "Yes" || i == "YES" || i == "Y" || i == "y" || i == "Ye" || i == "YE" || i == "yes" || i == "ye" {
                	fmt.Printf("Then name a directory name\n")
		        var dir string
			fmt.Scanf("%s",&dir)
			fmt.Println(os.Mkdir(dir, 22))
		} else {
			fmt.Printf("Have fun, then\n")
		}
	}
	IfDirectory("dude")
	//ww := os.Mkdir("dude", 22)

        //if fileinfo.IsDir(); err != nil {
        	//fmt.Printf("It's already created. Let's see if it's a file or a directory\n")
        //} else {
        	//fmt.Println(err)
        	//fmt.Printf("It's a already-made directory\n")
        //}
}

func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
    	return false
	}
    }
    return true
}

func IfDirectory(dirname string) bool {
    f, err := os.Open(dirname)
    if err != nil {
        //fmt.Println(err)
	return false
    }
    defer f.Close()
    fi, err := f.Stat()
    if err != nil {
        //fmt.Println(err)
	return false
    }
	switch mode := fi.Mode(); {
    	case mode.IsDir():
	  fmt.Printf("it's a dir\n")
	case mode.IsRegular():
	  fmt.Printf("It's a file\n")
	}
	return true
}
