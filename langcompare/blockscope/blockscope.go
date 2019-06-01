package main

import "strings"

type Person struct {
	Name string
}
type People []Person

func (pp People) Each(hfn func(i int, p Person)) {
	for i, p := range pp {
		hfn(i, p)
	}
}
func main() {
	people := People{Person{"tom"}, Person{"bom"}}
	names := []string{}
	people.Each(func(i int, p Person) {
		names = append(names, p.Name)
	})
	println(strings.Join(names, ","))
}
