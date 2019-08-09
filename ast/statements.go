package ast

import (
	"bytes"
	"monkey/token"
)

type Statement interface {
	Node
	statementNode()
}

/* --- Let Statement -------------------------------------------------------- */

type LetStatement struct {
	Token      token.Token
	Identifier *IdentifierLiteral
	Expression Expression
}

func (statement *LetStatement) statementNode()       {}
func (statement *LetStatement) TokenLiteral() string { return statement.Token.Literal }
func (statement *LetStatement) String() string {
	if statement == nil {
		return ""
	}

	var out bytes.Buffer
	out.WriteString(statement.TokenLiteral() + " ")

	if statement.Identifier != nil {
		out.WriteString(statement.Identifier.String() + " ")
	}

	out.WriteString("= ")

	if statement.Expression != nil {
		out.WriteString(statement.Expression.String())
	}

	return out.String()
}

/* --- Return Statement ----------------------------------------------------- */

type ReturnStatement struct {
	Token      token.Token
	Expression Expression
}

func (statement *ReturnStatement) statementNode()       {}
func (statement *ReturnStatement) TokenLiteral() string { return statement.Token.Literal }
func (statement *ReturnStatement) String() string {
	if statement == nil {
		return ""
	}

	var out bytes.Buffer
	out.WriteString(statement.TokenLiteral() + " ")

	if statement.Expression != nil {
		out.WriteString(statement.Expression.String())
	}

	return out.String()
}

/* --- Block Statement ------------------------------------------------------ */

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (statement *BlockStatement) statementNode()       {}
func (statement *BlockStatement) TokenLiteral() string { return statement.Token.Literal }
func (statement *BlockStatement) String() string {
	if statement == nil {
		return ""
	}

	var out bytes.Buffer

	out.WriteString("{")

	for _, stmt := range statement.Statements {

		out.WriteString(stmt.String())
		out.WriteString(";")

	}

	out.WriteString("}")

	return out.String()
}

/* --- Expression Statement ------------------------------------------------- */

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (statement *ExpressionStatement) statementNode()       {}
func (statement *ExpressionStatement) TokenLiteral() string { return statement.Token.Literal }
func (statement *ExpressionStatement) String() string {
	if statement == nil {
		return ""
	}

	if statement.Expression != nil {
		return statement.Expression.String()
	} else {
		return ""
	}
}
