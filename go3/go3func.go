package go3func

import (
	"fmt"
	"runtime"
)

func oSYS(q string) string {
	switch q := runtime.GOOS; q {
	default:
	}
		return q
}
