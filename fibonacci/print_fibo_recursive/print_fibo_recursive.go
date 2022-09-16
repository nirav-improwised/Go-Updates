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

func fibo(num int) {
	for i := 0; i <= num; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}

func main() {
	var num int
	fmt.Println("Enter no. of elements: ")
	fmt.Scan(&num)
	fibo(num)
}
