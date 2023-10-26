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

type FunctionNode struct {
	Type         types.Type
	Ident        IdentiferNode
	Declarations []Declaration
	CmpStmt      CompoundStatement
}

func (f *FunctionNode) node() {}
func (f *FunctionNode) String() string {
	var out bytes.Buffer

	out.WriteString(string(f.Type) + " ")
	out.WriteString(f.Ident.String())
	out.WriteString("(")
	var decs []string
	for _, dec := range f.Declarations {
		decs = append(decs, dec.String())
	}
	out.WriteString(strings.Join(decs, ", "))
	out.WriteString(")")
	out.WriteString("{")
	out.WriteString(f.CmpStmt.String())
	out.WriteString("}")

	return out.String()
}

type RootNode struct {
	Units []Node
}

func (r *RootNode) node() {}
func (r *RootNode) String() string {
	var outs []string
	for _, unit := range r.Units {
		outs = append(outs, unit.String())
	}

	return strings.Join(outs, "\n")
}
