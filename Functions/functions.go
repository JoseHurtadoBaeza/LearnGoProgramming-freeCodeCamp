package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}

	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)

	//sum("The sum is", 1, 2, 3, 4, 5)
	s := sum(1, 2, 3, 4, 5)
	//fmt.Println("The sum is ", s)
	//fmt.Println("The sum is ", *s)
	fmt.Println("The sum is ", s)

	// d, err := divide(5.0, 3.0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(d)

	for i := 0; i < 5; i++ {

		func(i int) {
			//msg := "Hello Go!"
			fmt.Println(i)
		}(i) // Immediately invoke the function

	}

	//f := func() {
	var f func() = func() {
		fmt.Println("Hello Go!")
	}
	f()

	// var divide func(float64, float64) (float64, error)
	// divide = func(a, b float64) (float64, error) {
	// 	if b == 0.0 {
	// 		return 0.0, fmt.Errorf("Cannot divide by zero")
	// 	} else {
	// 		return a / b, nil
	// 	}
	// }
	// d, error := divide(5.0, 3.0)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }
	// fmt.Println(d)

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is:", g.name)

}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

//func sum(msg string, values ...int) { // Told the go runtime to take all the arguments that have passed in and wrapped up in a slice named "values"
// func sum(values ...int) int {
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	//fmt.Println(msg, result)
// 	return result
// }

// func sum(values ...int) *int {
// 	fmt.Println(values)
// 	result := 0 // Declared on the execution stack of this function
// 	for _, v := range values {
// 		result += v
// 	}
// 	//fmt.Println(msg, result)
// 	return &result
// } // When this function exits that execution stack is destroyed, that memory is freed up

func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	//fmt.Println(msg, result)
	return
} // When this function exits that execution stack is destroyed, that memory is freed up

// func divide(a, b float64) (float64, error) {
// 	if b == 0.0 {
// 		//panic("Cannot provide zero as second value")
// 		return 0.0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// Struct as a "class" in object oriented
type greeter struct {
	greeting string
	name     string
}

// Method example
// func (g greeter) greet() {
// 	fmt.Println(g.greeting, g.name)
// 	g.name = ""
// }

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
