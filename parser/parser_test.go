package parser

import (
	"monkey/ast"
	"monkey/tokenizer"
	"testing"
)

func TestParser(t *testing.T) {
	testParseExpect(t, "a + b;", "(a + b);", 1)
	testParseExpect(t, "a + b + c;", "((a + b) + c);", 1)
	testParseExpect(t, "a + b * c;", "(a + (b * c));", 1)
	testParseExpect(t, "(a + b) * c;", "((a + b) * c);", 1)

	testParseExpect(t, "let a = 10;", "let a = 10;", 1)

	testParseExpect(t, "function (){ doStuff; };", "function(){doStuff;};", 1)
	testParseExpect(t, "function (a){ return a; };", "function(a){return a;};", 1)
	testParseExpect(t, "function (a, b){ return a + b; };", "function(a,b){return (a + b);};", 1)

	testParseExpectError(t, "(a+b)e;")
}

func testParseProgram(t *testing.T, input string, expectedStatements int) *ast.Program {
	tok := tokenizer.New(input)
	p := NewWithTest(tok, t)

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

func testParseExpect(t *testing.T, input string, output string, expectedStatements int) {
	t.Logf("--- testParseExpect '" + input + "' expect '" + output + "' ---")

	program := testParseProgram(t, input, expectedStatements)

	if program.String() != output {
		t.Errorf("testParseExpect failled expected '%s' to be '%s' got '%s'", input, output, program.String())
	}
}

func testParseExpectError(t *testing.T, input string) {
	t.Logf("--- testParseExpectError '" + input + "' ---")

	tok := tokenizer.New(input)
	p := NewWithTest(tok, t)

	program := p.Parse()

	if len(p.Errors) == 0 {
		t.Errorf("testParseExpectError failled expected '%s' to be erroneous, got '%s'", input, program.String())
	}
}

func checkParserErrors(t *testing.T, parser *Parser) {
	errors := parser.Errors
	if len(errors) != 0 {
		t.Errorf("Parser has %d errors", len(errors))

		for _, msg := range errors {
			t.Errorf("parser error: %q", msg)
		}
	}
}
