package main

import (
	"fmt"
	"math"
)

func main() {

	// IF STATEMENTS

	if true {
		fmt.Println("The test is true")
	}

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	//fmt.Println(pop) // ERROR: Out of her scope

	number := 50
	guess := 30
	if guess < 1 || returnTrue() {
		fmt.Println("The guess must be greater than 1!")
	} else if guess < 1 || guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}

		fmt.Println(number <= guess, number >= guess, number != guess)

	}
	fmt.Println(!true)

	myNum := 0.123456789
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	// SWITCH STATEMENTS

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	j := 10
	switch {
	case j <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case j <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	var i interface{} = 1 // The type interface can take any type pf data that we have in a Go application
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		// break
		// fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}

}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
