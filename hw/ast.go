package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
)

func main() {
	filePath := "/Users/linkw/go/src/awesomeProject/test/test1/test11"
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, filePath, func(info fs.FileInfo) bool {
		return true
	}, 0)
	if err != nil {
		panic(err)
	}
	for pName, pkg := range pkgs {
		fmt.Println(pName)
		fmt.Println(pkg)
	}
}


