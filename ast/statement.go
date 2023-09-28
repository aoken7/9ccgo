package ast

import (
	"bytes"
	"strings"
)

// 関数のブロック部分に該当
type CompoundStatement struct {
	// <compound-statement> ::= { {<declaration>}* {<statement>}* }
	Statements []Statement
}

func (cs *CompoundStatement) node()          {}
func (cs *CompoundStatement) statementNode() {}
func (cs *CompoundStatement) String() string {
	var outs []string

	for _, stmt := range cs.Statements {
		outs = append(outs, stmt.String())
	}

	return strings.Join(outs, ", ")
}

type ExpressionStatement struct {
	Expression Expression
}

func (es *ExpressionStatement) node()          {}
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) String() string { return es.Expression.String() }

type ReturnStatement struct {
	Expression Expression
}

func (rs *ReturnStatement) node()          {}
func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) String() string {
	return "return " + rs.Expression.String() + ";"
}

type IfStatement struct {
	Expression     Expression
	TrueStatement  Statement
	FalseStatement Statement
}

func (is *IfStatement) node()          {}
func (is *IfStatement) statementNode() {}
func (is *IfStatement) String() string {
	var out bytes.Buffer

	out.WriteString("if(")
	out.WriteString(is.Expression.String())
	out.WriteString("){ ")
	out.WriteString(is.TrueStatement.String())
	out.WriteString(" }")

	if is.FalseStatement != nil {
		out.WriteString("else{")
		out.WriteString(is.FalseStatement.String())
		out.WriteString("}")
	}

	return out.String()
}
