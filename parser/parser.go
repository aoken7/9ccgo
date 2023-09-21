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

func (p *Parser) multiple() ast.Node {
	node := p.primary()

	for {
		if p.consume(token.ASTERISK) {
			node = newInfixNode(node, p.primary(), "*")
		} else if p.consume(token.SLASH) {
			node = newInfixNode(node, p.primary(), "/")
		} else {
			return node
		}
	}
}

func (p *Parser) expr() ast.Node {
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

func (p *Parser) Parse() ast.Node {
	return p.expr()
}

func newInfixNode(l, r ast.Node, oper token.TokenType) ast.Node {
	node := &ast.InfixNode{}
	node.Operator = oper
	node.Lhs = l
	node.Rhs = r
	return node
}
