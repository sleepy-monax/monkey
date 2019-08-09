package ast

import "monkey/token"

type IdentifierLiteral struct {
	Token token.Token
	Value string
}

func (expression *IdentifierLiteral) expressionNode()      {}
func (expression *IdentifierLiteral) TokenLiteral() string { return expression.Token.Literal }

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (expression *BooleanLiteral) expressionNode()      {}
func (expression *BooleanLiteral) TokenLiteral() string { return expression.Token.Literal }

type IntegerLiteral struct {
	Token token.Token
	Value int
}

func (expression *IntegerLiteral) expressionNode()      {}
func (expression *IntegerLiteral) TokenLiteral() string { return expression.Token.Literal }

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*IdentifierLiteral
	Body       *BlockStatement
}

func (expression *FunctionLiteral) expressionNode()      {}
func (expression *FunctionLiteral) TokenLiteral() string { return expression.Token.Literal }
