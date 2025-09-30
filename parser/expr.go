package parser

import "lang/ast"

func (p *Parser) parseExpr() any {
	return p.parseTerm()
}

func (p *Parser) parseTerm() any {
	node := p.parseFactor()

	for p.current().Typ == "+" || p.current().Typ == "-" {
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
	node := p.parseOperand()

	for p.current().Typ == "*" || p.current().Typ == "/" {
		op := p.current().Typ
		p.eat(op)
		right := p.parseOperand()
		node = &ast.BinNode{
			Op:    op,
			Left:  node,
			Right: right,
		}
	}
	return node
}
func (p *Parser) parseOperand() any {
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
