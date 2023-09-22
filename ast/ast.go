package ast

import (
	"9ccgo/token"
	"bytes"
	"strconv"
)

type Node interface {
	node()
	String() string
}

type OperatorNode interface {
	Node
	OperatorType() token.TokenType
}

type InfixOperatorNode struct {
	Operator token.TokenType
	Lhs      Node
	Rhs      Node
}

func (i *InfixOperatorNode) node()                         {}
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
	Rhs      Node
}

func (p *PrefixOperatorNode) node()                         {}
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

func (i *IntegerNode) node()          {}
func (i *IntegerNode) String() string { return strconv.Itoa(i.Value) }
