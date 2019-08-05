package monkey

import "testing"

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};
	
	let result = add(five, ten);`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{TOKEN_LET, "let"},
		{TOKEN_IDENTIFIER, "five"},
		{TOKEN_EQUAL, "="},
		{TOKEN_NUMBER, "5"},
		{TOKEN_SEMICOLON, ";"},
		{TOKEN_LET, "let"},
		{TOKEN_IDENTIFIER, "ten"},
		{TOKEN_EQUAL, "="},
		{TOKEN_NUMBER, "10"},
		{TOKEN_SEMICOLON, ";"},
		{TOKEN_LET, "let"},
		{TOKEN_IDENTIFIER, "add"},
		{TOKEN_EQUAL, "="},
		{TOKEN_FUNCTION, "fn"},
		{TOKEN_OPENING_PARENTHESIS, "("},
		{TOKEN_IDENTIFIER, "x"},
		{TOKEN_COMMA, ","},
		{TOKEN_IDENTIFIER, "y"},
		{TOKEN_CLOSING_PARENTHESIS, ")"},
		{TOKEN_OPENING_BRACE, "{"},
		{TOKEN_IDENTIFIER, "x"},
		{TOKEN_PLUS, "+"},
		{TOKEN_IDENTIFIER, "y"},
		{TOKEN_SEMICOLON, ";"},
		{TOKEN_CLOSING_BRACE, "}"},
		{TOKEN_SEMICOLON, ";"},
		{TOKEN_LET, "let"},
		{TOKEN_IDENTIFIER, "result"},
		{TOKEN_EQUAL, "="},
		{TOKEN_IDENTIFIER, "add"},
		{TOKEN_OPENING_PARENTHESIS, "("},
		{TOKEN_IDENTIFIER, "five"},
		{TOKEN_COMMA, ","},
		{TOKEN_IDENTIFIER, "ten"},
		{TOKEN_CLOSING_PARENTHESIS, ")"},
		{TOKEN_SEMICOLON, ";"},
	}

	state := NewTokenizerState(input)

	for i, test := range tests {
		token := state.NextToken()

		t.Logf("Token{%q, %q, ln%d col%d}", token.Type, token.Literal, token.Line, token.Column)

		if token.Type != test.expectedType {
			t.Fatalf("test[%d] - TokenType wrong. expected=%q, got=%q", i, test.expectedType, token.Type)
		}

		if token.Literal != test.expectedLiteral {
			t.Fatalf("test[%d] - Literal wrong. expected=%q, got=%q", i, test.expectedLiteral, token.Literal)
		}
	}
}
