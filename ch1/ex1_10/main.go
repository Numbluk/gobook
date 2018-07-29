// Fetchall fetches URLs in parallel and reports their times and sizes
// This exercise is to see if this is run successively if the calls are
// cached (quicker). This also prints the results to a file.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Ignoring file writing for now, concurrently writing to files
// will come later
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[2:] {
		go fetch(url, ch) // start a goroutine
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to open: %v", err)
	}

	defer file.Close()
	for range os.Args[2:] {
		// file.Write([]byte(<-ch))
		file.WriteString(<-ch)
		// fmt.Println(<-ch)
	}
	var finalTime = fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	file.WriteString(finalTime)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
