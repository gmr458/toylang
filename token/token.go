package token

type TokenType string

type Token struct {
	// Type of the token
	Type TokenType

	// Holds the literal value of the token.
	Literal string
}

const (
	// Token type for illegal token/character.
	ILLEGAL = "ILLEGAL"

	// Token type for end of file.
	EOF = "EOF"

	// Token type for identifier.
	IDENT = "IDENT"

	// Token type for integer.
	INT = "INT"

	// Token type for assignation operator.
	ASSIGN = "="

	// Token type for plus operator.
	PLUS = "+"

	// Token type for minus operator.
	MINUS = "-"

	// Token type for bang operator.
	BANG = "!"

	// Token type for asterisk operator.
	ASTERISK = "*"

	// Token type for slash operator.
	SLASH = "/"

	// Token type for less than operator.
	LT = "<"

	// Token type for greater than operator.
	GT = ">"

	// Token type for equal operator.
	EQ = "=="

	// Token type for not equal operator.
	NOT_EQ = "!="

	// Token type for comma delimiter.
	COMMA = ","

	// Token type for semicolon delimiter.
	SEMICOLON = ";"

	// Token type for left paren delimiter.
	LPAREN = "("

	// Token type for right paren delimiter.
	RPAREN = ")"

	// Token type for left brace delimiter.
	LBRACE = "{"

	// Token type for right brace delimiter.
	RBRACE = "}"

	// Token type for function keyword.
	FUNCTION = "FUNCTION"

	// Token type for let keyword.
	LET = "LET"

	// Token type for true keyword.
	TRUE = "TRUE"

	// Token type for false keyword.
	FALSE = "FALSE"

	// Token type for if keyword.
	IF = "IF"

	// Token type for else keyword.
	ELSE = "ELSE"

	// Token type for return keyword.
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks the keywords table to see whether the given identifier
// is in fact a keyword. If it is, it returns the keyword’s TokenType constant.
// If it isn’t, we just get back token.IDENT, which is the TokenType for all
// user-defined identifiers.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
