package ast

import "bytes"

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

func (p *Program) String() string {
	if p == nil {
		return ""
	}

	var out bytes.Buffer

	for _, s := range p.Statements {
		if s != nil {
			out.WriteString(s.String())
			out.WriteString(";")
		}
	}

	return out.String()
}
