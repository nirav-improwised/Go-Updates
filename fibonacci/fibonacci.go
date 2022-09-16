package main

import "fmt"

func fibo(num int) []int {
	m := 1
	n := 1
	temp := 0
	var arr = []int{}
	if num > 0 {
		arr = append(arr, m)
	}
	if num > 1 {
		arr = append(arr, n)
	}
	if num > 2 {
		for i := 2; i < num; i++ {
			arr = append(arr, m+n)
			temp = m
			m = n
			n += temp
		}
	}
	return arr
}

func main() {
	var n int
	fmt.Println("Enter no. of elements: ")
	fmt.Scan(&n)
	fmt.Println(fibo(n))
}
