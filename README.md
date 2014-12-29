# Go Joiner

A generic `strings.Join` implementation that uses the `go generate` command from Go 1.4. See [this post for details](http://blog.golang.org/generate).

## Try the example

Install this package:

```
go install github.com/bslatkin/joiner
```

Go into the `example` directory:

```
go generate
go run *.go
```

You should see:

```
My favorite Simpsons Characters:
main.Person{FirstName:"Sideshow", LastName:"Bob", HairColor:"red"}
main.Person{FirstName:"Homer", LastName:"Simpson", HairColor:"n/a"}
main.Person{FirstName:"Lisa", LastName:"Simpson", HairColor:"blonde"}
main.Person{FirstName:"Marge", LastName:"Simpson", HairColor:"blue"}
main.Person{FirstName:"Mr", LastName:"Burns", HairColor:"gray"}
My favorite numbers:
10
33
24
47
My favorite places:
San Francisco
Katmandu
Sydney
```

## Generate your own joiners

Add this to the top of a file where you want to generate joiners:

```go
//go:generate joiner $GOFILE
```

Annotate your types with `// @joiner` to generate the code:

```go
// @joiner
type MyData struct{
    X float32
    Y float32
}
```

Then in your project directory run the generate command:

```
go generate
```

This will generate a set of files in the same directory with the generated code. If your types already satisfy `fmt.Stringer` in the same file then it won't auto-generate the stringer method.

```go
func (t MyData) String() string {
    return fmt.Sprintf("%#v", t)
}

type JoinMyData []MyData

func (j JoinMyData) With(sep string) string {
    all := make([]string, 0, len(j))
    for _, s := range j {
        all = append(all, s.String())
    }
    return strings.Join(all, sep)
}
```

To use the joiner, just do:

```go
JoinMyData(dataArray).With("\n")
```

The `Join*` prefixed type conversion will always be defined for your types. The `With` method indicates the separator string to use. The return type is a `string`.

## About

By [Brett Slatkin](http://www.onebigfluke.com)
