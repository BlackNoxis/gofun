package main

import (
        "fmt"
        //"strings"
        "os"
        //"errors"
)

func main() {
	
	if Exists("dude") {
		fmt.Printf("yay\n")
		} else {
		var i string
		fmt.Scanf("meh. you want to create it? Yes/No\n", &i)
		i_yes := []string{"Y","yes","Yes","Ye","YES","YeS","Ye","y"}
		i_no := []string{"No","N","NO","nO"}
		aq := len(i_yes)
		zq := len(i_no)
		
	}

	ww := os.Mkdir("dude", 22)

        if fileinfo.IsDir() {
        	fmt.Printf("It's already created. Let's see if it's a file or a directory\n")
        } else {
        	fmt.Println(err)
        	fmt.Printf("It's a already-made directory\n")
        }
}

func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
    	return false
	}
    }
    return true
}
