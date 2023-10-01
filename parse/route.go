package parse

import (
	"fmt"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fsctl/builder"
	"github.com/farseer-go/utils/condition"
	"go/ast"
	"strings"
)

// RouteComment 路由注解、函数类型
type RouteComment struct {
	Area           string            // 区域 @area
	Url            string            // 路由地址 @get {area}/order/{action}
	Method         string            // 路由method @get {area}/order/{action}
	filters        []string          // 过滤器 @filter jwt
	IocNames       map[string]string // 注入别名 @di repository default
	StatusMessage  string            // api返回message @message 成功
	PackagePath    string            // 包路径
	PackageName    string            // 包名
	FuncName       string            // handle名称
	paramName      []funcParam       // handle入参
	ProjectPath    string            // 项目根目录
	TopPackageName string            // 顶级包名
}

type funcParam struct {
	paramName     string // 参数名称
	paramTypeName string // 参数类型
	iocName       string // ioc别名
	typeName      string // 参数类型
}

// ParsePackageComment 解析包注解
func (receiver *RouteComment) ParsePackageComment(ant *Annotation) {
	if ant == nil {
		return
	}
	if ant.IsArea() {
		receiver.Area = ant.Args[0]
		return
	}
}

// ParseFuncComment 解析函数注解
func (receiver *RouteComment) ParseFuncComment(ant *Annotation) {
	if ant == nil {
		return
	}
	// 解析路由地址 @get @post @put @delete
	if ant.IsApi() {
		receiver.Method = strings.ToUpper(ant.Cmd)
		receiver.Url = ant.Args[0]
		return
	}

	// 解析过滤器 @filter
	if ant.IsFilter() {
		for _, arg := range ant.Args {
			receiver.filters = append(receiver.filters, arg+"{}")
		}
		return
	}

	// 解析过滤器 @di
	if ant.IsDi() {
		receiver.IocNames[ant.Args[0]] = ant.Args[1]
		return
	}

	// 解析返回Message @message
	if ant.IsMessage() {
		receiver.StatusMessage = ant.Args[0]
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
		var paramTypeName string
		var typeName string

		// 参数类型
		switch fieldType := field.Type.(type) {
		// 其它包的类型
		case *ast.SelectorExpr:
			packageName := fieldType.X.(*ast.Ident).Name
			paramTypeName = packageName + "." + fieldType.Sel.Name
			//if paramTypeName == "time.Time" {
			//	isBasic = true
			//	continue
			//}
			// 通过包名，去获取包路径
			typeName = receiver.parseType(astFile, packageName, fieldType.Sel.Name)
		case *ast.Ident:
			paramTypeName = fieldType.Name
			typeName = paramTypeName
		case *ast.ArrayType:
			paramTypeName = "[]" + fieldType.Elt.(*ast.Ident).Name
			typeName = paramTypeName
		default:
			paramTypeName = field.Names[0].Name
			typeName = paramTypeName
		}

		for _, fieldName := range field.Names {
			iocName := ""
			// 指定了ioc名称
			iocN, exists := receiver.IocNames[fieldName.Name]
			if exists {
				iocName = iocN
			}
			receiver.paramName = append(receiver.paramName, funcParam{
				paramName:     fieldName.Name, // 参数名
				paramTypeName: paramTypeName,  // 参数类型名称
				typeName:      typeName,       // 参数类型名称
				iocName:       iocName,
			})
		}
	}
}

// IsHaveComment 是否有解析到
func (receiver *RouteComment) IsHaveComment() bool {
	return receiver.Url != ""
}

func (receiver *RouteComment) parseType(astFile *ast.File, packageName string, paramTypeName string) string {
	var typeName string
	for _, importSpec := range astFile.Imports {
		// 去除前后""
		packagePath := strings.Trim(importSpec.Path.Value, "\"")
		if !strings.HasSuffix(packagePath, packageName) {
			continue
		}
		packagePath = strings.TrimPrefix(packagePath, receiver.TopPackageName)

		AstDirTypeDecl(receiver.ProjectPath+packagePath, func(filePath string, astFile *ast.File, genDecl *ast.GenDecl) {
			for _, specs := range genDecl.Specs {
				switch d := specs.(type) {
				case *ast.TypeSpec:
					if paramTypeName == d.Name.Name {
						switch t := d.Type.(type) {
						case *ast.InterfaceType:
							typeName = "interface"
							return
						case *ast.StructType:
							typeName = "struct"
							return
						case *ast.Ident:
							typeName = t.Name
							return
						}
					}
				}
			}
		})
		if typeName != "" {
			return typeName
		}
	}
	return typeName
}

