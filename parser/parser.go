package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/token"
	"monkey/tokenizer"
	"strconv"
	"testing"
)

type Parser struct {
	depth int
	t     *testing.T

	tokenizer    *tokenizer.Tokenizer
	Errors       []string
	currentToken token.Token
	peekToken    token.Token
}

func New(tokenizer *tokenizer.Tokenizer) *Parser {
	parser := &Parser{tokenizer: tokenizer}

	parser.nextToken()
	parser.nextToken()

	parser.depth = 0
	parser.t = nil

	return parser
}

func NewWithTest(tokenizer *tokenizer.Tokenizer, t *testing.T) *Parser {
	parser := New(tokenizer)

	parser.t = t

	return parser
}

func (parser *Parser) Parse() *ast.Program {
	parser.trace("Parsing")

	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.currentToken.Type != token.EOF {
		stmt := parser.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		parser.nextToken()
	}

	parser.untrace("Parse")

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
		parser.errorf(parser.peekToken, "expected token after %s be %s, got %s instead", parser.currentToken.Type, t, parser.peekToken.Type)
		return false
	}
}

func (parser *Parser) expectCurrent(t token.TokenType) bool {
	if parser.currentTokenIs(t) {
		return true
	} else {
		parser.errorf(parser.peekToken, "expected token be %s, got %s instead", t, parser.currentToken.Type)
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
	parser.trace("parseStatement")

	var statement ast.Statement = nil

	if parser.currentTokenIs(token.Let) {
		statement = parser.parseLetStatement()
	} else if parser.currentTokenIs(token.Return) {
		statement = parser.parseReturnStatement()
	} else if parser.currentTokenIs(token.OpeningBrace) {
		statement = parser.parseBlockStatement()
	} else {
		statement = parser.parseExpressionStatement()
	}

	if !parser.currentTokenIs(token.Semicolon) {
		parser.errorf(parser.currentToken, "Expected %s not %s at end of statement", token.Semicolon, parser.currentToken.Type)
	}

	parser.untrace("parseStatement")

	return statement
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	parser.trace("parseLetStatement")

	statement := &ast.LetStatement{Token: parser.currentToken}

	if !parser.expectPeek(token.Identifier) {

		parser.untrace("parseLetStatement")

		return nil
	}

	statement.Identifier = &ast.IdentifierLiteral{Token: parser.currentToken, Value: parser.currentToken.Literal}

	if !parser.expectPeek(token.Assign) {

		parser.untrace("parseLetStatement")

		return nil
	} else {
		parser.nextToken()
	}

	statement.Expression = parser.parseExpression(PrecedenceLowest)

	if parser.peekTokenIs(token.Semicolon) {
		parser.nextToken()
	}

	parser.untrace("parseLetStatement")

	return statement
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	parser.trace("parseReturnStatement")

	statement := &ast.ReturnStatement{Token: parser.currentToken}

	parser.nextToken()

	statement.Expression = parser.parseExpression(PrecedenceLowest)

	if parser.peekTokenIs(token.Semicolon) {
		parser.nextToken()
	}

	parser.untrace("parseReturnStatement")

	return statement
}

func (parser *Parser) parseBlockStatement() *ast.BlockStatement {
	parser.trace("parseBlockStatement")

	block := &ast.BlockStatement{Token: parser.currentToken}

	block.Statements = []ast.Statement{}

	parser.nextToken()

	for parser.currentToken.Type != token.EOF && parser.currentToken.Type != token.ClosingBrace {
		stmt := parser.parseStatement()

		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		parser.nextToken()
	}

	parser.untrace("parseBlockStatement")

	return block
}

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	parser.trace("parseExpressionStatement")

	stmt := &ast.ExpressionStatement{Token: parser.currentToken}
	stmt.Expression = parser.parseExpression(PrecedenceLowest)

	if parser.peekTokenIs(token.Semicolon) {
		parser.nextToken()
	}

	parser.untrace("parseExpressionStatement")

	return stmt
}

/* --- Expression ----------------------------------------------------------- */

const (
	PrecedenceLowest = iota
	PrecedenceLogic
	PrecedenceEquals
	PrecedenceComparator
	PrecedenceSum
	PrecedenceProduct
	PrecedencePrefix
)

