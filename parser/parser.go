package parser

import (
	"fmt"
	"strconv"

	"github.com/gmr458/toylang/ast"
	"github.com/gmr458/toylang/lexer"
	"github.com/gmr458/toylang/token"
)

const (
	_ int = iota
	LOWEST

	// ==
	EQUALS

	// > or <
	LESSGREATER

	// +
	SUM

	// *
	PRODUCT

	// -X or !X
	PREFIX

	// myFunction(x)
	CALL
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	// Pointer to an instance of lexer.Lexer to repeatedly
	// call NextToken() to get then next token in the input.
	l *lexer.Lexer

	// Current token under examination
	curToken token.Token

	// Next token
	peekToken token.Token

	errors []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

// New returns a new initialized Parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// registerPrefix adds a key value pair of token.TokenType and a
// prefixParseFn function to the map prefixParseFns
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerInfix adds a key value pair of token.TokenType and a
// infixParseFn function to the map infixParseFns
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// curTokenIs reports whether a given token.TokenType is equal
// to the current token.
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs reports whether a given token.TokenType is equal
// to the peek token.
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek checks the type of the peekToken and only if the
// type is correct does it advance the tokens by calling nextToken.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if nil == prefix {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// ParseProgram construct the root node of the AST, an *ast.Program.
// It then iterates over every token in the input until it encounters
// an token.EOF token. It does this by repeatedly calling nextToken,
// which advances both p.curToken and p.peekToken. In every iteration
// it calls parseStatement, whose job it is to parse a statement. If
// parseStatement returned something other than nil, a ast.Statement,
// its return value is added to Statements slice of the AST root node.
// When nothing is left to parse the *ast.Program root node is returned.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}
