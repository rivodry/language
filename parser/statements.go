package parser

import "lang/ast"

func (p *Parser) parseStatement() any {
	switch p.current().Typ {
	case "IF":
		{

			return p.parseIfStatement()
		}

	default:
		{
			return p.parseExpressionStatement()
		}
	}

}
func (p *Parser) parseExpressionStatement() any {
	expr := p.parseExpr()
	p.eat("TERM")
	return ast.ExpressionStatement{
		Expr: expr,
	}
}
func (p *Parser) parseIfStatement() any {
	p.eat("IF")
	expr := p.parseExpr()
	p.eat("THEN")

	a := []any{}
	for p.current().Typ != "END" && p.current().Typ != "EOF" {

		a = append(a, p.parseDeclaration())

	}
	p.eat("END")

	return ast.IfNode{
		If:   expr,
		Then: a,
	}
}
