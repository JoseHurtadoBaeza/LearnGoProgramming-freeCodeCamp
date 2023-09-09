package main

import "fmt"

func main() {

	// ARRAYS

	//grades := [...]int{97, 85, 93}
	// var students [3]string
	// fmt.Printf("Students: %v\n", students)
	// students[0] = "Lisa"
	// students[2] = "Charlie"
	// students[1] = "Arnold"
	// fmt.Printf("Student #1: %v\n", students[1])
	// fmt.Printf("Number of Students: %v\n", len(students))

	// var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// fmt.Println(identityMatrix)

	// var identityMatrix [3][3]int
	// identityMatrix[0] = [3]int{1, 0, 0}
	// identityMatrix[1] = [3]int{0, 1, 0}
	// identityMatrix[2] = [3]int{0, 0, 1}
	// fmt.Println(identityMatrix)

	// a := [...]int{1, 2, 3}
	// b := a // b := &a this is a pointers example with arrays
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)

	// SLICES

	// a := []int{1, 2, 3}
	// b := a
	// b[1] = 5 // It is also modified in the array "a"
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[:]   // slice of all elements
	// c := a[3:]  // slice from the 4th element to end
	// d := a[:6]  // slice first 6 elements
	// e := a[3:6] // slice the 4th, 5th, and 6th elements
	// a[5] = 42
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// a := make([]int, 3, 100) // Type, length and capacity
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// a := []int{}
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 1)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// //a = append(a, 2, 3, 4, 5)
	// a = append(a, []int{2, 3, 4, 5}...)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]
	b := append(a[:2], a[3:]...)
	fmt.Println(b) // [1 2 4 5]
	fmt.Println(a) // [1 2 4 5 5]
}
