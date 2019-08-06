package token

// TokenType is the type of the token
type TokenType string

const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"

	// Identifier and literals
	Identifier = "IDENTIFIER"
	Integer    = "NUMBER"

	// Operators
	Assign = "EQUAL"
	Plus   = "PLUS"

	// Delemiters
	Comma              = "COMMA"
	Semicolon          = "SEMICOLON"
	OpeningParenthesis = "OPENING_PARENTHESIS"
	ClosingParenthesis = "CLOSING_PARENTHESIS"
	OpeningBrace       = "OPENING_BRACE"
	ClosingBrace       = "CLOSING_BRACE"
	OpeningBracket     = "OPENING_BRACKET"
	ClosingBracket     = "CLOSING_BRACKET"

	// Keywords
	Function = "FUNCTION"
	Let      = "LET"
)

var Keywords = map[string]TokenType{
	"fn":  Function,
	"let": Let,
}

type Token struct {
	Type    TokenType
	Literal string

	Line   int
	Column int
}
