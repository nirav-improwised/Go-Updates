// package main

// import "fmt"

//	func main() {
//		fmt.Println("Welcome to Booking-App")
//		fmt.Println("Collect your tickets here")
//	}
//-------------------------------------------------------
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello World!")
// }
//----------------------------------------------------
// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var a string
// 	var b int
// 	var c bool

//		fmt.Println(a)
//		fmt.Println(b)
//		fmt.Println(c)
//	}
//
// ------------------------------------------
// package main

// import "fmt"

//	func main() {
//		var n int
//		n = 'r'
//		fmt.Print(n)
//	}
//
// --------------------------------------------
// package main

// import "fmt"

// var a int     //we can declare variables outside a function using "var"
// var b int = 2
// var c = 3

//	func main() {
//		a = 1
//		fmt.Println(a)
//		fmt.Println(b)
//		fmt.Print(c)//Why Does it show % sign at this place in the output???????????????????
//	}
//
// -----------------------------------------------------
// package main

// import (
//
//	"fmt"
//
// )
// a:=1//This line will give error because we cannot use := outside a function
//
//	func main() {
//		fmt.Println(a)
//	}
//
// -----------------------------------------------------
// -------------------------for loop program----------------------------
// package main

// import "fmt"

//	func main() {
//		var n int
//		fmt.Print("Enter a num:")
//		fmt.Scan(&n)
//		for i := 0; i < 10; i++ {
//			fmt.Printf("%v x %v = %v\n", n, (i + 1), (n * (i + 1)))
//		}
//	}
//
// -----------------------------------------------------------------
// --------------array of strings && for index, item loop----------------------
// package main

// import "fmt"

// var cars = [10]string{"Freestyle", "EcoSport", "Gurkha", "Virtus", "Defender", "Wrangler", "Bronco", "Kushaq", "Punch", "Harrier"}

//	func main() {
//		for index, car := range cars {
//			fmt.Printf("index:%v\tcar:%v\n", index, car)
//		}
//	}
//
// ------------------------------------------------------------------
// -------------------Infinite loop, conditional statement, break----------------------------
// package main

// import "fmt"

//	func main() {
//		var char string
//		for {
//			fmt.Print("Enter 'q' to end the loop or any other to continue:")
//			fmt.Scan(&char)
//			if char == "q" {
//				break
//			}
//		}
//	}
//
// --------------------Finding max of 3 numbers----------------------------
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, b, c float32
// 	fmt.Print("Enter 3 numbers: ")
// 	fmt.Scan(&a, &b, &c)
// 	if a >= b {
// 		if a >= c {
// 			fmt.Print(a)
// 		} else {
// 			fmt.Print(c)
// 		}
// 	} else if b >= c {
// 		fmt.Print(b)
// 	} else {
// 		fmt.Print(c)
// 	}
// }
//