// -------------------Printing concrete Type and value of interface----------------------
package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

// salary of permanent employee is the sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
		fmt.Printf("\nType=%T\tValue:%v", v, v)
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

}

// ------------------Interface embedding & composition, failed example, Inappropriate---------------------------------
// package main

// import "fmt"

// type car_technicalities interface {
// 	engine()
// 	brakes()
// 	drive()
// 	transmission()
// 	suspensions()
// }
// type car_features interface {
// 	comfort_convinience()
// 	infotainment()
// }
// type car_specifications interface {
// 	car_technicalities
// 	car_features
// }

// type car_tech struct {
// 	cengine, cbrakes, cdrive, ctransmission, csuspensions string
// }

// type car_ft struct {
// 	cinfotainment, crear_ac_vents, cadjustable_hr string
// }

// func (ct car_tech) engine() {
// 	fmt.Println(ct.engine)
// }
// func (ct car_tech) brakes() {
// 	fmt.Println(ct.brakes)
// }
// func (ct car_tech) drive() {
// 	fmt.Println(ct.drive)
// }
// func (ct car_tech) transmission() {
// 	fmt.Println(ct.transmission)
// }
// func (ct car_tech) suspensions() {
// 	fmt.Println(ct.suspensions)
// }

// func (ct car_ft) infotainment() {
// 	fmt.Println(ct.infotainment)
// }
// func (ct car_ft) comfort_convinience() {
// 	fmt.Println(ct.crear_ac_vents)
// 	fmt.Println(ct.cadjustable_hr)
// }

// func details(cs []car_specifications) {

// 	cs.engine()
// 	cs.brakes()
// 	cs.drive()
// 	cs.transmission()
// 	cs.suspensions()
// 	cs.comfort_convinience()
// 	cs.infotainment()
// }

// func main() {
// 	ct := car_tech{"Punch", "1.2L Petrol", "Disc & Drum", "2W Front", "MT & AMT", "McPherson"}
// 	cf := car_ft{"Touch, Apple car play & Android auto", "YES", "YES"}
// 	car_specs := car_specifications{ct, cf}
// 	details(car_specs)
// }
