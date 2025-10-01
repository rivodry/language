package parser

import "lang/ast"

func (p *Parser) parseStatement() any {
	return p.parseExpressionStatement()
}
func (p *Parser) parseExpressionStatement() any {
	expr := p.parseExpr()
	p.eat("TERM")
	return ast.ExpressionStatement{
		Expr: expr,
	}
}
