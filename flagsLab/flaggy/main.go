package main

import (
	"fmt"

	"github.com/integrii/flaggy"
)

// Declare variables and their defaults
var (
	stringFlag = "defaultValue"
	intFlag    int
	actionFlag string
)

// go run main.go hogemnoge -s=hoge -i=1
// go run main.go -s hoge -i 1 fff
// go run main.go -s=hoge -i=1 fff
func main() {

	// Add a flag
	flaggy.String(&stringFlag, "s", "string", "A test string flag")
	flaggy.Int(&intFlag, "i", "int", "int value")
	flaggy.AddPositionalValue(&actionFlag, "action", 1, true, "action")
	flaggy.Parse()

	// Use the flag
	fmt.Println(stringFlag, intFlag, actionFlag)
}
