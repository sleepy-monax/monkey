package tokenizer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};
	
	let result = add(five, ten);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Identifier, "five"},
		{token.Assign, "="},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "ten"},
		{token.Assign, "="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.OpeningParenthesis, "("},
		{token.Identifier, "x"},
		{token.Comma, ","},
		{token.Identifier, "y"},
		{token.ClosingParenthesis, ")"},
		{token.OpeningBrace, "{"},
		{token.Identifier, "x"},
		{token.Plus, "+"},
		{token.Identifier, "y"},
		{token.Semicolon, ";"},
		{token.ClosingBrace, "}"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "result"},
		{token.Assign, "="},
		{token.Identifier, "add"},
		{token.OpeningParenthesis, "("},
		{token.Identifier, "five"},
		{token.Comma, ","},
		{token.Identifier, "ten"},
		{token.ClosingParenthesis, ")"},
		{token.Semicolon, ";"},
	}

	state := New(input)

	for i, test := range tests {
		tok := state.NextToken()

		t.Logf("Token{%q, %q, ln%d col%d}", tok.Type, tok.Literal, tok.Line, tok.Column)

		if tok.Type != test.expectedType {
			t.Fatalf("test[%d] - TokenType wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("test[%d] - Literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}
