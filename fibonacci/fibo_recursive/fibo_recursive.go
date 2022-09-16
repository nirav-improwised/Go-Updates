package main

import "fmt"

func fibonacci(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else if num == 2 {
		return 1
	} else {
		return fibonacci(num-2) + fibonacci(num-1)
	}
}

func main() {
	var num int
	fmt.Println("Enter no. of elements: ")
	fmt.Scan(&num)
	fmt.Println(fibonacci(num))
}
