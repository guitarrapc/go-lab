package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"regexp/syntax"
	"strings"
	"unicode/utf8"

	regexast "github.com/guitarrapc/go-lab/regexLab/regexast"
)

func main() {
	args := []string{
		// `^D:/GitHub/guitarrapc/MixedContent.*/csharp/src/.*/bin/.+/netcoreapp2.2$`,
		// `^D:/GitHub/guitarrapc/MixedContentChecker/csharp/src/.*/bin/.+/netcoreapp2.2$`,
		// `^D:/GitHub/ghoasd/`,
		// `^D:/GitHub/ghoasd世界/.*`,
		// `^D:/GitHub/ghoasd.*`,
		// `^D:/asdf\d{0}/ghoasd.*`,
		// `^D:/GitHub/ghoasd.*hogemoge/fugafuga/.*gua/`,
		// `^C:/Users/ikiru\.yoshizaki/Documents/Git/guitarrapc/Log4NetConfigurations/src/.*/bin/Debug$`,
		// `^C:/Users/guitarrapc/Downloads/watchdog/logfiles`,
		// "^C:/Users/guitarrapc/Downloads/watchdog/logfiles",
		// `^C:/Users/ikiru\.yoshizaki/go/src/github\.com/guitarrapc/watchdog-symlinker/logfiles.*$`,
		// "^C:/Users/ikiru\\.yoshizaki/go/src/github\\.com/guitarrapc/watchdog-symlinker/logfiles.*$",
		`^C:/Users/ikiru\.yoshizaki/go/src/github\.com/guitarrapc/watchdog-symlinker/logfiles.*$`,
	}
	for _, arg := range args {
		basePath, err := getBasePath(arg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(basePath)

		pattern := regexp.MustCompile(arg)
		dirs, err := dirwalk(basePath, true)
		dirs = append(dirs, basePath)
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

	// for _, a := range asts {
	// 	fmt.Println(a)
	// }

	var b strings.Builder
	start := false
	begin := false
	var s string
	for _, a := range asts {
		if a.IsStart {
			start = true
		}
		if start && a.IsRune {
			b.WriteString(a.Value)
			begin = true
		}
		if begin && !a.IsRune {
			// check path is valid and fix
			s = b.String()
			fmt.Println(s)
			if isExists(s) {
				break
			}
			if getLastRune(s, 1) != "/" {
				i := strings.LastIndex(s, "/") + 1
				s = string(s[:i])
			}
			break
		}
	}

	return s, nil
}

func dirwalk(path string, toSlash bool) (fullPaths []string, err error) {
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

func getLastRune(s string, c int) string {
	j := len(s)
	for i := 0; i < c && j > 0; i++ {
		_, size := utf8.DecodeLastRuneInString(s[:j])
		j -= size
	}
	return s[j:]
}

func substring(str string, start int, length int) string {
	if start < 0 || length <= 0 {
		return str
	}
	r := []rune(str)
	if start+length > len(r) {
		return string(r[start:])
	}

	return string(r[start : start+length])
}

func isExists(path string) bool {
	if f, err := os.Stat(path); !os.IsNotExist(err) {
		return f.IsDir()
	}
	return false
}
