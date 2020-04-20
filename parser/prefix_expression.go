package parser

import "monkey-interpreter/ast"

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Right:    nil,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}
