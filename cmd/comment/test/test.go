package test

import (
	"fmt"
	"runtime"
)

func PrintGQYDebug() {
	if _, file, line, ok := runtime.Caller(1); ok {
		fmt.Printf("\033[1;33;47m%s\033[0m", "[GQY DEBUG]")
		fmt.Printf("file: %s, line %d\n", file, line)
	}
}
