package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("existing : end without / should be true")
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/go/src`))
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/go/src/github.com`))
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/go/src/github.com/guitarrapc`))

	fmt.Println("existing : end with / should be true")
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/go/src/`))
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/go/src/github.com/`))
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/go/src/github.com/guitarrapc/`))

	fmt.Println("not exists folder should be false")
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/go/src/hoge`))
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/go/src/hogemoge`))
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/go/src/fugafuga/`))

	fmt.Println("is file should be false")
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/.gitconfig`))
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/.gitignore_global`))
	fmt.Println(IsExists(`C:/Users/ikiru.yoshizaki/.scoop`))
}

// isExists
func IsExists(path string) bool {
	if f, err := os.Stat(path); !os.IsNotExist(err) {
		return f.IsDir()
	}
	return false
}
