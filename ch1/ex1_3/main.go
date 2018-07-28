// This measures the differences in time it takes to echo command line
// arguments using various methods vs strings.Join
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("For loop efficiency: ")
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args[1:]); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("Args:", s)
	fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
	fmt.Println("")

	fmt.Println("Range loop efficiency: ")
	start = time.Now()
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("Args:", s)
	fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
	fmt.Println("")

	fmt.Println("strings.Join efficiency: ")
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}