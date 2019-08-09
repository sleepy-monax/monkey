package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/token"
	"monkey/tokenizer"
)

type Parser struct {
	tokenizer    *tokenizer.Tokenizer
	errors       []string
	currentToken token.Token
	peekToken    token.Token
}

func New(tokenizer *tokenizer.Tokenizer) *Parser {
	parser := &Parser{tokenizer: tokenizer}

	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.currentToken.Type != token.EOF {
		stmt := parser.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		parser.nextToken()
	}

	return program
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.tokenizer.NextToken()
}

func (parser *Parser) expectPeek(t token.TokenType) bool {
	if parser.peekTokenIs(t) {
		parser.nextToken()
		return true
	} else {
		parser.errors = append(parser.errors, fmt.Sprintf("ln%d col%d expected next token be %s, got %s instead", parser.peekToken.Line, parser.peekToken.Column, t, parser.peekToken.Type))
		return false
	}
}

func (parser *Parser) currentTokenIs(t token.TokenType) bool {
	return parser.currentToken.Type == t
}

func (parser *Parser) peekTokenIs(t token.TokenType) bool {
	return parser.peekToken.Type == t
}

/* --- Statements ----------------------------------------------------------- */

func (parser *Parser) parseStatement() ast.Statement {
	if parser.currentTokenIs(token.Let) {
		return parser.parseLetStatement()
	} else if parser.currentTokenIs(token.Return) {
		return parser.parseReturnStatement()
	} else if parser.currentTokenIs(token.OpeningBracket) {
		return parser.parseBlockStatement()
	} else {
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: parser.currentToken}

	if !parser.expectPeek(token.Identifier) {
		return nil
	}

	statement.Identifier = &ast.IdentifierLiteral{Token: parser.currentToken, Value: parser.currentToken.Literal}

	if !parser.expectPeek(token.Assign) {
		return nil
	} else {
		parser.nextToken()
	}

	statement.Expression = parser.parseExpression()

	if parser.peekTokenIs(token.Semicolon) {
		parser.nextToken()
	}

	return statement
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: parser.currentToken}

	parser.nextToken()

	statement.ReturnValue = parser.parseExpression()

	if parser.peekTokenIs(token.Semicolon) {
		parser.nextToken()
	}

	return statement
}

func (parser *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: parser.currentToken}

	block.Statements = []ast.Statement{}

	for parser.currentToken.Type != token.EOF {
		stmt := parser.parseStatement()

		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		parser.nextToken()
	}

	return block
}

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	return nil
}

/* --- Expression ----------------------------------------------------------- */

func (parser *Parser) parseExpression() ast.Expression {
	for !parser.peekTokenIs(token.Semicolon) {
		parser.nextToken()
	}

	return nil
}
