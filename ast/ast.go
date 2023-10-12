package ast

import (
	"9ccgo/types"
	"bytes"
	"strings"
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
	Type            types.Type
	InitDeclarators []InitDeclarator
}

func (d *Declaration) node() {}
func (d *Declaration) String() string {
	var out bytes.Buffer

	out.WriteString(string(d.Type))
	out.WriteString(" ")

	var decs []string
	for _, s := range d.InitDeclarators {
		decs = append(decs, s.String())
	}
	out.WriteString(strings.Join(decs, ","))
	out.WriteString(";")

	return out.String()
}

type InitDeclarator struct {
	Ident IdentiferNode
	Right Expression
}

func (d *InitDeclarator) node() {}
func (d *InitDeclarator) String() string {
	var out bytes.Buffer

	out.WriteString(d.Ident.String())

	if d.Right != nil {
		out.WriteString(" = ")
		out.WriteString(d.Right.String())
	}

	return out.String()
}
