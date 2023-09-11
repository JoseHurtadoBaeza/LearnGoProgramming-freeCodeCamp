package main

import (
	"fmt"
	"time"
)

//var wg = sync.WaitGroup{}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func main() {

	// ch := make(chan int)
	// wg.Add(2)
	// go func() { // Receiving data by the channel, this is the receiving go routine
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	// go func() { // The sending go routine
	// 	i := 42
	// 	ch <- i
	// 	i = 27
	// 	wg.Done()
	// }()
	// wg.Wait()

	// ch := make(chan int)
	// for j := 0; j < 5; j++ {
	// 	wg.Add(2)
	// 	go func() {
	// 		i := <-ch
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()
	// 	go func() {
	// 		ch <- 42
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	// ch := make(chan int)
	// wg.Add(2)
	// go func(ch <-chan int) { // Receiving only channel
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	//ch <- 27
	// 	wg.Done()
	// }(ch)
	// go func(ch chan<- int) { // Sending only channel
	// 	ch <- 42
	// 	//fmt.Println(<-ch)
	// 	wg.Done()
	// }(ch)
	// wg.Wait()

	// ch := make(chan int, 50) // Create an internal data store that can store in this case, 50 integers.
	// wg.Add(2)
	// go func(ch <-chan int) { // Receiving only channel
	// 	for i := range ch {
	// 		fmt.Println(i)
	// 	}
	// 	wg.Done()
	// }(ch)
	// go func(ch chan<- int) { // Sending only channel
	// 	ch <- 42
	// 	ch <- 27
	// 	close(ch)
	// 	wg.Done()
	// }(ch)
	// wg.Wait()

	// ch := make(chan int, 50)
	// wg.Add(2)
	// go func(ch <-chan int) {
	// 	for {
	// 		if i, ok := <-ch; ok {
	// 			fmt.Println(i)
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	wg.Done()
	// }(ch)
	// go func(ch chan<- int) {
	// 	ch <- 42
	// 	ch <- 27
	// 	close(ch)
	// 	wg.Done()
	// }(ch)

	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}

}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
