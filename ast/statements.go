package ast

import (
	"bytes"
	"fmt"
	"monkey/token"
)

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token      token.Token
	Identifier *IdentifierLiteral
	Expression Expression
}

func (statement *LetStatement) statementNode()       {}
func (statement *LetStatement) TokenLiteral() string { return statement.Token.Literal }
func (statement *LetStatement) String() string {
	return fmt.Sprintf("%s %s = %s", statement.TokenLiteral(), statement.Identifier.String(), statement.Expression.String())
}

type ReturnStatement struct {
	Token      token.Token
	Expression Expression
}

func (statement *ReturnStatement) statementNode() {}

func (statement *ReturnStatement) TokenLiteral() string { return statement.Token.Literal }
func (statement *ReturnStatement) String() string {
	return fmt.Sprintf("%s %s", statement.TokenLiteral(), statement.Expression.String())
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (statement *BlockStatement) statementNode()       {}
func (statement *BlockStatement) TokenLiteral() string { return statement.Token.Literal }
func (statement *BlockStatement) String() string {
	var out bytes.Buffer

	out.WriteString("{")

	for _, stmt := range statement.Statements {

		out.WriteString(stmt.String())
		out.WriteString(";")

	}

	out.WriteString("}")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (statement *ExpressionStatement) statementNode() {}

func (statement *ExpressionStatement) TokenLiteral() string { return statement.Token.Literal }

func (statement *ExpressionStatement) String() string {
	return statement.Expression.String()
}
