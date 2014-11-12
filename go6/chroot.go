package main

import ( 
	"fmt"
	//"strings"
	"os"
	//"log"
	//"errors"
)

func main() {
	//p := []int{2, 3, 5, 7, 11, 13}
	//fmt.Println("p ==", p)

	//qq := "whatever whatever whatever"
	//ww := "v"
	//ee := "o"

	//fmt.Println(p[0],p[1],p[2])
	//for i := 0; i < len(p); i++ {
	//	fmt.Println(float64(min(p)))
	//}
	//fmt.Println(strings.Contains(qq, ww))
	//fmt.Println(strings.Replace(qq, ww, ee, 2))

	//var vv string

	finfo, err := os.Stat("filename.txt")
	if err != nil {
    	// no such file or dir
	fmt.Println(err)
	fmt.Printf("Let's create the directory")
	vv := os.Mkdir("dude", 22)
	fmt.Println(vv)
    	return
	}
	if finfo.IsDir() {
    	// it's a file
	fmt.Printf("It's a file")
	
	} else {
    	// it's a directory
	fmt.Println(err)
	fmt.Printf("It's a already-made directory")
	}

        //vv := os.Mkdir("dude", 22)
	//if vv != nil {
	//fmt.Printf("No result")
	//log.Fatal(vv)
	//return
	//}

	//fmt.Println(vv)
	//v := Cmd( `/usr/bin/equo` , )
	//v.Path = "/usr/bin/equo"
	//v.Dir = "/usr/bin"
	
}

//func min(xs []int) int {
//    		m := xs[0]
//    		for i := 1; i < len(xs); i++ {
 //   			if xs[i] < m {
   // 				m = xs[i]
    //			}
    //		}
    //		return m
//}

type Cmd struct {
	Path string
	Dir string
}
