package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile(`^D:/GitHub/guitarrapc/MixedContentChecker/csharp/src/.*/bin/.+/netcoreapp2.2$`)
	dirs, err := dirwalk(`D:\GitHub\guitarrapc\MixedContentChecker`, true)
	if err != nil {
		fmt.Println(err)
	}
	for _, dir := range dirs {
		fmt.Println(pattern.MatchString(dir), dir)
	}
}

func dirwalk(path string, toSlash bool) (fullPaths []string, err error) {
	fullPaths = nil
	files, _err := ioutil.ReadDir(path)
	if _err != nil {
		err = _err
		return
	}

	for _, file := range files {
		if file.IsDir() {
			if toSlash {
				fullPaths = append(fullPaths, filepath.ToSlash(filepath.Join(path, file.Name())))
			} else {
				fullPaths = append(fullPaths, filepath.Join(path, file.Name()))
			}

			walkPaths, _err := dirwalk(filepath.Join(path, file.Name()), toSlash)
			fullPaths = append(fullPaths, walkPaths...)
			err = _err
			continue
		}
	}
	return
}
