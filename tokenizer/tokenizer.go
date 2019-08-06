package tokenizer

import "monkey/token"

type Tokenizer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte

	currentLine   int
	currentColumn int
}

func New(input string) *Tokenizer {
	state := &Tokenizer{
		input:       input,
		currentLine: 1,
		currentChar: 1,
	}

	state.readChar()

	return state
}

func (state *Tokenizer) newToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:   tokenType,
		Column: state.currentColumn,
		Line:   state.currentLine,
	}
}

func (state *Tokenizer) newTokenChar(tokenType token.TokenType, char byte) token.Token {
	token := state.newToken(tokenType)
	token.Literal = string(char)

	return token
}

func (state *Tokenizer) newTokenString(tokenType token.TokenType, literal string) token.Token {
	token := state.newToken(tokenType)
	token.Literal = literal

	return token
}

func (state *Tokenizer) readChar() {
	if state.readPosition >= len(state.input) {
		state.currentChar = 0
	} else {
		state.currentChar = state.input[state.readPosition]
	}

	if state.currentChar == '\n' {
		state.currentLine++
		state.currentColumn = 0
	} else {
		state.currentColumn++
	}

	state.position = state.readPosition
	state.readPosition++
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isIdentifier(char byte) bool {
	return isLetter(char) || char == '_'
}

func isWitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func isNumber(char byte) bool {
	return '0' <= char && char <= '9'
}

func (state *Tokenizer) eatWhitespace() {
	for isWitespace(state.currentChar) {
		state.readChar()
	}
}

func lookupIdentifier(identifier string) token.TokenType {
	if tokenType, ok := token.Keywords[identifier]; ok {
		return tokenType
	}

	return token.Identifier
}

func (state *Tokenizer) readIdentifier() token.Token {
	token := state.newTokenString(token.Identifier, "")

	position := state.position

	for isIdentifier(state.currentChar) {
		state.readChar()
	}

	token.Literal = state.input[position:state.position]
	token.Type = lookupIdentifier(token.Literal)

	return token
}

func (state *Tokenizer) readNumber() token.Token {
	token := state.newToken(token.Integer)

	position := state.position

	for isNumber(state.currentChar) {
		state.readChar()
	}

	token.Literal = state.input[position:state.position]

	return token
}

// NextToken get the next token a currentPosition in the input string.
func (state *Tokenizer) NextToken() token.Token {
	state.eatWhitespace()

	tok := state.newTokenChar(token.Illegal, state.currentChar)

	switch state.currentChar {
	case '=':
		tok = state.newTokenChar(token.Assign, state.currentChar)
	case '+':
		tok = state.newTokenChar(token.Plus, state.currentChar)
	case '-':
		tok = state.newTokenChar(token.Minus, state.currentChar)
	case '!':
		tok = state.newTokenChar(token.Bang, state.currentChar)
	case '*':
		tok = state.newTokenChar(token.Asterisk, state.currentChar)
	case '/':
		tok = state.newTokenChar(token.Slash, state.currentChar)
	case '<':
		tok = state.newTokenChar(token.LessThan, state.currentChar)
	case '>':
		tok = state.newTokenChar(token.BiggerThan, state.currentChar)
	case ',':
		tok = state.newTokenChar(token.Comma, state.currentChar)
	case ';':
		tok = state.newTokenChar(token.Semicolon, state.currentChar)
	case '(':
		tok = state.newTokenChar(token.OpeningParenthesis, state.currentChar)
	case ')':
		tok = state.newTokenChar(token.ClosingParenthesis, state.currentChar)
	case '{':
		tok = state.newTokenChar(token.OpeningBrace, state.currentChar)
	case '}':
		tok = state.newTokenChar(token.ClosingBrace, state.currentChar)
	case '[':
		tok = state.newTokenChar(token.OpeningBracket, state.currentChar)
	case ']':
		tok = state.newTokenChar(token.ClosingBracket, state.currentChar)
	case 0:
		tok = state.newTokenString(token.EOF, "")

	default:
		if isIdentifier(state.currentChar) {
			return state.readIdentifier()
		} else if isNumber(state.currentChar) {
			return state.readNumber()
		}
	}

	state.readChar()

	return tok
}
