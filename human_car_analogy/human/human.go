package human

type Human interface {
	Body_parts() []string
	Performance() string
}

type Man string

func (m Man) Body_parts() []string {
	p := []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
	return p
}

func (m Man) Performance() string {
	return "8 Hours/day"
}
