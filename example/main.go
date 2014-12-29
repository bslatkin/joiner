package main

import (
	"fmt"

	_ "github.com/bslatkin/joiner"
)

// go:generate joiner
type Person struct {
	FirstName string
	LastName  string
	HairColor string
}

// go:generate joiner
type MyType int

// go:generate joiner
type Place struct {
	Name string
	Lat  float32
	Lon  float32
}

func (p Place) String() string {
	return p.Name
}

type LeaveThisAlone struct {
	Foo string
	Bar int
}

func main() {
	people := []Person{
		Person{"Sideshow", "Bob", "red"},
		Person{"Homer", "Simpson", "n/a"},
		Person{"Lisa", "Simpson", "blonde"},
		Person{"Marge", "Simpson", "blue"},
		Person{"Mr", "Burns", "gray"},
	}
	fmt.Printf("My favorite Simpsons Characters:\n%s\n", JoinPerson(people).With("\n"))

	myInts := []MyType{MyType(10), MyType(33), MyType(24), MyType(47)}
	fmt.Printf("My favorite numbers:\n%s\n", JoinMyType(myInts).With("\n"))

	places := []Place{
		Place{"San Francisco", 37.7833, -122.4167},
		Place{"Katmandu", 27.7000, 85.3333},
		Place{"Sydney", -33.8600, 151.2094},
	}
	fmt.Printf("My favorite places:\n%s\n", JoinPlace(places).With("\n"))
}
