package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"regexp/syntax"
	"strings"

	regexast "github.com/guitarrapc/go-lab/regexLab/regexast"
)

func main() {
	args := []string{
		`^D:/GitHub/guitarrapc/MixedContentChecker/csharp/src/.*/bin/.+/netcoreapp2.2$`,
		`^D:/GitHub/ghoasd.*`,
		`^D:/asdf\d{0}/ghoasd.*`,
		`^C:/Users/ikiru\.yoshizaki/Documents/Git/guitarrapc/Log4NetConfigurations/src/.*/bin/Debug$`,
	}
	for _, arg := range args {
		basePath, err := getBasePath(arg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		pattern := regexp.MustCompile(arg)
		dirs, err := dirwalk(basePath, true)
		if err != nil {
			fmt.Println(err)
		}
		for _, dir := range dirs {
			fmt.Println(pattern.MatchString(dir), dir)
		}
	}
}

func getBasePath(path string) (string, error) {
	asts, err := regexast.ParseRegex(path, syntax.Perl)
	if err != nil {
		//fmt.Println(err)
		return "", err
	}

	var b strings.Builder
	begin := false
	// for _, a := range asts {
	// 	fmt.Println(a)
	// }
	for _, a := range asts {
		if begin && !a.IsRune {
			break
		}
		if begin && a.IsRune {
			b.WriteString(a.Value)
		}
		if a.IsStart {
			begin = true
		}
	}
	return b.String(), nil
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
