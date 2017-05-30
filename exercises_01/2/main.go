package main

import (
	"fmt"

	"github.com/urbq/GolangWebDev/exercises_01/person"
)

func main() {
	pe := person.Person{Fname: "Max", Lname: "Hase"}
	sa := person.SecretAgent{Person: person.Person{Fname: "Bim", Lname: "Bo"}, LicenseToKill: true}
	fmt.Println(pe.Fname)
	pe.Speak()
	fmt.Println(sa.Fname)
	sa.Speak()
	sa.Person.Speak()
}
