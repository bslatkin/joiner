// ./bad_example.go:40: cannot convert people (type []Person) to type Join
//
// It's not immediately clear why this doesn't work. See bad_example2.go

package main

import (
	"fmt"
	"strings"
)

type Join []fmt.Stringer

func (j Join) With(sep string) string {
	stred := make([]string, 0, len(j))
	for _, s := range j {
		stred = append(stred, s.String())
	}
	return strings.Join(stred, sep)
}

type Person struct {
	FirstName string
	LastName  string
	HairColor string
}

func (s *Person) String() string {
	return fmt.Sprintf("%#v", s)
}

func main() {
	people := []Person{
		Person{"Sideshow", "Bob", "red"},
		Person{"Homer", "Simpson", "n/a"},
		Person{"Lisa", "Simpson", "blonde"},
		Person{"Marge", "Simpson", "blue"},
		Person{"Mr", "Burns", "gray"},
	}
	fmt.Printf("My favorite Simpsons Characters:%s\n", Join(people).With("\n"))
}
