package parser

import "lang/ast"

func (p *Parser) parseDeclaration() any {

	switch p.current().Typ {
	case "VAR":
		{
			p.eat("VAR")
			return p.parseVar()
		}
	case "SET":
		{
			p.eat("SET")
			return p.parseSet()
		}
	default:
		{
			return p.parseStatement()
		}
	}

}
func (p *Parser) parseSet() any {
	name := p.current().Value
	p.eat("IDENT")
	p.eat("=")
	set := p.parseExpr()
	p.eat("TERM")
	return ast.SetNode{
		Name: name,
		Set:  set,
	}
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
