package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for fileNum, url := range os.Args[1:] {
		go fetch(url, ch, fileNum) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // recieve from channel ch
	}

	end := time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed\n", end)
}

func fetch(url string, ch chan<- string, fileNum int) {

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	// If outputing to a file for inspection, uncomment the following block
	// fileName := "fileName" + string(fileNum) + ".txt"
	// out, err := os.Create(fileName)
	// if err != nil {
	// 	ch <- fmt.Sprintf("while creating file %s: %v", fileName, err)
	// }
	// nbytes, err := io.Copy(out, resp.Body)

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
