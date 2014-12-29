// ./bad_example2.go:51: cannot use people (type []Person) as type []fmt.Stringer in argument to Join
//
// This doesn't work because an array of interfaces is not equivalent to an
// array of the type that implements that interface.
//
// See http://research.swtch.com/interfaces for more details.

package main

import (
	"fmt"
	"strings"
)

type Joinable []fmt.Stringer

func Join(in []fmt.Stringer) Joinable {
	out := make(Joinable, 0, len(in))
	for _, x := range in {
		out = append(out, x)
	}
	return out
}

func (j Joinable) With(sep string) string {
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
