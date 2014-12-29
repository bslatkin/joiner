// generated by joiner -- DO NOT EDIT
package main

import (
	"fmt"
	"strings"
)

func (t Person) String() string {
	return fmt.Sprintf("%#v", t)
}

type JoinPerson []Person

func (j JoinPerson) With(sep string) string {
	all := make([]string, 0, len(j))
	for _, s := range j {
		all = append(all, s.String())
	}
	return strings.Join(all, sep)
}

func (t MyType) String() string {
	return fmt.Sprintf("%#v", t)
}

type JoinMyType []MyType

func (j JoinMyType) With(sep string) string {
	all := make([]string, 0, len(j))
	for _, s := range j {
		all = append(all, s.String())
	}
	return strings.Join(all, sep)
}

type JoinPlace []Place

func (j JoinPlace) With(sep string) string {
	all := make([]string, 0, len(j))
	for _, s := range j {
		all = append(all, s.String())
	}
	return strings.Join(all, sep)
}
