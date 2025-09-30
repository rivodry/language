package parser

import (
	"lang/ast"
	"lang/lex"
)

type Parser struct {
	tokens []lex.Token
	pos    int
}

func Parse(tokens []lex.Token) ast.BinNode {
	p := &Parser{tokens: tokens}
	node := p.parseExpr()
	return *node.(*ast.BinNode)
}

func (p *Parser) current() lex.Token {
	if p.pos >= len(p.tokens) {
		return lex.Token{Typ: "EOF"}
	}
	return p.tokens[p.pos]
}

func (p *Parser) eat(expected string) lex.Token {
	tok := p.current()
	if tok.Typ == expected {
		p.pos++
		return tok
	}
	panic("Unexpected token: expected " + expected + ", got " + tok.Typ)
}

func (p *Parser) parseExpr() any {
	node := p.parseTerm()

	for p.current().Typ == "+" || p.current().Typ == "-" {
		op := p.current().Typ
		p.eat(op)
		right := p.parseTerm()
		node = &ast.BinNode{
			Op:    op,
			Left:  node,
			Right: right,
		}
	}
	return node
}

func (p *Parser) parseTerm() any {
	node := p.parseFactor()

	for p.current().Typ == "*" || p.current().Typ == "/" {
		op := p.current().Typ
		p.eat(op)
		right := p.parseFactor()
		node = &ast.BinNode{
			Op:    op,
			Left:  node,
			Right: right,
		}
	}
	return node
}

func (p *Parser) parseFactor() any {
	tok := p.current()

	switch tok.Typ {
	case "NUMBER":
		p.eat("NUMBER")
		return &ast.NumNode{Value: tok.Value}
	case "IDENT":
		p.eat("IDENT")
		return &ast.IdentNode{Ident: tok.Value}
	case "(":
		p.eat("(")
		expr := p.parseExpr()
		p.eat(")")
		return expr
	default:
		panic("Unexpected token in factor: " + tok.Typ)
	}
}
