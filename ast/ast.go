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

type InfixNode struct {
	Operator token.TokenType
	Lhs      Node
	Rhs      Node
}

func (i *InfixNode) node() {}
func (i *InfixNode) String() string {
	var out bytes.Buffer

	out.WriteString("(" + i.Lhs.String())
	out.WriteString(" " + string(i.Operator) + " ")
	out.WriteString(i.Rhs.String() + ")")

	return out.String()
}

type IntegerNode struct {
	Value int
}

func (i *IntegerNode) node()          {}
func (i *IntegerNode) String() string { return strconv.Itoa(i.Value) }
