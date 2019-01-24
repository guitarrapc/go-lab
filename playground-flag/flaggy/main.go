package main

import "github.com/integrii/flaggy"

// Declare variables and their defaults
var (
	stringFlag = "defaultValue"
	intFlag    int
	actionFlag string
)

func main() {

	// Add a flag
	flaggy.String(&stringFlag, "s", "string", "A test string flag")
	flaggy.Int(&intFlag, "i", "int", "int value")
	flaggy.AddPositionalValue(&actionFlag, "action", 1, true, "action")

	// Parse the flag
	flaggy.Parse()

	// Use the flag
	println(stringFlag, intFlag, actionFlag)
}
