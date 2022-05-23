package complexity

import (
	"go/ast"
	"go/token"
)

func Count(node ast.Node) int {
	count := 1
	ast.Inspect(node, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.IfStmt:
			count += countCondition(n.Cond)
		case *ast.ForStmt, *ast.RangeStmt:
			count++
		case *ast.CaseClause:
			if n.List == nil {
				count++
				break
			}
			for _, stmt := range n.List {
				count += countCondition(stmt)
			}
		case *ast.CommClause:
			count++
		}
		return true
	})
	return count
}

func countCondition(condNode ast.Node) int {
	count := 1
	ast.Inspect(condNode, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.BinaryExpr:
			if n.Op == token.LAND || n.Op == token.LOR {
				count++
			}
		}
		return true
	})
	return count
}
