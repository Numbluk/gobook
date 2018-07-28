// Prints the count, text of lines, and name of the file for lines
// that appear more than once in the named input files
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	file_appearances := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, file_appearances)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1_4: %v\n", err)
				continue
			}
			countLines(f, counts, file_appearances)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println("Files dups appear in:", strings.Join(file_appearances[line], ", "))
			fmt.Printf("%d\t%s\n\n", n, line)
		}
	}
}

func file_in(file string, file_appearances []string) bool {
	for _, file_name := range file_appearances {
		if file == file_name {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, file_appearances map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !file_in(f.Name(), file_appearances[input.Text()]) {
			file_appearances[input.Text()] = append(file_appearances[input.Text()], f.Name())
		}
	}
}

