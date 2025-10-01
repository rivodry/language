package parser

import (
	"lang/lex"
)

func Parse(tokens []lex.Token) []any {
	p := &Parser{tokens: tokens}
	statements := []any{}
	for !(p.current().Typ == "EOF") {
		statements = append(statements, p.parseDeclaration())
	}
	return statements
}
