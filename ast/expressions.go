package ast

import (
	"bytes"
	"monkey/token"
)

type Expression interface {
	Node
	expressionNode()
}

/* --- Prefix Operator Expression ------------------------------------------- */

type PrefixOperatorExpression struct {
	Token token.Token

	Operator string
	Right    Expression
}

func (expression *PrefixOperatorExpression) expressionNode()      {}
func (expression *PrefixOperatorExpression) TokenLiteral() string { return expression.Token.Literal }
func (expression *PrefixOperatorExpression) String() string {
	if expression == nil {
		return ""
	}

	var out bytes.Buffer

	out.WriteString("(" + expression.Operator + " ")

	if expression.Right != nil {
		out.WriteString(expression.Right.String())
	}

	out.WriteString(")")

	return out.String()
}

/* --- Infix Operator Expression -------------------------------------------- */

type InfixOperatorExpression struct {
	Token token.Token

	Operator string
	Left     Expression
	Right    Expression
}

func (expression *InfixOperatorExpression) expressionNode()      {}
func (expression *InfixOperatorExpression) TokenLiteral() string { return expression.Token.Literal }
func (expression *InfixOperatorExpression) String() string {
	if expression == nil {
		return ""
	}

	var out bytes.Buffer
	out.WriteString("(")

	if expression.Left != nil {
		out.WriteString(expression.Left.String())
	}

	out.WriteString(" " + expression.Operator + " ")

	if expression.Right != nil {
		out.WriteString(expression.Right.String())
	}

	out.WriteString(")")

	return out.String()
}

/* --- Postfix Operator Expression ------------------------------------------ */

type PostfixOperatorExpression struct {
	Token token.Token

	Operator string
	Left     Expression
}

func (expression *PostfixOperatorExpression) expressionNode()      {}
func (expression *PostfixOperatorExpression) TokenLiteral() string { return expression.Token.Literal }
func (expression *PostfixOperatorExpression) String() string {
	if expression == nil {
		return ""
	}

	var out bytes.Buffer
	out.WriteString("(")

	if expression.Left != nil {
		out.WriteString(expression.Left.String())
	}

	out.WriteString(" " + expression.Operator + ")")

	return out.String()
}
