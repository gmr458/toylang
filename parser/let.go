package parser

import (
	"github.com/gmr458/toylang/ast"
	"github.com/gmr458/toylang/token"
)

// parseLetStatement constructs an *ast.LetStatement node with the
// token it’s currently sitting on (a token.LET token) and then
// advances the tokens while making assertions about the next token
// with calls to expectPeek. First it expects a token.IDENT token,
// which it then uses to construct an *ast.Identifier node. Then it
// expects an equal sign and finally it jumps over the expression
// following the equal sign until it encounters a semicolon.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
