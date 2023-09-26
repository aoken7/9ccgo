package ast

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
