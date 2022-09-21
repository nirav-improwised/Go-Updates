package main

import (
	"fmt"
	h "human_car_analogy/human"
	v "human_car_analogy/vehicle"
)

func main() {
	fmt.Println("Human Car Analogy")
	var car v.Vehicle = v.Car("Volvo")
	// car = v.Car("Volvo")
	var human h.Human = h.Man("Programmer")
	// human = h.Man("Programmer")

	for index, item := range car.Parts() {
		fmt.Printf("%s <===========> %s\n", item, human.Body_parts()[index])
	}
}
