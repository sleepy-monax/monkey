package ast

import "monkey/token"

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token      token.Token
	Identifier *IdentifierLiteral
	Expression Expression
}

func (statement *LetStatement) statementNode() {
}

func (statement *LetStatement) TokenLiteral() string {
	return statement.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (statement *ReturnStatement) statementNode() {
}

func (statement *ReturnStatement) TokenLiteral() string {
	return statement.Token.Literal
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (statement *BlockStatement) statementNode() {
}

func (statement *BlockStatement) TokenLiteral() string {
	return statement.Token.Literal
}

type ExpressionStatement struct {
	Token      token.Token
	Expression []*Expression
}

func (statement *ExpressionStatement) statementNode() {
}

func (statement *ExpressionStatement) TokenLiteral() string {
	return statement.Token.Literal
}
