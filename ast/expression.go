package ast

import (
	"9ccgo/token"
	"bytes"
	"strconv"
)

type OperatorNode interface {
	Expression
	OperatorType() token.TokenType
}

type InfixOperatorNode struct {
	Operator token.TokenType
	Lhs      Expression
	Rhs      Expression
}

func (i *InfixOperatorNode) node()                         {}
func (i *InfixOperatorNode) expressionNode()               {}
func (i *InfixOperatorNode) OperatorType() token.TokenType { return i.Operator }
func (i *InfixOperatorNode) String() string {
	var out bytes.Buffer

	out.WriteString("(" + i.Lhs.String())
	out.WriteString(" " + string(i.Operator) + " ")
	out.WriteString(i.Rhs.String() + ")")

	return out.String()
}

type PrefixOperatorNode struct {
	Operator token.TokenType
	Rhs      Expression
}

func (p *PrefixOperatorNode) node()                         {}
func (p *PrefixOperatorNode) expressionNode()               {}
func (p *PrefixOperatorNode) OperatorType() token.TokenType { return p.Operator }
func (p PrefixOperatorNode) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(string(p.Operator))
	out.WriteString(p.Rhs.String())
	out.WriteString(")")

	return out.String()
}

type IntegerNode struct {
	Value int
}

func (i *IntegerNode) node()           {}
func (i *IntegerNode) expressionNode() {}
func (i *IntegerNode) String() string  { return strconv.Itoa(i.Value) }

type IdentiferNode struct {
	Identifer string
	Offset    int
}

func (i *IdentiferNode) node()           {}
func (i *IdentiferNode) expressionNode() {}
func (i *IdentiferNode) String() string  { return i.Identifer }

type AssignmentNode struct {
	Ident Expression
	Right Expression
}

func (i *AssignmentNode) node()           {}
func (i *AssignmentNode) expressionNode() {}
func (i *AssignmentNode) String() string {
	var out bytes.Buffer

	out.WriteString(i.Ident.String())
	out.WriteString(" = ")
	out.WriteString(i.Right.String())

	return out.String()
}

type FunctionCallNode struct {
	Idetifer IdentiferNode
}

func (f *FunctionCallNode) node()           {}
func (f *FunctionCallNode) expressionNode() {}
func (f *FunctionCallNode) String() string {
	var out bytes.Buffer

	out.WriteString(f.Idetifer.Identifer)

	return out.String()
}
