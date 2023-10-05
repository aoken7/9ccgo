package ast

import (
	"bytes"
)

type Node interface {
	node()
	String() string
}

type Statement interface {
	/* <statement> ::= <labeled-statement>
	   | <expression-statement>
	   | <compound-statement>
	   | <selection-statement>
	   | <iteration-statement>
	   | <jump-statement>
	*/
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Declaration struct {
	Type  string
	Ident IdentiferNode
	right Node
}

func (d *Declaration) node() {}
func (d *Declaration) String() string {
	var out bytes.Buffer

	out.WriteString(string(d.Type) + " ")
	out.WriteString(d.Ident.String() + " ")

	if d.right != nil {
		out.WriteString("=")
		out.WriteString(d.right.String())
	}

	out.WriteString(";")
	return out.String()
}
