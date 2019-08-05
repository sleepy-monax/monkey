package monkey

type TokenType string

const (
	TOKEN_ILLEGAL     = "ILLEGAL"
	TOKEN_END_OF_FILE = "END_OF_FILE"

	// Identifier and literals
	TOKEN_IDENTIFIER = "IDENTIFIER"
	TOKEN_NUMBER     = "NUMBER"

	// Operators
	TOKEN_EQUAL = "="
	TOKEN_PLUS  = "+"

	// Delemiters
	TOKEN_COMMA               = ","
	TOKEN_SEMICOLON           = ";"
	TOKEN_OPENING_PARENTHESIS = "("
	TOKEN_CLOSING_PARENTHESIS = ")"
	TOKEN_OPENING_BRACE       = "{"
	TOKEN_CLOSING_BRACE       = "}"
	TOKEN_OPENING_BRACKET     = "["
	TOKEN_CLOSING_BRACKET     = "]"

	// Keyword
	TOKEN_FUNCTION = "FUNCTION"
	TOKEN_LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string

	Line   int
	Column int
}

// TokenizerState is the current state of the tokenizer
type TokenizerState struct {
	// FIXME: maybe this should not be a string
	Input        string
	Position     int
	ReadPosition int
	CurrentChar  byte

	CurrentLine   int
	CurrentColumn int
}

// NewTokenizerState create a new tokenizer state from an input string.
func NewTokenizerState(input string) *TokenizerState {
	state := &TokenizerState{
		Input:       input,
		CurrentLine: 1,
		CurrentChar: 1,
	}

	state.NextChar()

	return state
}

// NewTokenChar create a new token from a char
func (state *TokenizerState) NewTokenChar(tokenType TokenType, char byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

// NewTokenString create a new token from a string
func (state *TokenizerState) NewTokenString(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
		Column:  state.CurrentColumn,
		Line:    state.CurrentLine,
	}
}

// NextChar get the next char form the input string.
func (state *TokenizerState) NextChar() {
	if state.ReadPosition >= len(state.Input) {
		state.CurrentChar = 0
	} else {
		state.CurrentChar = state.Input[state.ReadPosition]
	}

	if state.CurrentChar == '\n' {
		state.CurrentLine++
		state.CurrentColumn = 0
	} else {
		state.CurrentColumn++
	}

	state.Position = state.ReadPosition
	state.ReadPosition++
}

// NextToken get the next token a currentPosition in the input string
func (state *TokenizerState) NextToken() Token {
	var token Token

	switch state.CurrentChar {
	case '=':
		token = state.NewTokenChar(TOKEN_EQUAL, state.CurrentChar)
	case '+':
		token = state.NewTokenChar(TOKEN_PLUS, state.CurrentChar)
	case ',':
		token = state.NewTokenChar(TOKEN_COMMA, state.CurrentChar)
	case ';':
		token = state.NewTokenChar(TOKEN_SEMICOLON, state.CurrentChar)
	case '(':
		token = state.NewTokenChar(TOKEN_OPENING_PARENTHESIS, state.CurrentChar)
	case ')':
		token = state.NewTokenChar(TOKEN_CLOSING_PARENTHESIS, state.CurrentChar)
	case '{':
		token = state.NewTokenChar(TOKEN_OPENING_BRACE, state.CurrentChar)
	case '}':
		token = state.NewTokenChar(TOKEN_CLOSING_BRACE, state.CurrentChar)
	case '[':
		token = state.NewTokenChar(TOKEN_OPENING_BRACKET, state.CurrentChar)
	case ']':
		token = state.NewTokenChar(TOKEN_CLOSING_BRACKET, state.CurrentChar)
	case 0:
		token = state.NewTokenString(TOKEN_END_OF_FILE, "")
	}

	state.NextChar()

	return token
}
