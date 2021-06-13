package repl

import (
	"ASMLang/lexer"
	"ASMLang/token"
	"bufio"
	"fmt"
	"io"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		input := scanner.Text()
		if len(input) <= 0 {
			continue
		}
		l := lexer.NewLexer(input)
		tok := l.NextToken()
		for tok.Type != token.EOF {
			fmt.Printf("type: %s, lexeme: %s\n", tok.Type, tok.Literal)
			tok = l.NextToken()
		}
	}
}
