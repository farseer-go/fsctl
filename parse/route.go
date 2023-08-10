package parse

import (
	"fmt"
	"go/ast"
	"strings"
)

// RouteComment 路由注解、函数类型
type RouteComment struct {
	Area          string            // 区域 @area
	Url           string            // 路由地址 @get {area}/order/{action}
	Method        string            // 路由method @get {area}/order/{action}
	filters       []string          // 过滤器 @filter jwt
	IocNames      map[string]string // 注入别名 @di repository default
	StatusMessage string            // api返回message @message 成功
	PackagePath   string            // 包路径
	PackageName   string            // 包名
	FuncName      string            // handle名称
	paramName     []funcParam       // handle入参
}

type funcParam struct {
	paramName string // 参数名称
	paramType string // 参数类型
	isBasic   bool   // 是否为基础变量
	iocName   string // ioc别名
}

// ParsePackageComment 解析包注解
func (receiver *RouteComment) ParsePackageComment(comment string) {
	if comment == "" {
		return
	}
	comments := strings.Split(comment, " ")
	if strings.ToLower(comments[0]) == "area" {
		receiver.Area = comments[1]
		return
	}
}

// ParseFuncComment 解析函数注解
func (receiver *RouteComment) ParseFuncComment(comment string) {
	if comment == "" {
		return
	}
	comments := strings.Split(comment, " ")
	// 解析路由地址 @get @post @put @delete
	if strings.ToLower(comments[0]) == "get" ||
		strings.ToLower(comments[0]) == "post" ||
		strings.ToLower(comments[0]) == "put" ||
		strings.ToLower(comments[0]) == "delete" {
		receiver.Method = strings.ToUpper(comments[0])
		receiver.Url = comments[1]
		return
	}

	// 解析过滤器 @filter
	if strings.ToLower(comments[0]) == "filter" {
		receiver.filters = comments[1:]
		return
	}

	// 解析过滤器 @di
	if strings.ToLower(comments[0]) == "di" && len(comments) == 3 {
		receiver.IocNames[comments[1]] = comments[2]
		return
	}

	// 解析返回Message @message
	if strings.ToLower(comments[0]) == "message" {
		receiver.StatusMessage = comments[1]
		return
	}
}

// ParseFuncType 解析函数名称、参数类型
func (receiver *RouteComment) ParseFuncType(astFile *ast.File, funcDecl *ast.FuncDecl) {
	// 函数的名称
	receiver.PackageName = astFile.Name.Name
	receiver.FuncName = funcDecl.Name.Name
	// 解析函数的入参
	for _, field := range funcDecl.Type.Params.List {
		var paramType string
		var isBasic bool
		// 参数类型
		switch fieldType := field.Type.(type) {
		case *ast.SelectorExpr:
			paramType = fieldType.X.(*ast.Ident).Name + "." + fieldType.Sel.Name
			isBasic = paramType == "time.Time"
		case *ast.Ident:
			paramType = fieldType.Name
			isBasic = true
		}

		for _, fieldName := range field.Names {
			iocName := ""
			// 指定了ioc名称
			iocN, exists := receiver.IocNames[fieldName.Name]
			if exists {
				iocName = iocN
			}
			receiver.paramName = append(receiver.paramName, funcParam{
				paramName: fieldName.Name, // 参数名
				paramType: paramType,      // 参数类型
				isBasic:   isBasic,        // 是否为基础变量
				iocName:   iocName,
			})
		}
	}
}

// IsHaveComment 是否有解析到
func (receiver *RouteComment) IsHaveComment() bool {
	return receiver.Url != ""
}

// CheckIsRoute 检查route.go文件
func CheckIsRoute(routePath string) (isRoute bool) {
	// 检查根目录是否有route.go文件，如果有则删除
	ASTGenDecl(routePath, func(genDecl *ast.GenDecl) {
		for _, spec := range genDecl.Specs {
			switch s := spec.(type) {
			case *ast.ValueSpec:
				// 是否包含 route 变量
				if len(s.Names) == 1 && s.Names[0].Obj.Kind == ast.Var && s.Names[0].Name == "route" && len(s.Values) == 1 {
					if compositeLit, isCompositeLit := s.Values[0].(*ast.CompositeLit); isCompositeLit {
						if arrayType, isArrayType := compositeLit.Type.(*ast.ArrayType); isArrayType {
							if selector, isSelectorExpr := arrayType.Elt.(*ast.SelectorExpr); isSelectorExpr {
								if ident, isIdent := selector.X.(*ast.Ident); isIdent {
									isRoute = ident.Name == "webapi" && selector.Sel.Name == "Route"
								}
							}
						}
					}
				}
			}
		}
	})
	return
}

// BuildRoute 生成route.go文件
func BuildRoute(routePath string, routeComments []RouteComment) string {
	// 引用包（使用map，为了去重）
	imports := make(map[string]any)
	imports["github.com/farseer-go/webapi"] = struct{}{}
	for _, rc := range routeComments {
		imports[rc.PackagePath] = struct{}{}
	}

	var builder strings.Builder
	builder.WriteString("// 该文件由fsctl route命令自动生成，请不要手动修改此文件\n")
	builder.WriteString("package main\n")
	// import
	builder.WriteString("\n")
	for packName := range imports {
		builder.WriteString(fmt.Sprintf("import \"%s\"\n", packName))
	}
	builder.WriteString("\n")
	// var route = []webapi.Route{ }
	builder.WriteString("var route = []webapi.Route{\n")

	for _, comment := range routeComments {
		builder.WriteString(fmt.Sprintf("\t{"))
		builder.WriteString(fmt.Sprintf("Method: \"%s\", ", comment.Method))
		builder.WriteString(fmt.Sprintf("Url: \"%s\", ", comment.Url))
		builder.WriteString(fmt.Sprintf("Action: %s.%s, ", comment.PackageName, comment.FuncName))
		builder.WriteString(fmt.Sprintf("Message: \"%s\", ", comment.StatusMessage))
		// 函数的入参
		builder.WriteString(fmt.Sprintf("Params: []string{"))
		for i := 0; i < len(comment.paramName); i++ {
			// 基础类型，直接使用参数名称
			if comment.paramName[i].isBasic {
				builder.WriteString(fmt.Sprintf("\"%s\"", comment.paramName[i].paramName))
			} else { // 非基础类型，使用ioc别名
				builder.WriteString(fmt.Sprintf("\"%s\"", comment.paramName[i].iocName))
			}
			if i < len(comment.paramName)-1 {
				builder.WriteString(", ")
			}
		}
		builder.WriteString(fmt.Sprintf("}"))

		builder.WriteString(fmt.Sprintf("},\n"))
	}

	builder.WriteString("}\n")
	return builder.String()
}
