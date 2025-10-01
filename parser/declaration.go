package parser

func (p *Parser) parseDeclaration() any {
	if p.current().Typ == "VAR" {
		return p.parseVar()
	}
	return p.parseStatement()
}
func (p *Parser) parseVar() any {

}
