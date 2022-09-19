//-----------------------using function inside structure-------------------------------------------------------
// package main

// import "fmt"

// // type fit func(float32, float32)string

// type Person struct {
// 	name string
// 	fit  func(float32, float32) string
// }

// func main() {
// 	var weight, height float32
// 	fmt.Println("Enter weight(kg) and height(feet)")
// 	fmt.Scan(&weight, &height)
// 	p1 := Person{name: "Nirav",
// 		fit: func(w float32, h float32) string {
// 			if w >= (h * 10) {
// 				return "fit"
// 			} else {
// 				return "unfit"
// 			}
// 		},
// 	}
// 	fmt.Println("p1 name:", p1.name)
// 	fmt.Println("p1 fitness:", p1.fit(weight, height))
// }

// ------------------------------inheritance/composition with struct & also using interface-------------------------------------
// package main

// import "fmt"

// type see_details interface {
// 	details()
// }

// type Person struct {
// 	name        string
// 	age         int
// 	nationality string
// }

// type fav_car struct {
// 	car string
// 	Person
// }

// type fav_bike struct {
// 	bike string
// 	Person
// }

// func (c fav_car) details() {
// 	fmt.Println("Name:", c.name)
// 	fmt.Println("age:", c.age)
// 	fmt.Println("nationality:", c.nationality)
// 	fmt.Println("fav_car:", c.car)
// }
// func (b fav_bike) details() {
// 	fmt.Println("Name:", b.name)
// 	fmt.Println("age:", b.age)
// 	fmt.Println("nationality:", b.nationality)
// 	fmt.Println("fav_bike:", b.bike)
// }

// func print_details(s []see_details) {
// 	for _, item := range s {
// 		item.details()
// 	}
// }
// func main() {
// 	p1 := Person{"Nirav", 22, "Indian"}
// 	c1 := fav_car{"Volvo S60 D4", p1}
// 	b1 := fav_bike{"Pulser 200 NS", p1}
// 	i1 := []see_details{c1, b1}
// 	print_details(i1)
// }

// ------------------------------------multiple inheritance with structure-------------------------------------
// package main

// import "fmt"

// type male struct {
// 	male_count int
// }
// type female struct {
// 	female_count int
// }
// type transGender struct {
// 	transGender_count int
// }
// type children struct {
// 	children_count int
// }
// type population struct {
// 	male
// 	female
// 	transGender
// 	children
// }

// func (p population) see_population() {
// 	fmt.Println(p.male_count + p.female_count + p.transGender_count + p.children_count)
// }
// func main() {
// 	m := male{500}
// 	f := female{1000}
// 	t := transGender{400}
// 	c := children{2000}
// 	p := population{m, f, t, c}
// 	p.see_population()
// }
//--------------------------------------------------------------------------------