package lexer

import "github.com/gmr458/toylang/token"

type Lexer struct {
	// Input source code.
	input string

	// Current position in input, it points to character
	// in the input that corresponds to the ch byte.
	position int

	// Current reading position in input, always points to
	// the next character in the input.
	readPosition int

	// Current character under examination.
	ch byte
}

// New returns a new initialized Lexer.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) isEndOfInput() bool {
	return l.readPosition >= len(l.input)
}

// peekChar is really similar to readChar, except that it doesn’t increment l.position
// and l.readPosition. It returns the character to which readPosition is pointing.
func (l *Lexer) peekChar() byte {
	if l.isEndOfInput() {
		return 0
	}
	return l.input[l.readPosition]
}

// readChar sets the index of the next character (readPosition)
// and advance the position.
func (l *Lexer) readChar() {
	if l.isEndOfInput() {
		l.ch = 0 // sets ch to 0 (ASCII code for the NUL character)
	} else {
		l.ch = l.input[l.readPosition] // sets ch to the next character
	}

	l.position = l.readPosition // update position according to ch
	l.readPosition += 1         // update readPosition so always point to the next character
}

// readIdentifier reads in an identifier and advances the lexer’s
// positions until it encounters a non-letter-character.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads in a number and advances the lexer’s
// positions until it encounters a non-digit-character.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// New returns a new initialized Token.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// NextToken looks at the current character under examination (l.ch)
// and returns a token depending on which character it is and before
// returning the token it advance the pointers into the input, so
// when this method is called again the l.ch field is already updated.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // returns here bacause readIdentifier() already calls readChar()
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok // returns here bacause readNumber() already calls readChar()
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
