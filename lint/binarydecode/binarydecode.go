// Package binarydecode defines an Analyzer that reports *[]byte passthroughs
// that skip the content-type check.
//
// The SDK client hands a response body to the caller as raw bytes when the
// decode target is a *[]byte. Without a content-type check a JSON, XML or HTML
// body that arrives with a success status would be returned as if it were an
// image or an archive.
package binarydecode

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// CheckFunc is the package-level function that must guard every *[]byte passthrough.
const CheckFunc = "isBinaryContentType"

const doc = `report *[]byte type assertions that are not guarded by ` + CheckFunc + `

A type assertion or type switch case to *[]byte hands a response body through
as raw bytes. The branch it guards must call ` + CheckFunc + ` from the same
package, so that only image and archive responses are passed through:

	if data, ok := v.(*[]byte); ok {
		if !` + CheckFunc + `(contentType) {
			return fmt.Errorf("unexpected content type %q for a binary response", contentType)
		}
		*data = b
		return nil
	}

The guarded branch is the condition and body of the if statement whose
initializer makes the assertion, the case clause of a type switch, or else the
enclosing function.`

// Analyzer reports *[]byte passthroughs that skip the content-type check.
var Analyzer = &analysis.Analyzer{
	Name:     "binarydecode",
	Doc:      doc,
	URL:      "https://github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/tree/v4/lint/binarydecode",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

var bytesPointer = types.NewPointer(types.NewSlice(types.Typ[types.Byte]))

func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{(*ast.TypeAssertExpr)(nil), (*ast.CaseClause)(nil)}

	// stack ends with n itself.
	insp.WithStack(nodeFilter, func(n ast.Node, push bool, stack []ast.Node) bool {
		if !push {
			return true
		}
		switch n := n.(type) {
		case *ast.TypeAssertExpr:
			// n.Type is nil for the x.(type) guard of a type switch; its
			// cases are checked one by one as *ast.CaseClause.
			if n.Type != nil && isBytesPointer(pass, n.Type) {
				checkGuarded(pass, n, guardedRegion(stack))
			}
		case *ast.CaseClause:
			// ..., *ast.TypeSwitchStmt, *ast.BlockStmt, *ast.CaseClause
			if len(stack) < 3 {
				return true
			}
			if _, inTypeSwitch := stack[len(stack)-3].(*ast.TypeSwitchStmt); !inTypeSwitch {
				return true
			}
			for _, expr := range n.List {
				if isBytesPointer(pass, expr) {
					checkGuarded(pass, expr, []ast.Node{n})
				}
			}
		}
		return true
	})
	return nil, nil
}

// isBytesPointer reports whether expr denotes the type *[]byte.
func isBytesPointer(pass *analysis.Pass, expr ast.Expr) bool {
	t := pass.TypesInfo.TypeOf(expr)
	return t != nil && types.Identical(t, bytesPointer)
}

// guardedRegion returns the nodes that the type assertion at the top of stack
// guards: the condition and body of the if statement it initializes, or else
// the body of the innermost enclosing function.
func guardedRegion(stack []ast.Node) []ast.Node {
	// stack ends with ..., *ast.IfStmt, *ast.AssignStmt, *ast.TypeAssertExpr
	// for `if data, ok := v.(*[]byte); ok`.
	if len(stack) >= 3 {
		if ifStmt, ok := stack[len(stack)-3].(*ast.IfStmt); ok && ifStmt.Init == stack[len(stack)-2] {
			return []ast.Node{ifStmt.Cond, ifStmt.Body}
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		switch fn := stack[i].(type) {
		case *ast.FuncLit:
			return []ast.Node{fn.Body}
		case *ast.FuncDecl:
			return []ast.Node{fn.Body}
		}
	}
	return nil
}

// checkGuarded reports at node unless one of region calls CheckFunc.
func checkGuarded(pass *analysis.Pass, node ast.Node, region []ast.Node) {
	for _, n := range region {
		if callsCheckFunc(pass, n) {
			return
		}
	}
	pass.Reportf(node.Pos(), "*[]byte passthrough skips the content-type check: call %s before returning raw bytes", CheckFunc)
}

// callsCheckFunc reports whether root contains a call to CheckFunc of the package under analysis.
func callsCheckFunc(pass *analysis.Pass, root ast.Node) bool {
	found := false
	ast.Inspect(root, func(n ast.Node) bool {
		if found {
			return false
		}
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		ident, ok := ast.Unparen(call.Fun).(*ast.Ident)
		if !ok {
			return true
		}
		fn, ok := pass.TypesInfo.Uses[ident].(*types.Func)
		found = ok && fn.Pkg() == pass.Pkg && fn.Name() == CheckFunc
		return !found
	})
	return found
}
