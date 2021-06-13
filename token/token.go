package token

/*
* ; this is a comment
*
* loadc contantIndex 	; loads a constant in stack
* store identifier   	; pops from stack and store in global array.
* add					; pops 2 values from stack and adds them
* sub					; pops 2 values from stack and subs them
* mul					; pops 2 values from stack and mul them
* div					; pops 2 values from stack and div them
* mod					; pops 2 values from stack and perform a modulus operation
* jumpf	index			; jumps to index if the top most value of the stack is false
* jump index			; jumps to index
*
 */

type Type string

const (
	EOF = "EOF"

	// stack operations
	PUSH = "PUSH"
	POP  = "POP"

	IDENT  = "IDENT"
	NUMBER = "NUMBER"
	STRING = "STRING"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	NULL   = "NULL"

	// arithmetic operators
	ADD = "ADD"
	SUB = "SUB"
	MUL = "MUL"
	DIV = "DIV"
	MOD = "MOD"
	POW = "POW"

	// relational operators
	LE  = "LE"
	LEQ = "LEQ"
	GE  = "GE"
	GEQ = "GEQ"
	EQ  = "EQ"
	NEQ = "NEQ"

	// logical operators
	NOT = "NOT"
	NEG = "NEG"
	AND = "AND"
	OR  = "OR"

	JUMPF = "JUMPF"
	JUMP  = "JUMP"

	SETG = "SETG"
	GETG = "GETG"
	SETL = "SETL"
	GETL = "GETL"
)

type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"push":  PUSH,
	"pop":   POP,
	"add":   ADD,
	"sub":   SUB,
	"mul":   MUL,
	"div":   DIV,
	"mod":   MOD,
	"pow":   POW,
	"le":    LE,
	"leq":   LEQ,
	"ge":    GE,
	"geq":   GEQ,
	"not":   NOT,
	"neg":   NEG,
	"and":   AND,
	"or":    OR,
	"jumpf": JUMPF,
	"jump":  JUMP,
	"setg":  SETG,
	"getg":  GETG,
	"setl":  SETL,
	"getl":  GETL,
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

func IsKeyword(ident string) Type {
	if keyword, ok := keywords[ident]; ok {
		return keyword
	}
	return IDENT
}
