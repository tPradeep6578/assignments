package main

import (
	"fmt"

	arithmetic "github.com/tpradeep6578/assignments/go-practice/arithmetic"
	datatype "github.com/tpradeep6578/assignments/go-practice/datatypes"
	"github.com/tpradeep6578/assignments/go-practice/sample"
	scope "github.com/tpradeep6578/assignments/go-practice/scope"
	variable "github.com/tpradeep6578/assignments/go-practice/variableDeclaration"
)

func main() {
	fmt.Println("main print")
	sample.SampleFunc()
	variable.VariableFunc()
	scope.ScopeDeclaration()
	scope.ScopeTestFunc()
	datatype.DataTypeFunc()
	arithmetic.ArithmeticOperationFunc()
}
