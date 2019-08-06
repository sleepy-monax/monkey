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
	Minus      = "Minus"
	Bang       = "Bang"
	Asterisk   = "Asterisk"
	Slash      = "Slash"
	LessThan   = "LessThan"
	BiggerThan = "BiggerThan"

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
	True     = "True"
	False    = "False"
	If       = "If"
	Else     = "Else"
	Return   = "Return"
)

var Keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

var Operators = map[string]TokenType{
	"=": Assign,
	"+": Plus,
	"-": Minus,
	"!": Bang,
	"*": Asterisk,
	"/": Slash,
	"<": LessThan,
	">": BiggerThan,
}

type Token struct {
	Type    TokenType
	Literal string

	Line   int
	Column int
}
