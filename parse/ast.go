package parse

import (
	"github.com/farseer-go/utils/file"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func AstDirFuncDecl(path string, fund func(filePath string, astFile *ast.File, funcDecl *ast.FuncDecl)) {
	files := file.GetFiles(path, "*.go", true)
	for _, filePath := range files {
		paths := strings.Split(filePath, "/")
		fileName := paths[len(paths)-1:][0]
		if strings.HasPrefix(fileName, "_") || strings.HasSuffix(filePath, "_test.go") {
			continue
		}

		// 解析Func定义
		fileSet := token.NewFileSet()
		f, _ := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
		for _, decl := range f.Decls {
			switch d := decl.(type) {
			case *ast.FuncDecl:
				fund(filePath, f, d)
			}
		}
	}
}

func AstDirTypeDecl(path string, fund func(filePath string, astFile *ast.File, genDecl *ast.GenDecl)) {
	files := file.GetFiles(path, "*.go", true)
	for _, filePath := range files {
		paths := strings.Split(filePath, "/")
		fileName := paths[len(paths)-1:][0]
		if strings.HasPrefix(fileName, "_") || strings.HasSuffix(filePath, "_test.go") {
			continue
		}

		// 解析Func定义
		fileSet := token.NewFileSet()
		f, _ := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
		for _, decl := range f.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				fund(filePath, f, d)
			}
		}
	}
}

func AstFileGenDecl(filePath string, fund func(genDecl *ast.GenDecl)) {
	fileSet := token.NewFileSet()
	astFile, _ := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
	for _, decl := range astFile.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			fund(d)
		}
	}
}
