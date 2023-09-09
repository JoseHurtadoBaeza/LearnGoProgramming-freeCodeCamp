package main

import (
	"fmt"
)

// type Doctor struct {
// 	actorName  string
// 	number     int
// 	companions []string
// 	episodes   []string
// }

// With this way we are exporting the data
// type Doctor struct {
// 	ActorName  string
// 	Number     int
// 	Companions []string
// 	Episodes   []string
// }

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

// type Animal2 struct {
// 	Name   string `required max:"100"`
// 	Origin string
// }

func main() {

	// MAPS

	//statePopulations := make(map[string]int)
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	} // The order of the elements is not guaranteed in maps

	//m := map[[3]int]string{}
	fmt.Println(statePopulations)
	statePopulations["Georgia"] = 10310371
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Ohio"])
	fmt.Println(statePopulations["Georgia"]) // 0
	fmt.Println(statePopulations["Oho"])     // 0

	pop, ok := statePopulations["Oho"]
	fmt.Println(pop, ok)

	pop, ok = statePopulations["Ohio"]
	fmt.Println(pop, ok)

	_, ok = statePopulations["Ohio"]
	fmt.Println(ok)

	fmt.Println(len(statePopulations))

	sp := statePopulations
	delete(sp, "Ohio") // Also remove the element in the map "statePopulations"
	fmt.Println(sp)
	fmt.Println(statePopulations)

	// STRUCTS

	// aDoctor := Doctor{
	// 	number:    3,
	// 	actorName: "Jon Pertwee",
	// 	episodes:  []string{},
	// 	companions: []string{
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Sarah Jane Smith",
	// 	},
	// }

	// fmt.Println(aDoctor)
	// fmt.Println(aDoctor.actorName)
	// fmt.Println(aDoctor.companions[1])

	aDoctor := struct{ name string }{name: "John Pertwee"}
	// anotherDoctor := &aDoctor // Example of pointer
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)

	// b := Bird{}
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 48
	// b.CanFly = false
	// fmt.Println(b.Name)

	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b.Name)

	// t := reflect.TypeOf(Animal2{})
	// field, _ := t.FieldByName("Name")
	// fmt.Println(field.Tag)

}
