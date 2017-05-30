package main

import "github.com/urbq/GolangWebDev/exercises_01/person"

type human interface {
	Speak()
}

func talk(h human) {
	h.Speak()
}

func main() {
	pe := person.Person{Fname: "Max", Lname: "Hase"}
	sa := person.SecretAgent{Person: person.Person{Fname: "Bim", Lname: "Bo"}, LicenseToKill: true}
	talk(pe)
	talk(sa)
}
