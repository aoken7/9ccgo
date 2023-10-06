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
			panic(fmt.Sprintf("undefined variable: %s", ident))
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

	return p.primary(env)
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

func (p *Parser) assign(env *Env) ast.Expression {
	node := p.equality(env)

	return node
}

func (p *Parser) expression(env *Env) ast.Expression {
	return p.assign(env)
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
		stmt = p.compoundStatement()
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
	panic(fmt.Sprintf("expected token.TYPE. but got=%T", p.curToken))
}

func (p *Parser) declarator(env *Env) *ast.Declaration {
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

func (p *Parser) directDeclarator(env *Env) *ast.Declaration {
	// <direct-declarator> ::= <identifier>
	//                       | ( <declarator> )
	//                       | <direct-declarator> [ {<constant-expression>}? ]
	//                       | <direct-declarator> ( <parameter-type-list> )
	//                       | <direct-declarator> ( {<identifier>}* )
	return &ast.Declaration{Ident: p.identifier(env)}
}

func (p *Parser) declaration(env *Env) *ast.Declaration {
	// <declaration> ::=  {<declaration-specifier>}+ {<init-declarator>}* ;
	ds := p.declarationSpecifier()
	id := p.initDeclarator(env, ds)
	p.consume(token.SEMICOLON)
	return id
}

func (p *Parser) initDeclarator(env *Env, typ types.Type) *ast.Declaration {
	// <init-declarator> ::= <declarator>
	//                 | <declarator> = <initializer>
	dec := p.declarator(env)

	p.consume(token.ASSIGN)

	dec.Type = typ
	dec.Right = p.initializer(env)

	return dec
}

func (p *Parser) initializer(env *Env) ast.Node {
	// <initializer> ::= <assignment-expression>
	//				   | { <initializer-list> }
	//                 | { <initializer-list> , }

	return p.expression(env)
}

// TODO: 親の変数にアクセスする方法を考える
func (p *Parser) compoundStatement() ast.Statement {
	env := &Env{env: map[string]int{}}
	node := &ast.CompoundStatement{}

	p.consume("{")

	for !p.consume(token.RBRACE) && !p.expect(token.EOF) {
		if p.peek() == token.TYPE {
			node.Statements = append(node.Statements, p.declaration(env))
		} else {
			node.Statements = append(node.Statements, p.stmtement(env))
		}
	}

	return node
}

func (p *Parser) Parse() ast.Node {
	return p.compoundStatement()
}

func newInfixNode(l, r ast.Expression, oper token.TokenType) ast.Expression {
	node := &ast.InfixOperatorNode{}
	node.Operator = oper
	node.Lhs = l
	node.Rhs = r
	return node
}
