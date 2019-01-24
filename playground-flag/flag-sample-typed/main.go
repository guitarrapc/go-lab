package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		i int
		s string
		b bool
	)
	flag.IntVar(&i, "int", 0, "int flag")
	flag.StringVar(&s, "str", "default", "string flag")
	flag.BoolVar(&b, "bool", false, "bool flag")
	flag.Parse()
	fmt.Println(i, s, b)
}
