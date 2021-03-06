package complexity

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func GetFuncNode(t *testing.T, code string) ast.Node {
	t.Helper()

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range file.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok {
			return fd
		}
	}
	t.Fatal("no function declear found")
	return nil
}

func TestComplexity(t *testing.T) {
	testcases := []struct {
		name       string
		code       string
		complexity int
	}{
		{
			name: "simple",
			code: `package main
func Double(n int) int {
	return n * 2
}`,
			complexity: 1,
		},
		{
			name: "if statement",
			code: `package main
func Double(n int) int {
	if n%2 == 0 {
		return 0
	}
	return n
}`,
			complexity: 2,
		},
		{
			name: "for statement",
			code: `package main
func Sum(n int) int {
	c := 0
	for i := 0; i < n; i++ {
		c += i
	}
	return c
}`,
			complexity: 2,
		},
		{
			name: "case clause simple",
			code: `package main
func Print(n int) {
	switch n {
	case 0:
		println("zero")
	}
}`,
			complexity: 2,
		},
		{
			name: "case clause default",
			code: `package main
func Print(n int) {
	switch n {
	case 0:
		println("zero")
	default:
		println("not zero")
	}
}`,
			complexity: 3,
		},
		{
			name: "case clause multi case",
			code: `package main
func Print(n int) {
	switch n {
	case 0, 1, 2, 3:
		println("between zero and three")
	}
}`,
			complexity: 5,
		},
		{
			name: "condition and",
			code: `package main
func Print(n int) {
	if 0 < n && n < 10 {
		println("between zero and ten")
	}
}`,
			complexity: 3,
		},
		{
			name: "condition or",
			code: `package main
func Print(n int) {
	if n < 10 || 10 < n {
		println("not between zero and ten")
	}
}`,
			complexity: 3,
		},
		{
			name: "range statement",
			code: `package main
func Print(n int) {
	l := []int{n}
	for i := range l {
		println(i)
	}
}`,
			complexity: 2,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			a := GetFuncNode(t, testcase.code)
			c := Count(a)

			if c != testcase.complexity {
				t.Errorf("got=%d, want=%d", c, testcase.complexity)
			}
		})
	}
}
