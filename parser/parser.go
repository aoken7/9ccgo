package parser

import (
	"9ccgo/ast"
	"9ccgo/token"
	"fmt"
	"strconv"
)

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
	if p.tokens[p.readPosition].Type == token.EOF {
		return
	}

	p.curToken = p.tokens[p.readPosition]
	p.position = p.readPosition
	p.readPosition++
}

func (p *Parser) peek() token.TokenType {
	return p.tokens[p.readPosition].Type
}

func (p *Parser) consume(t token.TokenType) bool {
	if p.curToken.Type == t {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) primary() ast.Node {
	if p.consume(token.LPAREN) {
		node := p.expr()
		if !p.consume(token.RPAREN) {
			panic(fmt.Sprintf("expected token is ')'. got %v", p.curToken))
		}
		return node
	} else if p.peek() == token.ASSIGN {
		ident := p.curToken.Literal
		p.nextToken()

		return &ast.IdentiferNode{
			Identifer: ident,
			Offset:    int(ident[0] - 'a'),
		}
	}

	// 多分Int
	num, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	p.nextToken()
	return &ast.IntegerNode{Value: num}
}

func (p *Parser) unary() ast.Node {
	if p.consume("-") {
		return &ast.PrefixOperatorNode{Operator: "-", Rhs: p.primary()}
	}

	return p.primary()
}

func (p *Parser) multiple() ast.Node {
	node := p.unary()

	for {
		if p.consume(token.ASTERISK) {
			node = newInfixNode(node, p.unary(), "*")
		} else if p.consume(token.SLASH) {
			node = newInfixNode(node, p.unary(), "/")
		} else {
			return node
		}
	}
}

func (p *Parser) add() ast.Node {
	node := p.multiple()

	for {
		if p.consume(token.PLUS) {
			node = newInfixNode(node, p.multiple(), "+")
		} else if p.consume(token.MINUS) {
			node = newInfixNode(node, p.multiple(), "-")
		} else {
			return node
		}
	}
}

func (p *Parser) relational() ast.Node {
	node := p.add()

	for {
		if p.consume(token.LSS) {
			node = newInfixNode(node, p.add(), "<")
		} else if p.consume(token.LEQ) {
			node = newInfixNode(node, p.add(), "<=")
		} else if p.consume(token.GTR) {
			node = newInfixNode(node, p.add(), ">")
		} else if p.consume(token.GEQ) {
			node = newInfixNode(node, p.add(), ">=")
		} else {
			return node
		}
	}
}

func (p *Parser) equality() ast.Node {
	node := p.relational()

	for {
		if p.consume(token.EQL) {
			node = newInfixNode(node, p.relational(), "==")
		} else if p.consume(token.NEQ) {
			node = newInfixNode(node, p.relational(), "!=")
		} else {
			return node
		}
	}
}

func (p *Parser) assign() ast.Node {
	node := p.equality()

	if p.consume(token.ASSIGN) {
		node = newInfixNode(node, p.assign(), "=")
	}

	return node
}

func (p *Parser) expr() ast.Node {
	return p.assign()
}

func (p *Parser) stmt() ast.Node {
	node := p.expr()
	// TODO:consumeでは無くexpectを使う
	p.consume(token.SEMICOLON)
	return node
}

func (p *Parser) Parse() ast.Node {
	return p.stmt()
}

func newInfixNode(l, r ast.Node, oper token.TokenType) ast.Node {
	node := &ast.InfixOperatorNode{}
	node.Operator = oper
	node.Lhs = l
	node.Rhs = r
	return node
}
