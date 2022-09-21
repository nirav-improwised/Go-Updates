// -------#Program 5: This program prints the (sum of sqaures)+(sum of cubes) of all digits in a number-----------------
package main

import "fmt"

func digits(num int, dCH chan int) {
	for num != 0 {
		digit := num % 10
		dCH <- digit
		num = num / 10
	}
	close(dCH)
}

func squares(num int, sCH chan int) {
	sdCH := make(chan int)
	sum := 0
	go digits(num, sdCH)
	for i := range sdCH {
		sum += i * i
	}
	sCH <- sum
	close(sCH)
}

func cubes(num int, cCH chan int) {
	cdCH := make(chan int)
	sum := 0
	go digits(num, cdCH)
	for i := range cdCH {
		sum += (i * i * i)
	}
	cCH <- sum
	close(cCH)
}

func main() {
	var n int
	fmt.Println("Enter a number:")
	fmt.Scan(&n)
	sqchnl := make(chan int)
	cbchnl := make(chan int)
	go squares(n, sqchnl)
	go cubes(n, cbchnl)
	square, cube := <-sqchnl, <-cbchnl
	fmt.Println("Square+Cubes of digits is: ", square+cube)
}

// ------#4 Program to demonstrate use of for range with channel-----
// package main

// import "fmt"

// func print_n(ch_n chan int) {
// 	for i := 0; i < 10; i++ {
// 		ch_n <- i
// 	}
// 	close(ch_n)
// }

// func main() {
// 	ch := make(chan int)
// 	go print_n(ch)
// 	for n := range ch {
// 		fmt.Println(n)
// 	}
// }

//-----#3 Program--Closing channel--Infinite-----------------
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func producer(chnl chan int) {
// 	for i := 0; i < 10; i++ {
// 		chnl <- i
// 	}
// 	close(chnl)
// }

//		func main() {
//			ch := make(chan int)
//			go producer(ch)
//			for {
//				v := <-ch
//				fmt.Println(v)
//				time.Sleep(500 * time.Millisecond)
//			}
//	}
//
// ------#2 Program for channel, it returns the sqaure of the length of a string------
// package main

// import (
// 	"fmt"
// )

// func calcLength(st string, lCh chan int) {
// 	lCh <- len(st)
// }
// func calcSqaure(num int, sqch chan int) {
// 	sqch <- (num * num)
// }

// func main() {

// 	var st_user string
// 	lCh := make(chan int)
// 	sqch := make(chan int)
// 	fmt.Println("Enter a string: ")
// 	fmt.Scan(&st_user)
// 	go calcLength(st_user, lCh)
// 	len_st := <-lCh
// 	go calcSqaure(len_st, sqch)
// 	sq_len_st := <-sqch
// 	fmt.Printf("lenght of string is %d and sqaure of %d is %d\n", len_st, len_st, sq_len_st)

// }

//-----#1 Program Channel----------------
// package main

// import "fmt"

// func main() {
// 	var a chan int
// 	if a == nil {
// 		fmt.Printf("a_value:%v & a_type:%T\n", a, a)
// 		a = make(chan int)
// 		fmt.Println("After declaring 'a' using make function")
// 		fmt.Printf("a_value:%d & a_type:%T\n", a, a)
// 	}
// }
