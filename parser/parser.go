package parser

import (
	"9ccgo/ast"
	"9ccgo/token"
	"9ccgo/types"
	"fmt"
	"strconv"
)

type Env struct {
	env    map[string]int
	offset int
}

type Parser struct {
	tokens       []token.Token
	curToken     token.Token
	position     int
	readPosition int
}

func New(t []token.Token) *Parser {
	p := &Parser{tokens: t}
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	if p.expect(token.EOF) {
		return
	}

	p.curToken = p.tokens[p.readPosition]
	p.position = p.readPosition
	p.readPosition++
}

func (p *Parser) peek() token.TokenType {
	return p.curToken.Type
}

func (p *Parser) consume(t token.TokenType) bool {
	if p.expect(t) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) expect(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) translationUnit() ast.Node {
	// <translation-unit> ::= {<external-declaration>}*
	env := &Env{env: map[string]int{}}

	rootNode := &ast.RootNode{}
	for !p.expect(token.EOF) {
		rootNode.Units = append(rootNode.Units, p.externalDeclaration(*env))
	}
	return rootNode
}

func (p *Parser) externalDeclaration(env Env) ast.Node {
	// <external-declaration> ::= <function-definition>
	//                         | <declaration>

	return p.functionDefinition(env)
}

func (p *Parser) functionDefinition(env Env) ast.Node {
	// <function-definition> ::= {<declaration-specifier>}* <declarator> {<declaration>}* <compound-statement>
	// TODO: グローバル変数に対応

	function := &ast.FunctionNode{}
	function.Type = p.declarationSpecifier()
	function.Ident = p.declarator(&env)
	p.consume("(")
	for p.curToken.Type != token.RPAREN {
		function.Declarations = append(function.Declarations, *p.declaration(&env))
	}
	p.consume(")")
	p.consume("{")
	function.CmpStmt = *p.compoundStatement(env)
	p.consume("}")
	return function
}

func (p *Parser) postfixExpression(env *Env) ast.Expression {
	//<postfix-expression> ::= <primary-expression>
	//                   | <postfix-expression> [ <expression> ]
	//                   | <postfix-expression> ( {<assignment-expression>}* )
	//                   | <postfix-expression> . <identifier>
	//                   | <postfix-expression> -> <identifier>
	//                   | <postfix-expression> ++
	//                   | <postfix-expression> --

	exp := p.primary(env)
	if p.consume("(") {
		ident, ok := exp.(*ast.IdentiferNode)
		if !ok {
			fmt.Printf("expected *ast.IdentiferNode. got %v", ident)
		}

		funcCall := &ast.FunctionCallNode{Idetifer: *ident}

		for !p.expect(")") {
			funcCall.Args = append(funcCall.Args, p.assignment_expression(env))
			p.consume(",")
		}

		p.consume(")")

		return funcCall
	}

	return exp
}

func (p *Parser) primary(env *Env) ast.Expression {
	if p.consume(token.LPAREN) {
		node := p.expression(env)
		if !p.consume(token.RPAREN) {
			panic(fmt.Sprintf("expected token is ')'. got %v", p.curToken))
		}
		return node
	} else if p.expect(token.IDENT) {
		// p.comsume() するとidentが取れないのでcurTokenで判定
		ident := p.curToken.Literal
		p.nextToken()

		offset, ok := env.env[ident]
		if !ok {
			//panic(fmt.Sprintf("undefined variable: %s", ident))
		}

		return &ast.IdentiferNode{
			Identifer: ident,
			Offset:    offset,
		}
	}

	// 多分Int
	num, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		panic(fmt.Sprintf("expected token is number. got %v", p.curToken))
	}

	p.nextToken()
	return &ast.IntegerNode{Value: num}
}

func (p *Parser) unary(env *Env) ast.Expression {
	if p.consume("-") {
		return &ast.PrefixOperatorNode{Operator: "-", Rhs: p.primary(env)}
	}

	return p.postfixExpression(env)
}

func (p *Parser) multiple(env *Env) ast.Expression {
	node := p.unary(env)

	for {
		if p.consume(token.ASTERISK) {
			node = newInfixNode(node, p.unary(env), "*")
		} else if p.consume(token.SLASH) {
			node = newInfixNode(node, p.unary(env), "/")
		} else {
			return node
		}
	}
}

func (p *Parser) add(env *Env) ast.Expression {
	node := p.multiple(env)

	for {
		if p.consume(token.PLUS) {
			node = newInfixNode(node, p.multiple(env), "+")
		} else if p.consume(token.MINUS) {
			node = newInfixNode(node, p.multiple(env), "-")
		} else {
			return node
		}
	}
}

func (p *Parser) relational(env *Env) ast.Expression {
	node := p.add(env)

	for {
		if p.consume(token.LSS) {
			node = newInfixNode(node, p.add(env), "<")
		} else if p.consume(token.LEQ) {
			node = newInfixNode(node, p.add(env), "<=")
		} else if p.consume(token.GTR) {
			node = newInfixNode(node, p.add(env), ">")
		} else if p.consume(token.GEQ) {
			node = newInfixNode(node, p.add(env), ">=")
		} else {
			return node
		}
	}
}

func (p *Parser) equality(env *Env) ast.Expression {
	node := p.relational(env)

	for {
		if p.consume(token.EQL) {
			node = newInfixNode(node, p.relational(env), "==")
		} else if p.consume(token.NEQ) {
			node = newInfixNode(node, p.relational(env), "!=")
		} else {
			return node
		}
	}
}

