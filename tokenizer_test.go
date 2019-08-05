package monkey

import "testing"

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{TOKEN_EQUAL, "="},
		{TOKEN_PLUS, "+"},
		{TOKEN_OPENING_PARENTHESIS, "("},
		{TOKEN_CLOSING_PARENTHESIS, ")"},
		{TOKEN_OPENING_BRACE, "{"},
		{TOKEN_CLOSING_BRACE, "}"},
		{TOKEN_COMMA, ","},
		{TOKEN_SEMICOLON, ";"},
		{TOKEN_END_OF_FILE, ""},
	}

	state := NewTokenizerState(input)

	for i, test := range tests {
		token := state.NextToken()

		t.Logf("Token{%q, %q}", token.Type, token.Literal)

		if token.Type != test.expectedType {
			t.Fatalf("test[%d] - TokenType wrong. expected=%q, got=%q", i, test.expectedType, token.Type)
		}

		if token.Literal != test.expectedLiteral {
			t.Fatalf("test[%d] - Literal wrong. expected=%q, got=%q", i, test.expectedLiteral, token.Literal)
		}
	}
}
