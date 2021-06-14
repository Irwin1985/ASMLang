package code

type Instructions []int
type OpCode int

const (
	push int = iota
	pop
	add
	sub
	mul
	div
	mod
	pow
	le
	leq
	ge
	geq
	not
	neg
	and
	or
	jumpf
	jump
	setg
	getg
	setl
	getl
	true
	false
	null
)
