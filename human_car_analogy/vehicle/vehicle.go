package vehicle

type Vehicle interface {
	Parts() []string
	Speed() string
}

type Car string

func (c Car) Parts() []string {
	r := []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Tank"}
	return r
}

func (c Car) Speed() string {
	return "80 kmph"
}
