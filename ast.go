package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func AST(a string) {
	fileSet := token.NewFileSet()
	pkgs, _ := parser.ParseDir(fileSet, "./", func(info os.FileInfo) bool {
		return true
	}, parser.ParseComments)
	ast.Print(fileSet, pkgs)

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				decl.Pos()
				decl.End()
			}
		}
	}

	fmt.Println(a)
}
