package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/token"
	"monkey/tokenizer"
	"strconv"
)

type Parser struct {
	tokenizer    *tokenizer.Tokenizer
	Errors       []string
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
		parser.errorf(parser.peekToken, "expected next token be %s, got %s instead", t, parser.peekToken.Type)
		return false
	}
}

func (parser *Parser) currentTokenIs(t token.TokenType) bool {
	return parser.currentToken.Type == t
}

func (parser *Parser) peekTokenIs(t token.TokenType) bool {
	return parser.peekToken.Type == t
}

func (parser *Parser) errorf(tok token.Token, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	parser.Errors = append(parser.Errors, fmt.Sprintf("Ln %d, Col %d: %s", tok.Line, tok.Column, msg))
}

func (parser *Parser) error(tok token.Token, msg string) {
	parser.Errors = append(parser.Errors, msg)
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

	statement.Expression = parser.parseExpression(PrecedenceLowest)

	if parser.peekTokenIs(token.Semicolon) {
		parser.nextToken()
	}

	return statement
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: parser.currentToken}

	parser.nextToken()

	statement.Expression = parser.parseExpression(PrecedenceLowest)

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
	stmt := &ast.ExpressionStatement{Token: parser.currentToken}
	stmt.Expression = parser.parseExpression(PrecedenceLowest)

	if parser.peekTokenIs(token.Semicolon) {
		parser.nextToken()
	}

	return stmt
}

/* --- Expression ----------------------------------------------------------- */

const (
	PrecedenceLowest = iota
	PrecedenceEquals
	PrecedenceComparator
	PrecedenceSum
	PrecedenceProduct
	PrecedencePrefix
)

var precedences = map[token.TokenType]int{
	token.Equal:    PrecedenceEquals,
	token.NotEqual: PrecedenceEquals,

	token.LessThan:   PrecedenceComparator,
	token.BiggerThan: PrecedenceComparator,

	token.Plus:  PrecedenceSum,
	token.Minus: PrecedenceSum,

	token.Assign: PrecedenceProduct,
	token.Slash:  PrecedenceProduct,
}

type (
	prefixParseFunction func(*Parser) ast.Expression
	infixParseFunction  func(*Parser, ast.Expression) ast.Expression
)

var prefixParseFunctions = map[token.TokenType]prefixParseFunction{
	token.Identifier: parseIdentifierLiteral,
	token.Integer:    parseIntergerLiteral,
	token.True:       parseBoolLiteral,
	token.False:      parseBoolLiteral,
}

var infixParseFunctions = map[token.TokenType]infixParseFunction{}

func (parser *Parser) peekPrecedence() int {
	if precedence, ok := precedences[parser.peekToken.Type]; ok {
		return precedence
	}

	return PrecedenceLowest
}

func (parser *Parser) currentPrecedence() int {
	if precedence, ok := precedences[parser.currentToken.Type]; ok {
		return precedence
	}

	return PrecedenceLowest
}

func (parser *Parser) parseExpression(precedences int) ast.Expression {
	prefixParse := prefixParseFunctions[parser.currentToken.Type]

	if prefixParse == nil {
		parser.errorf(parser.currentToken, "no prefix parse function for %s found", parser.currentToken.Type)
		return nil
	}

	leftExp := prefixParse(parser)

	for !parser.peekTokenIs(token.Semicolon) && precedences < parser.peekPrecedence() {
		infixParse := infixParseFunctions[parser.peekToken.Type]

		if infixParse == nil {
			return leftExp
		}

		parser.nextToken()

		leftExp = infixParse(parser, leftExp)
	}

	return leftExp
}

/* --- Literals ------------------------------------------------------------- */

func parseIdentifierLiteral(parser *Parser) ast.Expression {
	return &ast.IdentifierLiteral{Token: parser.currentToken, Value: parser.currentToken.Literal}
}

func parseIntergerLiteral(parser *Parser) ast.Expression {
	value, ok := strconv.ParseInt(parser.currentToken.Literal, 0, 64)
	if ok == nil {
		return &ast.IntegerLiteral{Token: parser.currentToken, Value: value}
	} else {
		return nil
	}
}

func parseBoolLiteral(parser *Parser) ast.Expression {
	if parser.currentTokenIs(token.True) {
		return &ast.BooleanLiteral{Token: parser.currentToken, Value: true}
	} else {
		return &ast.BooleanLiteral{Token: parser.currentToken, Value: false}
	}
}
