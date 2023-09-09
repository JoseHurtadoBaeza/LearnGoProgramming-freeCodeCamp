package main

import "fmt"

//const a int16 = 27

// const (
// 	a = iota
// 	b
// 	c
// )

// const (
// 	a2 = iota
// )

// const (
// 	_ = iota + 5
// 	catSpecialist
// 	dogSpecialist
// 	snakeSpecialist
// )

// const (
// 	_  = iota // ignore first value by assigning to blank identifier
// 	KB = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
// )

// const (
// 	isAdmin = 1 << iota
// 	isHeadquarters
// 	canSeeFinancials

// 	canSeeAfrica
// 	canSeeAsia
// 	canSeeEurope
// 	canSeeNorthAmerica
// 	canSeeSouthAmerica
// )

func main() {

	const myConst int = 42
	fmt.Printf("%v, %T", myConst, myConst)

	// const a int = 14
	// const b string = "foo"
	// const c float32 = 3.14
	// const d bool = true
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%v\n", d)

	//fmt.Printf("%v, %T", a, a)

	// const a = 42
	// var b int16 = 27
	// fmt.Printf("%v, %T", a+b, a+b)

	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%v\n", a2)

	//var specialistType int
	//fmt.Printf("%v\n", catSpecialist)

	// fileSize := 4000000000.
	// fmt.Printf("%.2fGB", fileSize/GB)

	// var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	// fmt.Printf("%b\n", roles)
	// fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	// fmt.Printf("Is Admin? %v", isHeadquarters&roles == isHeadquarters)
}
