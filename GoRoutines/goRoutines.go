package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{} // Design to synchronize multiple go routines together
var counter = 0
var m = sync.RWMutex{}

func main() {

	// go sayHello()
	// time.Sleep(100 * time.Millisecond)

	// var msg = "Hello"
	// wg.Add(1)
	// go func(msg string) {
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }(msg)
	// msg = "Goodbye"
	// time.Sleep(100 * time.Millisecond)
	// wg.Wait()

	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()

}

// func sayHello() {
// 	fmt.Println("Hello")
// }

func sayHello() {
	m.RLock()
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
