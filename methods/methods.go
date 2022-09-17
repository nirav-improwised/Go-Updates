// -----------------Using Pointer Receivers----------------------
package main

import "fmt"

type car_stock struct {
	punch   int
	nexon   int
	harrier int
	safari  int
}

func (stock *car_stock) update_stock(p, n, h, s int) {
	(*stock).punch += p
	(*stock).nexon += n
	(*stock).harrier += h
	(*stock).safari += s
}
func main() {
	jay_ganesh := car_stock{250, 200, 150, 100}
	fmt.Println("Before Update", jay_ganesh)
	jay_ganesh.update_stock(-50, -60, -30, -20)
	fmt.Println("After Update:", jay_ganesh)
}

// -------------Using Receivers-------------------
// package main

// import "fmt"

// type car_stock struct {
// 	punch   int
// 	nexon   int
// 	harrier int
// 	safari  int
// }

// func (stock car_stock) update_stock(p, n, h, s int) {
// 	stock.punch += p
// 	stock.nexon += n
// 	stock.harrier += h
// 	stock.safari += s
// }
// func main() {
// 	jay_ganesh := car_stock{250, 200, 150, 100}
// 	fmt.Println("Before Update", jay_ganesh)
// 	jay_ganesh.update_stock(-50, -60, -30, -20)
// 	fmt.Println("After Update:", jay_ganesh)
// }
