package go4func

import (
        "fmt"
        "runtime"
)

func oSYS(q string) string {
        switch q := runtime.GOOS; q {
	case "linux":
		fmt.Println("Linux.")		
        default:
        }
                return q
}

func suma (x, y int) int {
        sum := x + y
        return sum

}

func adauga() func(int,int) int {
	ne := 0
	return func(x, y int) int {
		ne = x + y
		return ne
	}
}
