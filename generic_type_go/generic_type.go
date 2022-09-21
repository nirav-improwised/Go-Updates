// -------------Using Generic type-----------------
package main

import (
	"fmt"
)

type worker string

// ----below code of block will work with generic types-----
func Work[T any](w T) {
	fmt.Printf("%v is working\n", w)
}

func DoWork[T any](things []T) {
	for _, v := range things {
		fmt.Printf("%T\n", v)
		Work(v)
	}
}

func main() {
	var a, b, c worker
	a = "A"
	b = "B"
	c = "C"
	DoWork([]worker{a, b, c})
}

//---------Go dev original example, below code canNOT be used with Generic type------------------
// package main

// import (
// 	"fmt"
// )

// type worker string

// func (w worker) Work(){
// 	fmt.Printf("%s is working\n", w)
// }

// func DoWork[T any](things []T) {
//     for _, v := range things {
//         v.Work()
//     }
// }

// func main() {
// 	var a,b,c worker
// 	a = "A"
// 	b = "B"
// 	c = "C"
// 	DoWork([]worker{a,b,c})
// }
