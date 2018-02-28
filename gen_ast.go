package main

import (
	"fmt"
	"strings"
    "io/ioutil"
	"os"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	os.Exit(actualMain())
}

func actualMain() int {
	files, err := projectFiles()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println(files)

	for _, path := range files {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		ast.Print(fset, f)
	}

	return 0
}

func projectFiles() (files []string, err error) {
	fileModes, err := ioutil.ReadDir(".")
    if err != nil {
		return
    }

	for _, f := range fileModes {
		if f.Mode().IsRegular() {
			name := f.Name()
			if name == "gen_ast.go" {
				continue //FIXME: remove check when no longer hacking
			}
			if strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go") {
				files = append(files, name)
			}
		}
    }
	return
}
