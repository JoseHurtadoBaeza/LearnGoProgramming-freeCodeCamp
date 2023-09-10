package main

import "fmt"

func main() {

	// a := 42
	// b := a // Copy whatever data stored in "a" and assigned to "b". Point to the same memory
	// fmt.Println(a, b)
	// a = 27
	// fmt.Println(a, b)

	// var a int = 42
	// var b *int = &a
	// fmt.Println(a, b)
	// fmt.Println(&a, b)
	// fmt.Println(a, *b)

	// a = 27
	// fmt.Println(a, *b)
	// *b = 14
	// fmt.Println(a, *b)

	// a := [3]int{1, 2, 3}
	// b := &a[0]
	// c := &a[1]
	// fmt.Printf("%v %p %p\n", a, b, c)

	var ms *myStruct
	fmt.Println(ms)
	//ms = &myStruct{foo: 42}
	ms = new(myStruct)
	//(*ms).foo = 42
	ms.foo = 42
	fmt.Println(ms)
	//fmt.Println((*ms).foo)
	fmt.Println(ms.foo)

	//a := [3]int{1, 2, 3}
	// a := []int{1, 2, 3} // But if we use [] then is a slice and works different
	// b := a
	// fmt.Println(a, b)
	// a[1] = 42
	// fmt.Println(a, b)

	a := map[string]string{"foo": "bar", "baz": "buz"} // Works like the slices
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)

}

type myStruct struct {
	foo int
}
