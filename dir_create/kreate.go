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
	   } else { os.Exit(1) }
	   if f.Name() == "bin" {
		fmt.Printf("oh hello, %v \n", f.Name())
	   } else { os.Exit(1) }
	}
}

func ifYES(i string) bool {
	if i == "Yes" || i == "YES" || i == "Y" || i == "y" || i == "Ye" || i == "YE" || i == "yes" || i == "ye" {
		return true
	}
	return false
}

func sourceCH(dir string) {
	current_dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	kurrent_dir := strings.TrimSpace(current_dir)
	os.Chdir(dir)
	url := "http://pkg.rogentos.ro/~rogentos/iso/Gentoo-Devel-x64" + ".tar.xz"
	downloadFromUrl(url)
	os.Chdir(kurrent_dir)
}

func omega(i string) {
	fmt.Printf("Hello there, name a directory that you want to see if it exists or not:\n")
	reader := bufio.NewReader(os.Stdin)
	instring, _ := reader.ReadString('\n')
	existenta := strings.TrimSpace(instring)

	if Exists(existenta) {
		fmt.Printf("Yay. It's there. Something.\n")
		dirname := existenta + string(filepath.Separator)
		if IfDirectory(dirname) == true {
			checkCH(dirname)
		}
	} else {
		fmt.Printf("meh. it`s not there, Joe. Let`s start the creation procedure. Do you want to create it? Yes/No\n")
		fmt.Scanf("%s", &i)
		if ifYES(i) == true {
			os.Mkdir(existenta, 22)
			fmt.Printf("Do you wish to download the chroot archive? Yes/No\n")
			fmt.Scanf("%s", &i)
			if ifYES(i) == true {
				sourceCH(existenta)
			} else {
				fmt.Printf("Okay. we will not download the chroot archive\n")
			}
		} else {
			fmt.Printf("Have fun, then\n")
		}
	}
}
