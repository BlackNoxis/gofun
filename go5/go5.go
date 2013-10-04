package main

import ( 
	"fmt"
	"strings"
	"os"
)

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	qq := "whatever whatever whatever"
	ww := "v"
	ee := "o"

	fmt.Println(p[0],p[1],p[2])
	for i := 0; i < len(p); i++ {
		fmt.Println(float64(min(p)))
	}
	fmt.Println(strings.Contains(qq, ww))
	fmt.Println(strings.Replace(qq, ww, ee, 2))
	
	vv := os.Mkdir("dude", 22)
	fmt.Println(vv)

	//v := Cmd( `/usr/bin/equo` , )
	//v.Path = "/usr/bin/equo"
	//v.Dir = "/usr/bin"
	
}

func min(xs []int) int {
    		m := xs[0]
    		for i := 1; i < len(xs); i++ {
    			if xs[i] < m {
    				m = xs[i]
    			}
    		}
    		return m
}

type Cmd struct {
	Path string
	Dir string
}
