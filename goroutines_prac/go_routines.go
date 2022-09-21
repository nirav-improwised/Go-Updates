// --------------Multiple Goroutine----------
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func numbers() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Print(i, " ")
// 		time.Sleep(50 * time.Millisecond)
// 	}
// }

// func spellings() {
// 	spell := [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
// 	for _, word := range spell {
// 		fmt.Print(word, " ")
// 		time.Sleep(75 * time.Millisecond)
// 	}
// }

// func main() {
// 	go numbers()
// 	go spellings()
// 	time.Sleep(2500 * time.Millisecond)
// 	fmt.Println("\n main Goroutine terminated")
// }

// --------------------------------First program Goroutine---------------
package main

import (
	"fmt"
	"time"
)

func Greet() {
	// time.Sleep(2 * time.Second)
	fmt.Println("Hello User! Have a Good Day.")
}
func main() {
	go Greet()
	time.Sleep(1 * time.Second)
	fmt.Println("main goroutine terminated")
}
