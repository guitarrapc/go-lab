package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile(filepath.ToSlash(`^D:\GitHub\guitarrapc\MixedContentChecker\csharp\src\.*\bin\.+\netcoreapp2.2$`))
	dirs := dirwalkDirectory(`D:\GitHub\guitarrapc\MixedContentChecker`, true)
	for _, dir := range dirs {
		fmt.Println(pattern.MatchString(dir), dir)
	}
}

func dirwalkDirectory(dir string, toSlash bool) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			if toSlash {
				paths = append(paths, filepath.ToSlash(filepath.Join(dir, file.Name())))
			} else {
				paths = append(paths, filepath.Join(dir, file.Name()))
			}

			paths = append(paths, dirwalkDirectory(filepath.Join(dir, file.Name()), toSlash)...)
			continue
		}
	}

	return paths
}
