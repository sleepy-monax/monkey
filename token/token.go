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
	Equal      = "Equal"
	NotEqual   = "NotEqual"

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
	While    = "While"
	Else     = "Else"
	Return   = "Return"
	And      = "And"
	Or       = "Or"
	Not      = "Not"
)

var Keywords = map[string]TokenType{
	"function": Function,
	"let":      Let,
	"true":     True,
	"false":    False,
	"if":       If,
	"while":    While,
	"else":     Else,
	"return":   Return,
	"and":      And,
	"or":       Or,
	"not":      Not,
}

var Operators = map[string]TokenType{
	"=":  Assign,
	"+":  Plus,
	"-":  Minus,
	"!":  Bang,
	"*":  Asterisk,
	"/":  Slash,
	"<":  LessThan,
	">":  BiggerThan,
	"==": Equal,
	"!=": NotEqual,
}

type Token struct {
	Type    TokenType
	Literal string

	Line   int
	Column int
}
