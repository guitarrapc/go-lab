package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

var helpFlag bool
var fooFlag string
var barFlag string

func main() {
	flag.BoolVarP(&helpFlag, "help", "p", false, "show help message")
	flag.StringVar(&fooFlag, "foo", "defoo", "help message foo")
	flag.StringVarP(&barFlag, "bar", "b", "debar", "help message bar")
	flag.Parse()

	if helpFlag {
		flag.PrintDefaults()
		return
	}

	fmt.Println(fooFlag, barFlag)
}
