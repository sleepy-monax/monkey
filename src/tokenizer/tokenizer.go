package tokenizer

// TokenType is the type of the token
type TokenType string

const (
	TOKEN_ILLEGAL     = "ILLEGAL"
	TOKEN_END_OF_FILE = "END_OF_FILE"

	// Identifier and literals
	TOKEN_IDENTIFIER = "IDENTIFIER"
	TOKEN_NUMBER     = "NUMBER"

	// Operators
	TOKEN_EQUAL = "EQUAL"
	TOKEN_PLUS  = "PLUS"

	// Delemiters
	TOKEN_COMMA               = "COMMA"
	TOKEN_SEMICOLON           = "SEMICOLON"
	TOKEN_OPENING_PARENTHESIS = "OPENING_PARENTHESIS"
	TOKEN_CLOSING_PARENTHESIS = "CLOSING_PARENTHESIS"
	TOKEN_OPENING_BRACE       = "OPENING_BRACE"
	TOKEN_CLOSING_BRACE       = "CLOSING_BRACE"
	TOKEN_OPENING_BRACKET     = "OPENING_BRACKET"
	TOKEN_CLOSING_BRACKET     = "CLOSING_BRACKET"

	// Keyword
	TOKEN_FUNCTION = "FUNCTION"
	TOKEN_LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  TOKEN_FUNCTION,
	"let": TOKEN_LET,
}

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

	state.ReadChar()

	return state
}

func (state *TokenizerState) NewToken(tokenType TokenType) Token {
	return Token{
		Type:   tokenType,
		Column: state.CurrentColumn,
		Line:   state.CurrentLine,
	}
}

// NewTokenChar create a new token from a char
func (state *TokenizerState) NewTokenChar(tokenType TokenType, char byte) Token {
	token := state.NewToken(tokenType)
	token.Literal = string(char)

	return token
}

// NewTokenString create a new token from a string
func (state *TokenizerState) NewTokenString(tokenType TokenType, literal string) Token {
	token := state.NewToken(tokenType)
	token.Literal = literal

	return token
}

// ReadChar get the next char form the input string.
func (state *TokenizerState) ReadChar() {
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

// IsLetter check if the char is a letter.
func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

// IsIdentifier check if the char is part of a valid identifier.
func IsIdentifier(char byte) bool {
	return IsLetter(char) || char == '_'
}

func IsWitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func IsNumber(char byte) bool {
	return '0' <= char && char <= '9'
}

func (state *TokenizerState) EatWhitespace() {
	for IsWitespace(state.CurrentChar) {
		state.ReadChar()
	}
}

// LookupIdentifier chech if the given identifier is a keyword or an identifier.
func LookupIdentifier(identifier string) TokenType {
	if tokenType, ok := keywords[identifier]; ok {
		return tokenType
	}

	return TOKEN_IDENTIFIER
}

// ReadIdentifier read the next identifier from the input string.
func (state *TokenizerState) ReadIdentifier() Token {
	token := state.NewTokenString(TOKEN_IDENTIFIER, "")

	position := state.Position

	for IsIdentifier(state.CurrentChar) {
		state.ReadChar()
	}

	token.Literal = state.Input[position:state.Position]
	token.Type = LookupIdentifier(token.Literal)

	return token
}

func (state *TokenizerState) ReadNumber() Token {
	token := state.NewToken(TOKEN_NUMBER)

	position := state.Position

	for IsNumber(state.CurrentChar) {
		state.ReadChar()
	}

	token.Literal = state.Input[position:state.Position]

	return token
}

// NextToken get the next token a currentPosition in the input string.
func (state *TokenizerState) NextToken() Token {
	state.EatWhitespace()

	token := state.NewTokenChar(TOKEN_ILLEGAL, state.CurrentChar)

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

	default:
		if IsIdentifier(state.CurrentChar) {
			return state.ReadIdentifier()
		} else if IsNumber(state.CurrentChar) {
			return state.ReadNumber()
		}
	}

	state.ReadChar()

	return token
}
