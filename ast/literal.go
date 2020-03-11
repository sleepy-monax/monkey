package ast

import (
	"bytes"
	"fmt"
	"monkey/token"
)

/* --- Identifier Literal --------------------------------------------------- */

type IdentifierLiteral struct {
	Token token.Token
	Value string
}

func (expression *IdentifierLiteral) expressionNode()      {}
func (expression *IdentifierLiteral) TokenLiteral() string { return expression.Token.Literal }
func (expression *IdentifierLiteral) String() string {
	if expression == nil {
		return ""
	}

	return expression.Value
}

/* --- Boolean Literal ------------------------------------------------------ */

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (expression *BooleanLiteral) expressionNode()      {}
func (expression *BooleanLiteral) TokenLiteral() string { return expression.Token.Literal }
func (expression *BooleanLiteral) String() string {
	if expression == nil {
		return ""
	}

	if expression.Value {
		return "true"
	} else {
		return "false"
	}
}

/* --- Integer Literal ------------------------------------------------------ */

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (expression *IntegerLiteral) expressionNode()      {}
func (expression *IntegerLiteral) TokenLiteral() string { return expression.Token.Literal }
func (expression *IntegerLiteral) String() string {
	if expression == nil {
		return ""
	}

	return fmt.Sprintf("%d", expression.Value)
}

/* --- Function Literal ----------------------------------------------------- */

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*IdentifierLiteral
	Body       *BlockStatement
}

func (expression *FunctionLiteral) expressionNode()      {}
func (expression *FunctionLiteral) TokenLiteral() string { return expression.Token.Literal }
func (expression *FunctionLiteral) String() string {
	if expression == nil {
		return ""
	}

	var out bytes.Buffer
	out.WriteString(expression.TokenLiteral())

	out.WriteString("(")

	for i, param := range expression.Parameters {
		out.WriteString(param.String())

		if i < len(expression.Parameters)-1 {
			out.WriteString(",")
		}
	}

	out.WriteString(")")

	if expression.Body != nil {
		out.WriteString(expression.Body.String())
	}

	return out.String()
}
