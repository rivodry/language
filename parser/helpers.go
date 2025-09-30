package parser

import "lang/lex"

type Parser struct {
	tokens []lex.Token
	pos    int
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
