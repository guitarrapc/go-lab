package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	fmt.Println("0", flag.Arg(0))
	fmt.Println("1", flag.Arg(1))
	fmt.Println("2", flag.Arg(2))
}
