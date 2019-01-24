package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Int("int", 0, "int flag")
	flag.String("str", "default", "string flag")
	flag.Bool("bool", false, "bool flag")
	flag.Parse()
	fmt.Println(flag.NArg(), flag.NFlag())
}