type (
	prefixParseFunction func(*Parser) ast.Expression
	infixParseFunction  func(*Parser, ast.Expression) ast.Expression
)

var precedences map[token.TokenType]int
var prefixParseFunctions map[token.TokenType]prefixParseFunction
var infixParseFunctions map[token.TokenType]infixParseFunction

func init() {
	precedences = map[token.TokenType]int{
		token.And: PrecedenceLogic,
		token.Or:  PrecedenceLogic,

		token.Equal:    PrecedenceEquals,
		token.NotEqual: PrecedenceEquals,

		token.LessThan:   PrecedenceComparator,
		token.BiggerThan: PrecedenceComparator,

		token.Plus:  PrecedenceSum,
		token.Minus: PrecedenceSum,

		token.Asterisk: PrecedenceProduct,
		token.Slash:    PrecedenceProduct,

		token.Bang: PrecedencePrefix,
	}

	prefixParseFunctions = map[token.TokenType]prefixParseFunction{
		token.Identifier: parseIdentifierLiteral,
		token.Integer:    parseIntergerLiteral,
		token.True:       parseBoolLiteral,
		token.False:      parseBoolLiteral,
		token.Function:   parseFunctionLiteral,

		token.Not:   parsePrefixOperatorExpression,
		token.Plus:  parsePrefixOperatorExpression,
		token.Minus: parsePrefixOperatorExpression,

		token.OpeningParenthesis: parseGroupedExpression,
		token.If:                 parseIfExpression,
		token.While:              parseWhileExpression,
	}

	infixParseFunctions = map[token.TokenType]infixParseFunction{
		token.And: parseInfixOperatorExpression,
		token.Not: parseInfixOperatorExpression,

		token.Equal:    parseInfixOperatorExpression,
		token.NotEqual: parseInfixOperatorExpression,

		token.LessThan:   parseInfixOperatorExpression,
		token.BiggerThan: parseInfixOperatorExpression,

		token.Plus:  parseInfixOperatorExpression,
		token.Minus: parseInfixOperatorExpression,

		token.Asterisk: parseInfixOperatorExpression,
		token.Slash:    parseInfixOperatorExpression,
	}
}

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
	parser.trace("parseExpression")

	if parser.currentTokenIs(token.Semicolon) {

		parser.untrace("parseExpression")
		return nil
	}

	prefixParse := prefixParseFunctions[parser.currentToken.Type]

	if prefixParse == nil {
		parser.errorf(parser.currentToken, "no prefix parse function for %s found", parser.currentToken.Type)

		parser.untrace("parseExpression")
		return nil
	}

	leftExp := prefixParse(parser)

	for !parser.peekTokenIs(token.Semicolon) && precedences < parser.peekPrecedence() {
		infixParse := infixParseFunctions[parser.peekToken.Type]

		if infixParse == nil {
			parser.errorf(parser.peekToken, "no infix parse function for %s found", parser.peekToken.Type)

			parser.untrace("parseExpression")
			return leftExp
		}

		parser.nextToken()

		leftExp = infixParse(parser, leftExp)
	}

	parser.untrace("parseExpression")
	return leftExp
}

func parseGroupedExpression(parser *Parser) ast.Expression {
	parser.trace("parseGroupedExpression")

	parser.nextToken()

	exp := parser.parseExpression(PrecedenceLowest)

	if !parser.expectPeek(token.ClosingParenthesis) {

		parser.untrace("parseGroupedExpression")
		return nil
	} else {

		parser.untrace("parseGroupedExpression")
		return exp
	}
}

func parseIfExpression(parser *Parser) ast.Expression {
	parser.trace("parseIfExpression")

	expression := &ast.IfExpression{Token: parser.currentToken}

	if !parser.expectPeek(token.OpeningParenthesis) {
		parser.untrace("parseIfExpression")
		return nil
	}

	parser.nextToken()

	expression.Condition = parser.parseExpression(PrecedenceLowest)

	if !parser.expectPeek(token.ClosingParenthesis) {
		parser.untrace("parseIfExpression")
		return nil
	}

	if !parser.expectPeek(token.OpeningBrace) {
		parser.untrace("parseIfExpression")
		return nil
	}

	expression.Consequence = parser.parseBlockStatement()

	if parser.peekTokenIs(token.Else) {
		parser.nextToken()

		if !parser.expectPeek(token.OpeningBrace) {
			parser.untrace("parseIfExpression")
			return nil
		}

		expression.Alternative = parser.parseBlockStatement()
	}

	parser.untrace("parseIfExpression")
	return expression
}

