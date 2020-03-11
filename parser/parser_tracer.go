package parser

import (
	"fmt"
	"strings"
)

const traceIdentPlaceholder string = "\t"

func (parser *Parser) identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, parser.depth-1)
}

func (parser *Parser) tracePrint(fs string) {
	if parser.t != nil {
		parser.t.Logf("%s%s\n", parser.identLevel(), fs)
	}
}

func (parser *Parser) incIdent() { parser.depth = parser.depth + 1 }
func (parser *Parser) decIdent() { parser.depth = parser.depth - 1 }

func (parser *Parser) trace(msg string) string {
	parser.incIdent()
	parser.tracePrint("BEGIN " + msg + fmt.Sprintf("(%s %s) ", parser.currentToken.Type, parser.peekToken.Type))
	return msg
}

func (parser *Parser) untrace(msg string) {
	parser.tracePrint("END " + msg + fmt.Sprintf("(%s %s) ", parser.currentToken.Type, parser.peekToken.Type))
	parser.decIdent()
}
