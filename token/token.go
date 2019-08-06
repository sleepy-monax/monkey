package token

// TokenType is the type of the token
type TokenType string

const (
	Illegal = "Illegal"
	EOF     = "EOF"

	// Identifier and literals
	Identifier = "Identifier"
	Integer    = "Integer"

	// Operators
	Assign     = "Assign"
	Plus       = "Plus"

	// Delemiters
	Comma              = "Comma"
	Semicolon          = "Semicolon"
	OpeningParenthesis = "OpeningParenthesis"
	ClosingParenthesis = "ClosingParenthesis"
	OpeningBrace       = "OpeningBrace"
	ClosingBrace       = "ClosingBrace"
	OpeningBracket     = "OpeningBracket"
	ClosingBracket     = "ClosingBracket"

	// Keywords
	Function = "Function"
	Let      = "Let"
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