func parseWhileExpression(parser *Parser) ast.Expression {
	parser.trace("parseWhileExpression")

	expression := &ast.WhileExpression{Token: parser.currentToken}

	if !parser.expectPeek(token.OpeningParenthesis) {
		parser.untrace("parseWhileExpression")
		return nil
	}

	parser.nextToken()

	expression.Condition = parser.parseExpression(PrecedenceLowest)

	if !parser.expectPeek(token.ClosingParenthesis) {
		parser.untrace("parseWhileExpression")
		return nil
	}

	if !parser.expectPeek(token.OpeningBrace) {
		parser.untrace("parseWhileExpression")
		return nil
	}

	expression.Body = parser.parseBlockStatement()

	parser.untrace("parseWhileExpression")
	return expression
}

func parsePrefixOperatorExpression(parser *Parser) ast.Expression {
	parser.trace("parsePrefixOperatorExpression")

	expression := &ast.PrefixOperatorExpression{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Literal,
	}

	parser.nextToken()
	expression.Right = parser.parseExpression(PrecedencePrefix)

	parser.untrace("parsePrefixOperatorExpression")
	return expression
}

func parseInfixOperatorExpression(parser *Parser, left ast.Expression) ast.Expression {
	parser.trace("parseInfixOperatorExpression")

	expression := &ast.InfixOperatorExpression{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Literal,
		Left:     left,
	}

	precedences := parser.currentPrecedence()
	parser.nextToken()
	expression.Right = parser.parseExpression(precedences)

	parser.untrace("parseInfixOperatorExpression")
	return expression
}

func parsePostfixOperatorExpression(parser *Parser, left ast.Expression) ast.Expression {
	expression := &ast.PostfixOperatorExpression{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Literal,
		Left:     left,
	}

	parser.nextToken()

	return expression
}

func unexpectedInfixToken(parser *Parser, left ast.Expression) ast.Expression {
	parser.errorf(parser.currentToken, "unexpected infix token %s", parser.currentToken.Type)
	return nil
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

func parseFunctionLiteral(parser *Parser) ast.Expression {
	parser.trace("parseFunctionLiteral")

	function := &ast.FunctionLiteral{Token: parser.currentToken}

	if !parser.expectPeek(token.OpeningParenthesis) {
		parser.untrace("parseFunctionLiteral")
		return nil
	}

	parser.nextToken()

	function.Parameters = parser.parseFunctionParameters()

	if !parser.expectPeek(token.ClosingParenthesis) {
		parser.untrace("parseFunctionLiteral")
		return nil
	}

	if !parser.expectPeek(token.OpeningBrace) {
		parser.untrace("parseFunctionLiteral")
		return nil
	}

	function.Body = parser.parseBlockStatement()

	parser.untrace("parseFunctionLiteral")
	return function
}

func (parser *Parser) parseFunctionParameters() []*ast.IdentifierLiteral {
	parser.trace("parseFunctionParameters")

	identifiers := []*ast.IdentifierLiteral{}

	for !parser.peekTokenIs(token.ClosingParenthesis) && parser.expectCurrent(token.Identifier) {
		identifiers = append(identifiers, &ast.IdentifierLiteral{Token: parser.currentToken, Value: parser.currentToken.Literal})

		parser.nextToken()

		if parser.currentTokenIs(token.Comma) {
			parser.nextToken()
		}
	}

	if parser.currentTokenIs(token.Identifier) {
		identifiers = append(identifiers, &ast.IdentifierLiteral{Token: parser.currentToken, Value: parser.currentToken.Literal})
	}

	if parser.currentTokenIs(token.ClosingParenthesis) {
		parser.untrace("parseFunctionParameters")
		return nil
	}

	parser.untrace("parseFunctionParameters")
	return identifiers
}
