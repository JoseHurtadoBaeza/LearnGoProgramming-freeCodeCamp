package main

import (
	"fmt"
	"log"
)

func main() {

	// fmt.Println("start")
	// defer fmt.Println("middle")
	// fmt.Println("end")

	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// a := "start"
	// defer fmt.Println(a) // start
	// a = "end"

	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// fmt.Println("start")
	// panic("something bad happened")
	// fmt.Println("end") // code unreachable

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error()) // This line doesn't allow the port be reused in another execution of the program
	// }

	// fmt.Println("start")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("Error:", err)
	// 	}
	// }()
	// panic("something bad happened")
	// fmt.Println("end")

	fmt.Println("start")
	panicker()
	fmt.Println("end")

}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			//panic(err)
		}
	}()
	panic("something bad happened")
	//fmt.Println("done panicking") // code unreachable
}
