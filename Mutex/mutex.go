package main

import (
	"fmt"
	"sync"
)

var x int

func incr(wg *sync.WaitGroup, m *sync.Mutex) {
	// m.Lock()
	x = x + 1
	// m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 100; i++ {
		go incr(&wg, &m)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("Final value of x is", x)
}

//
// ---------------Solving race conditions with channel
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var x int = 0

// func incr(wg *sync.WaitGroup, ch chan bool) {
// 	ch <- true
// 	x = x + 1
// 	<-ch
// 	wg.Done()
// }
// func main() {
// 	ch := make(chan bool, 1)
// 	var wg sync.WaitGroup
// 	for i := 0; i < 100; i++ {
// 		go incr(&wg, ch)
// 		wg.Add(1)
// 	}
// 	wg.Wait()
// 	fmt.Println("Final value of x is", x)
// }
