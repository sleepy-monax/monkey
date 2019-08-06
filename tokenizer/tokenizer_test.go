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
	
	let result = add(five, ten);

	=+-!*/<>;

	if (true) { return true; } else { return false; }

	let truth = ten ==10!= 5;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// let five = 5;
		{token.Let, "let"},
		{token.Identifier, "five"},
		{token.Assign, "="},
		{token.Integer, "5"},
		{token.Semicolon, ";"},

		// let ten = 10;
		{token.Let, "let"},
		{token.Identifier, "ten"},
		{token.Assign, "="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},

		// let add = fn(x, y) { x + y; };
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

		// let result = add(five, ten);
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

		// =+-!*/<>;
		{token.Assign, "="},
		{token.Plus, "+"},
		{token.Minus, "-"},
		{token.Bang, "!"},
		{token.Asterisk, "*"},
		{token.Slash, "/"},
		{token.LessThan, "<"},
		{token.BiggerThan, ">"},
		{token.Semicolon, ";"},

		// if (true) { return true; } else { return false; }
		{token.If, "if"},
		{token.OpeningParenthesis, "("},
		{token.True, "true"},
		{token.ClosingParenthesis, ")"},
		{token.OpeningBrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.ClosingBrace, "}"},
		{token.Else, "else"},
		{token.OpeningBrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.ClosingBrace, "}"},

		// let truth = ten ==10!= 5;
		{token.Let, "let"},
		{token.Identifier, "truth"},
		{token.Assign, "="},
		{token.Identifier, "ten"},
		{token.Equal, "=="},
		{token.Integer, "10"},
		{token.NotEqual, "!="},
		{token.Integer, "5"},
		{token.Semicolon, ";"},

		{token.EOF, ""},
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