func (p *Parser) assignment_expression(env *Env) ast.Expression {
	// <assignment-expression> ::= <conditional-expression>
	//                           | <unary-expression> <assignment-operator> <assignment-expression>
	node := p.equality(env)
	if p.consume(token.ASSIGN) {
		return &ast.AssignmentNode{
			Ident: node,
			Right: p.assignment_expression(env),
		}
	}
	return node
}

func (p *Parser) expression(env *Env) ast.Expression {
	return p.assignment_expression(env)
}

func (p *Parser) expressionStatement(env *Env) ast.Statement {
	node := p.expression(env)

	es := &ast.ExpressionStatement{}
	es.Expression = node
	return es
}

func (p *Parser) jumpStatement(env *Env) ast.Statement {
	if p.consume(token.RETURN) {
		exp := p.expression(env)
		return &ast.ReturnStatement{Expression: exp}
	}
	panic(fmt.Sprintf("parser err: expected return. but got=%T\n", p.peek()))
}

func (p *Parser) selectionStatement(env *Env) ast.Statement {
	if p.consume(token.IF) {
		selectionStmt := &ast.IfStatement{}
		p.consume("(")
		selectionStmt.Expression = p.expression(env)
		p.consume(")")
		selectionStmt.TrueStatement = p.stmtement(env)

		if p.consume(token.ELSE) {
			selectionStmt.FalseStatement = p.stmtement(env)
		}

		return selectionStmt
	}

	return nil
}

func (p *Parser) stmtement(env *Env) ast.Statement {
	var stmt ast.Statement
	switch p.curToken.Type {
	case token.RETURN:
		stmt = p.jumpStatement(env)
	case token.IF:
		stmt = p.selectionStatement(env)
	case token.LBRACE:
		stmt = p.compoundStatement(*env)
	default:
		stmt = p.expressionStatement(env)
	}
	// TODO:consumeでは無くexpectを使う
	p.consume(token.SEMICOLON)
	return stmt
}

func (p *Parser) declarationSpecifier() types.Type {
	//<declaration-specifier> ::= <storage-class-specifier>
	//                          | <type-specifier>
	//                          | <type-qualifier>
	if p.consume(token.TYPE) {
		return types.Int
	}
	panic(fmt.Sprintf("expected token.TYPE. but got=%T", p.curToken.Type))
}

func (p *Parser) declarator(env *Env) ast.IdentiferNode {
	// <declarator> ::= {<pointer>}? <direct-declarator>
	return p.directDeclarator(env)
}

func (p *Parser) identifier(env *Env) ast.IdentiferNode {
	if p.expect(token.IDENT) {
		ident := p.curToken.Literal
		p.nextToken()

		offset, ok := env.env[ident]
		if !ok {
			env.env[ident] = env.offset
			offset = env.offset
			//env.offset += 8
		}
		env.offset += 8

		return ast.IdentiferNode{
			Identifer: ident,
			Offset:    offset,
		}
	} else {
		panic(fmt.Sprintf("expected token.IDENT. but got=%T", p.curToken))
	}
}

func (p *Parser) directDeclarator(env *Env) ast.IdentiferNode {
	// <direct-declarator> ::= <identifier>
	//                       | ( <declarator> )
	//                       | <direct-declarator> [ {<constant-expression>}? ]
	//                       | <direct-declarator> ( <parameter-type-list> )
	//                       | <direct-declarator> ( {<identifier>}* )
	return p.identifier(env)
}

func (p *Parser) declaration(env *Env) *ast.Declaration {
	// <declaration> ::=  {<declaration-specifier>}+ {<init-declarator>}* ;
	declaration := &ast.Declaration{
		Type: p.declarationSpecifier(),
	}

	for p.peek() == token.IDENT {
		declaration.InitDeclarators =
			append(declaration.InitDeclarators, *p.initDeclarator(env))
		p.consume(",")
	}

	p.consume(token.SEMICOLON)

	return declaration
}

func (p *Parser) initDeclarator(env *Env) *ast.InitDeclarator {
	// <init-declarator> ::= <declarator>
	//                 | <declarator> = <initializer>

	initDec := &ast.InitDeclarator{
		Ident: p.declarator(env),
	}

	if p.consume(token.ASSIGN) {
		initDec.Right = p.initializer(env)
	}

	return initDec
}

func (p *Parser) initializer(env *Env) ast.Expression {
	// <initializer> ::= <assignment-expression>
	//				   | { <initializer-list> }
	//                 | { <initializer-list> , }

	return p.expression(env)
}

// TODO: 親の変数にアクセスする方法を考える
func (p *Parser) compoundStatement(env Env) *ast.CompoundStatement {
	// <compound-statement> ::= { {<declaration>}* {<statement>}* }
	//env := &Env{env: map[string]int{}}
	node := &ast.CompoundStatement{}

	p.consume("{")

	// 代入を担うのはdeclaration
	for !p.consume(token.RBRACE) && !p.expect(token.EOF) {
		if p.curToken.Type == token.TYPE {
			node.Statements = append(node.Statements, p.declaration(&env))
		} else {
			node.Statements = append(node.Statements, p.stmtement(&env))
		}
	}

	p.consume("}")

	return node
}

func (p *Parser) Parse() ast.Node {
	return p.translationUnit()
}

func newInfixNode(l, r ast.Expression, oper token.TokenType) ast.Expression {
	node := &ast.InfixOperatorNode{}
	node.Operator = oper
	node.Lhs = l
	node.Rhs = r
	return node
}
