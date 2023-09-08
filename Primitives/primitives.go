package main

import "fmt"

func main() {

	//var n bool = true

	// n := 1 == 1
	// m := 1 == 2

	// fmt.Printf("%v, %T\n", n, n)
	// fmt.Printf("%v, %T", m, m)

	// var n bool
	// fmt.Printf("%v, %T", n, n)

	// var n uint16 = 42
	// fmt.Printf("%v, %T", n, n)

	// a := 10
	// b := 3
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// fmt.Println(a % b)

	// var a int = 10
	// var b int8 = 3
	// fmt.Println(a + int(b))

	// a := 10             // In binary -> 1010
	// b := 3              // In binary -> 0011
	// fmt.Println(a & b)  // AND OPERATOR -> 0010
	// fmt.Println(a | b)  // OR OPERATOR -> 1011
	// fmt.Println(a ^ b)  // XOR OPERATOR -> 1001
	// fmt.Println(a &^ b) // AND NOT OPERATOR -> 1000

	// a := 8              // 2^3
	// fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	// fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0

	// var n float64 = 3.14
	// n = 13.7e72
	// n = 2.1e14
	// fmt.Printf("%v, %T", n, n)

	// a := 10.2
	// b := 3.7
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)

	// var n complex64 = 1 + 2i
	// fmt.Printf("%v, %T\n", n, n)

	// a := 1 + 2i
	// b := 2 + 5.2i
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)

	// var n complex128 = 1 + 2i
	// fmt.Printf("%v, %T\n", real(n), real(n))
	// fmt.Printf("%v, %T\n", imag(n), imag(n))

	// var n complex128 = complex(5, 12)
	// fmt.Printf("%v, %T\n", n, n)

	s := "this is a string"
	b := []byte(s) // Convertion to a slice of bytes
	fmt.Printf("%v, %T\n", b, b)

	//var r rune = 'a'
	//fmt.Printf("%v, %T", r, r)

}
