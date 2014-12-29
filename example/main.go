package main

import (
	"fmt"
	"strings"

	_ "github.com/bslatkin/joiner"
)

// go:generate joiner
type Person struct {
	FirstName string
	LastName  string
	HairColor string
}

func main() {
	people := []Person{
		Person{"Sideshow", "Bob", "red"},
		Person{"Homer", "Simpson", "n/a"},
		Person{"Lisa", "Simpson", "blonde"},
		Person{"Marge", "Simpson", "blue"},
		Person{"Mr", "Burns", "gray"},
	}
	fmt.Printf("My favorite Simpsons Characters:%s\n", JoinPerson(people).With("\n"))
}
