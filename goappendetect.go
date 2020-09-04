package goappendetect

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "goappendetect is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "goappendetect",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			params := n.Type.Params.List
			sliceArgs := sliceParams(pass, params)
			if len(sliceArgs) != 0 {
				checkAppend(pass, n.Body.List, sliceArgs)
			}
		}
	})

	return nil, nil
}

// sliceParams returns argument objects whose types are types.Slice
func sliceParams(pass *analysis.Pass, fields []*ast.Field) map[types.Object]bool {
	defs := pass.TypesInfo.Defs
	slices := make(map[types.Object]bool)
	for _, f := range fields {
		for _, n := range f.Names {
			obj, ok := defs[n]
			if !ok {
				continue
			}

			// pointer types do not pass here
			if _, ok = obj.Type().Underlying().(*types.Slice); !ok {
				continue
			}
			slices[obj] = true
		}
	}
	return slices
}

// checkAppend checks if there are assignments as follows: s = append(s, 1)
func checkAppend(pass *analysis.Pass, stmts []ast.Stmt, slices map[types.Object]bool) {
	uses := pass.TypesInfo.Uses
	for _, stmt := range stmts {
		assign, ok := stmt.(*ast.AssignStmt)
		if !ok {
			continue
		}
		for i, lh := range assign.Lhs {
			lh, ok := lh.(*ast.Ident)
			if !ok {
				continue
			}
			lobj, ok := uses[lh]
			if !ok {
				continue
			}
			// check if the left side is one of the args of the function
			if !slices[lobj] {
				continue
			}

			rh := assign.Rhs[i]
			call, ok := rh.(*ast.CallExpr)
			if !ok {
				continue
			}
			if call.Fun.(*ast.Ident).Name != "append" {
				continue
			}
			// check if the first arg of append is the same as the left side
			arg := call.Args[0].(*ast.Ident)
			robj := uses[arg]
			if !slices[robj] || lobj.Id() != robj.Id() {
				continue
			}
			pass.Reportf(stmt.Pos(), "this assignment is not detected outside of the func")
		}
	}
}
