package main

import (
	"fmt"
	"strconv"
)

//var i float32 = 42

// var (
// 	actorName    string = "Elisabeth Sladen"
// 	companion    string = "Sarah Jane Smith"
// 	doctorNumber int    = 3
// 	season       int    = 11
// )

// var (
// 	counter int = 0
// )

//var i int = 27

// Uppercase variables (outs block) makes the go compiler to expose this variable to the outside world.
//var I int = 42

func main() {

	// var i int
	// i = 42

	//var j float32 = 27

	//k := 99.

	//fmt.Println(i)
	//fmt.Printf("%v, %T", i, i)

	// fmt.Println(i)
	// var i int = 42
	//	i = 13
	// fmt.Println(i)

	//var theHTTP string = "https://google.com" // Example of name convention

	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)

}
