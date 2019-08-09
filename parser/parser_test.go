package parser

import (
	"monkey/ast"
	"monkey/tokenizer"
	"testing"
)

func testParseProgram(t *testing.T, input string, expectedStatements int) *ast.Program {
	tok := tokenizer.New(input)
	p := New(tok)

	program := p.Parse()

	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("Parse() returned nil")
	}

	if len(program.Statements) != expectedStatements {
		t.Fatalf("program.Statements does not contain %d statements. got=%d", expectedStatements, len(program.Statements))
	}

	return program
}

func checkParserErrors(t *testing.T, parser *Parser) {
	errors := parser.Errors
	if len(errors) != 0 {
		t.Errorf("Parser has %d errors", len(errors))

		for _, msg := range errors {
			t.Errorf("parser error: %q", msg)
		}

		t.FailNow()
	}
}

func TestLetStatement(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
	}{
		{"let x = 5;", "x"},
		{"let y = 10;", "y"},
		{"let foobar = 838383;", "foobar"},
	}

	for _, test := range tests {
		program := testParseProgram(t, test.input, 1)

		statement := program.Statements[0]

		if statement.TokenLiteral() != "let" {
			t.Errorf("s.TokenLiteral() not 'let'. got=%q", statement.TokenLiteral())
		}

		letStatement, ok := statement.(*ast.LetStatement)

		if !ok {
			t.Errorf("statement not *ast.LetStatement. got=%T", statement)
		}

		if letStatement.Identifier.Value != test.expectedIdentifier {
			t.Errorf("letStmt.Name.Value not '%s'. got=%s", test.expectedIdentifier, letStatement.Identifier.Value)
		}

		if letStatement.Identifier.TokenLiteral() != test.expectedIdentifier {
			t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", test.expectedIdentifier, letStatement.Identifier.TokenLiteral())
		}
	}
}

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"return 5;", 5},
		{"return true;", true},
		{"return foobar;", "foobar"},
	}

	for _, test := range tests {
		program := testParseProgram(t, test.input, 1)

		statement := program.Statements[0]

		if statement.TokenLiteral() != "return" {
			t.Errorf("s.TokenLiteral() not 'return'. got=%q", statement.TokenLiteral())
		}

		_, ok := statement.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("statement not *ast.returnStatement. got=%T", statement)
		}
	}
}

func TestExpressionStatement(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"5;", 5},
		{"true;", true},
		{"foobar;", "foobar"},
	}

	for _, test := range tests {
		program := testParseProgram(t, test.input, 1)

		statement := program.Statements[0]

		_, ok := statement.(*ast.ExpressionStatement)

		if !ok {
			t.Errorf("statement not *ast.ExpressionStatement. got=%T", statement)
		}
	}
}
