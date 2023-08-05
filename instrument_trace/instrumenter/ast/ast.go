package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
)

func Rewrite(filename string) ([]byte, error) {
	fset := token.NewFileSet()
	oldAST, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s: %w", filename, err)
	}

	if !hasFuncDecl(oldAST) {
		return nil, nil
	}

	astutil.AddImport(fset, oldAST, "github.com/NickYxy/GoTypingTest/instrument")

	addDeferTraceIntoFuncDecls(oldAST)

	buf := &bytes.Buffer{}
	err = format.Node(buf, fset, oldAST)
	if err != nil {
		return nil, fmt.Errorf("error formating new code: %w", err)
	}

	return buf.Bytes(), nil
}

func addDeferTraceIntoFuncDecls(f *ast.File) {
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if ok {
			addDeferStmt(fd)
		}
	}
}

func addDeferStmt(fd *ast.FuncDecl) (added bool) {
	stmts := fd.Body.List

	for _, stmt := range stmts {
		ds, ok := stmt.(*ast.DeferStmt)
		if !ok {
			continue
		}

		ce, ok := ds.Call.Fun.(*ast.CallExpr)
		if !ok {
			continue
		}

		se, ok := ce.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		x, ok := se.X.(*ast.Ident)
		if !ok {
			continue
		}

		if (x.Name == "functrace") && (se.Sel.Name == "Trace") {
			return false
		}
	}

	ds := &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "functrace"},
					Sel: &ast.Ident{Name: "Trace"},
				},
			},
		},
	}

	newList := make([]ast.Stmt, len(stmts)+1)
	copy(newList[1:], stmts)
	newList[0] = ds
	fd.Body.List = newList
	return true
}

func hasFuncDecl(f *ast.File) bool {
	if len(f.Decls) == 0 {
		return false
	}

	for _, decl := range f.Decls {
		_, ok := decl.(*ast.FuncDecl)
		if ok {
			return true
		}
	}

	return false
}
