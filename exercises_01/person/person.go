package person

import "fmt"

type Person struct {
	Fname string
	Lname string
}

type SecretAgent struct {
	Person
	LicenseToKill bool
}

func (p Person) Speak() {
	fmt.Println("I am", p.Fname, p.Lname)
}

func (s SecretAgent) Speak() {
	fmt.Println("Code name", s.Fname)
}
