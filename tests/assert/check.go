package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("existing : end without / should be true")
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/go/src`))
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/go/src/github.com`))
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/go/src/github.com/guitarrapc`))

	fmt.Println("existing : end with / should be true")
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/go/src/`))
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/go/src/github.com/`))
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/go/src/github.com/guitarrapc/`))

	fmt.Println("not exists folder should be false")
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/go/src/hoge`))
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/go/src/hogemoge`))
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/go/src/fugafuga/`))

	fmt.Println("is file should be false")
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/.gitconfig`))
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/.gitignore_global`))
	fmt.Println(isExists(`C:/Users/ikiru.yoshizaki/.scoop`))
}

// isExists
func isExists(path string) bool {
	if f, err := os.Stat(path); !os.IsNotExist(err) {
		return f.IsDir()
	}
	return false
}
