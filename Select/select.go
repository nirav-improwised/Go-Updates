package main

import (
	"fmt"
	"time"
)

func source1(ch chan string) {
	ch <- "Data from source1"
}
func source2(ch chan string) {
	ch <- "Data from source2"
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go source1(ch1)
	go source2(ch2)
	time.Sleep(250 * time.Millisecond)
	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	default:
		fmt.Println("This is default")
	}
}
