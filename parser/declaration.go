package parser

import "lang/ast"

func (p *Parser) parseDeclaration() any {
	if p.current().Typ == "VAR" {
		p.eat("VAR")
		return p.parseVar()
	}
	return p.parseStatement()
}
func (p *Parser) parseVar() any {
	name := p.current().Value
	p.eat("IDENT")
	p.eat("=")
	init := p.parseExpr()
	p.eat("TERM")
	return ast.VarNode{
		Name: name,
		Init: init,
	}
}
