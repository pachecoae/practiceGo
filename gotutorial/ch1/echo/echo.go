package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	s, sep, sa := "", " ", make([]string, 100)
	for index := 0; index < 100; index++ {
		sa[index] = strconv.Itoa(index)
	}

	// Testing print times when looping
	start := time.Now()
	for _, val := range sa {
		s += sep + val
	}
	fmt.Printf("%s\n", s)
	fmt.Printf("%s%f\n", "It's been: ", time.Since(start).Seconds())

	// Testing print times when joining
	start = time.Now()
	fmt.Printf("%s\n", strings.Join(sa, " "))
	fmt.Printf("%s%f\n", "It's been: ", time.Since(start).Seconds())

	// Printing Args
	for index := 0; index < len(os.Args); index++ {
		fmt.Println("Index ", index, ": ", os.Args[index])
	}

}
