package main

import (
	"fmt"
	"path/filepath"
	"bufio"
	"os"
	"strings"
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
	fmt.Printf("Hello there, name a directory that you want to see if it exists or not:\n")
	reader := bufio.NewReader(os.Stdin)

	instring, _ := reader.ReadString('\n')

	existenta := strings.TrimSpace(instring)

	if Exists(existenta) {
		fmt.Printf("yay. it's there. something.\n")
		dirname := existenta + string(filepath.Separator)
		if IfDirectory(dirname) == true {
			fmt.Println(dirname)
			d, err := os.Open(dirname)
			if err != nil {
				//fmt.Println(err) //error check
				os.Exit(1)
			}
			f, err := d.Readdir(-1)
			if err != nil {
				//fmt.Println(err) //error check
				os.Exit(1)
			}
			for _, f := range f {
				fmt.Println(f.Name(), f.Size()) //nope, not needed
				// I should add a mode  isregular ? with an `if` //or not
				if f.Name() == "root" {
					fmt.Printf("oh hello root\n")
					rootz := dirname + "root" + string(filepath.Separator)
					edd, err := os.Open(rootz)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					passf, err := edd.Readdir(-1)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					for _, passf := range passf {
						fmt.Println(f.Name() + "/" + passf.Name())
					}
					continue
				}
				if f.Name() == "bin" {
					fmt.Printf("oh hello bin\n")
				}
			 }
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

