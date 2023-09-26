package ast

import (
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
