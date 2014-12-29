// Copyright 2014 Brett Slatkin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

//go:generate joiner $GOFILE

import (
	"fmt"
)

// @joiner
type Person struct {
	FirstName string
	LastName  string
	HairColor string
}

// @joiner
type MyType int

// @joiner
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
