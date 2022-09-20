package main

import "fmt"

// type num_type int
type num struct {
	a, b int
}

func division_func(a, b int) (float32, error) {
	if b != 0 {
		return float32(a / b), nil
	} else {
		return float32(-999), &num{a, b}
	}
}

func main() {
	num1 := num{5, 0}
	ans, err := division_func(num1.a, num1.b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}
func (n *num) Error() string {
	return fmt.Sprintf("ZeroDivisionError %d/%d is mathematically undefined", n.a, n.b)
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// type MyError struct {
// 	When time.Time
// 	What string
// }

// func (e *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s",
// 		e.When, e.What)
// }

// func run() error {
// 	return &MyError{
// 		time.Now(),
// 		"it didn't work",
// 	}
// }

// func main() {
// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}
// }
