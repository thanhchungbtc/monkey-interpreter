package parser

import "monkey-interpreter/ast"

func (p *Parser) ParseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Left:     left,
		Operator: p.curToken.Literal,
		Right:    nil,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