// CheckIsRoute 检查route.go文件
func CheckIsRoute(routePath string) (isRoute bool) {
	// 检查根目录是否有route.go文件，如果有则删除
	AstFileGenDecl(routePath, func(genDecl *ast.GenDecl) {
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
func BuildRoute(routePath string, routeComments []RouteComment) {
	// 加载框架、应用的包
	lstPackageImport := loadImports(routeComments)

	// import
	var importBuilder strings.Builder
	lstPackageImport.Foreach(func(item *packageImportVO) {
		if importBuilder.Len() > 0 {
			importBuilder.WriteString("\n")
		}
		// 是否使用了包别名
		if item.alias == "" {
			importBuilder.WriteString(fmt.Sprintf("\t\"%s\"", item.packagePath))
		} else {
			importBuilder.WriteString(fmt.Sprintf("\t%s \"%s\"", item.alias, item.packagePath))
		}
	})

	builder.RouteBuilder(routePath, importBuilder.String(), func(routeItemTpl string) string {
		var routeBuilder strings.Builder
		for _, comment := range routeComments {
			if routeBuilder.Len() > 0 {
				routeBuilder.WriteString("\n")
			}
			contents := strings.ReplaceAll(routeItemTpl, "{method}", comment.Method)
			contents = strings.ReplaceAll(contents, "{url}", comment.Url)
			contents = strings.ReplaceAll(contents, "{funcName}", comment.PackageName+"."+comment.FuncName)
			contents = strings.ReplaceAll(contents, "{message}", comment.StatusMessage)
			contents = strings.ReplaceAll(contents, "{filters}", strings.Join(comment.filters, ","))
			// 函数的入参
			var paramBuilder strings.Builder
			for i := 0; i < len(comment.paramName); i++ {
				// 基础类型，直接使用参数名称，非基础类型，使用ioc别名
				paramName := condition.IsTrue(comment.paramName[i].typeName != "interface", comment.paramName[i].paramName, comment.paramName[i].iocName)
				paramBuilder.WriteString(fmt.Sprintf("\"%s\"", paramName))

				if i < len(comment.paramName)-1 {
					paramBuilder.WriteString(", ")
				}
			}
			contents = strings.ReplaceAll(contents, "{param}", paramBuilder.String())
			routeBuilder.WriteString(contents)
		}
		return routeBuilder.String()
	})
}

type packageImportVO struct {
	packagePath string // 包路径
	alias       string // 包别名
}

// 加载框架、应用的包
func loadImports(routeComments []RouteComment) collections.List[packageImportVO] {
	// 添加框架的包
	imports := collections.NewList[string]("github.com/farseer-go/webapi", "github.com/farseer-go/webapi/context")
	if collections.NewList(routeComments...).Where(func(item RouteComment) bool {
		for _, filter := range item.filters {
			if strings.HasPrefix(filter, "filter.") {
				return true
			}
		}
		return false
	}).Any() {
		imports.Add("github.com/farseer-go/webapi/filter")
	}
	// 添加应用的包
	for _, rc := range routeComments {
		imports.Add(rc.PackagePath)
	}

	lstPackageImport := collections.NewList[packageImportVO]()
	imports.Distinct().OrderByItem().Foreach(func(item *string) {
		lstPackageImport.Add(packageImportVO{packagePath: *item})
	})

	// 检查包是否有同名
	lstPackageNameTemp := collections.NewList[string]()
	for i := 0; i < lstPackageImport.Count(); i++ {
		packageNames := strings.Split(lstPackageImport.Index(i).packagePath, "/")
		packageName := packageNames[len(packageNames)-1:][0]
		// 说明包同名了，那么需要使用别名
		if lstPackageNameTemp.Contains(packageName) {
			vo := packageImportVO{
				packagePath: lstPackageImport.Index(i).packagePath,
				alias: packageName + parse.ToString(lstPackageNameTemp.Where(func(item string) bool {
					return item == packageName
				}).Count()+1),
			}

			lstPackageImport.Set(i, vo)
			for ir, _ := range routeComments {
				if routeComments[ir].PackagePath == vo.packagePath {
					routeComments[ir].PackageName = vo.alias
				}
			}
		}
		lstPackageNameTemp.Add(packageName)
	}
	return lstPackageImport
}
