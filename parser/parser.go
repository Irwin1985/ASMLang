package parser

import (
	"ASMLang/code"
	"ASMLang/lexer"
	"ASMLang/object"
	"ASMLang/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
	constants []object.Object
}

func NewParser(l *lexer.Lexer, constants []object.Object) *Parser {
	p := &Parser{
		l:         l,
		errors:    []string{},
		constants: constants,
	}
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) advance(tType token.Type) {
	if p.curToken.Type == tType {
		p.nextToken()
	} else {
		msg := "expected next token to be %s, got %s instead."
		p.errors = append(p.errors, msg)
	}
}

func (p *Parser) ParseAndCompile() error {
	switch p.curToken.Type {
	case token.PUSH:

	}
	return nil
}

func addConstant(opCode code.OpCode, operands ...int) {

}
