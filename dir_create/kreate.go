package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
	"path/filepath"
)

func checkCH(dir string) {
	d, err := os.Open(dir)
	if err != nil {
		os.Exit(1)
	}
	f, err := d.Readdir(-1)
	if err != nil {
		os.Exit(1)
	}
	for _, f := range f {
	   if f.Name() == "root" {
		fmt.Printf("oh hello, %v \n", f.Name())
	   }
	   if f.Name() == "bin" {
		fmt.Printf("oh hello, %v \n", f.Name())
	   }
	}
}

func omega() {
	fmt.Printf("Hello there, name a directory that you want to see if it exists or not:\n")
	reader := bufio.NewReader(os.Stdin)

	instring, _ := reader.ReadString('\n')

	existenta := strings.TrimSpace(instring)

	if Exists(existenta) {
		fmt.Printf("yay. it's there. something.\n")
		dirname := existenta + string(filepath.Separator)
		if IfDirectory(dirname) == true {
			checkCH(dirname)
		}
	} else {
		var i string
		fmt.Printf("meh. it`s not there, Joe. Let`s start the creation procedure. Do you want to create it? Yes/No\n")
		fmt.Scanf("%s", &i)
		//i_yes := []string{"Y","yes","Yes","Ye","YES","YeS","Ye","y"} //I've gone mad
		//i_no := []string{"No","N","NO","nO"} //and again
		if i == "Yes" || i == "YES" || i == "Y" || i == "y" || i == "Ye" || i == "YE" || i == "yes" || i == "ye" {
			os.Mkdir(existenta, 22)
			fmt.Printf("Do you wish to download the chroot archive? Yes/No\n")
			fmt.Scanf("%s", &i)
			if i == "Yes" || i == "YES" || i == "Y" || i == "y" || i == "Ye" || i == "YE" || i == "yes" || i == "ye" {
				current_dir, err := os.Getwd()
				kurrent_dir := strings.TrimSpace(current_dir) //let's translate this into readable string
				os.Chdir(existenta)
				//fmt.Println(kurrent_dir) //additional verify
				url := "http://pkg.rogentos.ro/~rogentos/iso/Gentoo-Devel-x64" + ".tar.gz"
				go downloadFromUrl(url)
				if err != nil {
					os.Exit(1)
				} else {
					os.Chdir(kurrent_dir)
					//fmt.Println(kurrent_dir) //additional verify
				}
			} else {
				fmt.Printf("Okay. we will not download the chroot archive\n")
			}
		} else {
			fmt.Printf("Have fun, then\n")
		}
	}
}
