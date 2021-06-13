package lexer

import (
	"ASMLang/token"
	"fmt"
	"os"
)

type Lexer struct {
	pos   int
	chr   byte
	input string
}

func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isAlpha(ch byte) bool {
	return isLetter(ch) || isAlpha(ch) || ch == '_'
}

func NewLexer(input string) *Lexer {
	lex := &Lexer{
		input: input,
		pos:   0,
	}
	lex.chr = lex.input[lex.pos]
	return lex
}

func newToken(tType token.Type, lexeme string) token.Token {
	t := token.Token{Type: tType, Literal: lexeme}
	return t
}

func (l *Lexer) advance() {
	l.pos += 1
	if l.pos >= len(l.input) {
		l.chr = 0
	} else {
		l.chr = l.input[l.pos]
	}
}

func (l *Lexer) skipWhitespace() {
	for isSpace(l.chr) {
		l.advance()
	}
}

func (l *Lexer) skipComment() {
	for l.chr != 0 && l.chr != '\n' {
		l.advance()
	}
}

func (l *Lexer) getNumber() token.Token {
	var lexeme string
	for {
		for l.chr != 0 && isDigit(l.chr) {
			lexeme += string(l.chr)
			l.advance()
		}
		if l.chr == '.' {
			lexeme += "."
			l.advance()
		} else {
			break
		}
	}
	return newToken(token.NUMBER, lexeme)
}

func (l *Lexer) getString() token.Token {
	var lexeme string
	var strDelim = l.chr
	l.advance() // skip the star delimiter
	for l.chr != 0 && l.chr != strDelim {
		lexeme += string(l.chr)
		l.advance()
	}
	l.advance() // skip ending delim
	return newToken(token.STRING, lexeme)
}

func (l *Lexer) getIdent() token.Token {
	var lexeme string
	for l.chr != 0 && isLetter(l.chr) {
		lexeme += string(l.chr)
		l.advance()
	}
	tokenType := token.IsKeyword(lexeme)

	return newToken(tokenType, lexeme)
}

func (l *Lexer) NextToken() token.Token {
	for l.chr != 0 {
		if isSpace(l.chr) {
			l.skipWhitespace()
			continue
		}
		if l.chr == ';' {
			l.skipComment()
			continue
		}

		if isDigit(l.chr) {
			return l.getNumber()
		}
		if l.chr == '"' || l.chr == '\'' {
			return l.getString()
		}
		if isAlpha(l.chr) {
			return l.getIdent()
		}
		fmt.Printf("unknown character %c\n", l.chr)
		os.Exit(1)
	}
	return newToken(token.EOF, "")
}
