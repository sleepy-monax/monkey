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

/* --- If Expression -------------------------------------------------------- */

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (expression *IfExpression) expressionNode()      {}
func (expression *IfExpression) TokenLiteral() string { return expression.Token.Literal }
func (expression *IfExpression) String() string {
	if expression == nil {
		return ""
	}

	var out bytes.Buffer
	out.WriteString(expression.TokenLiteral())

	if expression.Condition != nil {
		out.WriteString("(" + expression.Condition.String() + ")")
	}

	if expression.Consequence != nil {
		out.WriteString(expression.Consequence.String())
	}

	if expression.Alternative != nil {
		out.WriteString("else")
		out.WriteString(expression.Alternative.String())
	}

	return out.String()
}

type WhileExpression struct {
	Token     token.Token
	Condition Expression
	Body      *BlockStatement
}

func (expression *WhileExpression) expressionNode()      {}
func (expression *WhileExpression) TokenLiteral() string { return expression.Token.Literal }
func (expression *WhileExpression) String() string {
	if expression == nil {
		return ""
	}

	var out bytes.Buffer
	out.WriteString(expression.TokenLiteral())

	if expression.Condition != nil {
		out.WriteString("(" + expression.Condition.String() + ")")
	}

	if expression.Body != nil {
		out.WriteString(expression.Body.String())
	}

	return out.String()
}
