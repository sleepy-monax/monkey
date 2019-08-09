package ast

import (
	"fmt"
	"monkey/token"
)

type IdentifierLiteral struct {
	Token token.Token
	Value string
}

func (expression *IdentifierLiteral) expressionNode()      {}
func (expression *IdentifierLiteral) TokenLiteral() string { return expression.Token.Literal }
func (expression *IdentifierLiteral) String() string       { return expression.Value }

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (expression *BooleanLiteral) expressionNode()      {}
func (expression *BooleanLiteral) TokenLiteral() string { return expression.Token.Literal }
func (expression *BooleanLiteral) String() string {
	if expression.Value {
		return "true"
	} else {
		return "false"
	}
}

type IntegerLiteral struct {
	Token token.Token
	Value int
}

func (expression *IntegerLiteral) expressionNode()      {}
func (expression *IntegerLiteral) TokenLiteral() string { return expression.Token.Literal }
func (expression *IntegerLiteral) String() string       { return fmt.Sprintf("%d", expression.Value) }

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*IdentifierLiteral
	Body       *BlockStatement
}

func (expression *FunctionLiteral) expressionNode()      {}
func (expression *FunctionLiteral) TokenLiteral() string { return expression.Token.Literal }
func (expression *FunctionLiteral) String() string {
	return ""
}
