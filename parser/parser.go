package parser

import (
	"fmt"
	"lang/lex"
)

type Parser struct {
	Tokens []lex.Token
	I      int
}

func (P Parser) IsAtEnd() bool {
	return P.Peek().Typ == "EOF"
}
func (P Parser) Peek() lex.Token {
	return P.Tokens[P.I]
}
func (P *Parser) Advance() {
	if !P.IsAtEnd() {
		P.I++
	}
}
func Parse(L []lex.Token) {
	P := Parser{L, 0}
	for !P.IsAtEnd() {
		fmt.Println(P.Peek())
		P.Advance()
	}
}
