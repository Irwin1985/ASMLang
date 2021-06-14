package object

import "fmt"

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	NUMBER  = "NUMBER"
	STRING  = "STRING"
	BOOLEAN = "BOOLEAN"
	NULL    = "NULL"
)

type Number struct {
	Value float32
}

func (n *Number) Type() ObjectType { return NUMBER }
func (n *Number) Inspect() string {
	return fmt.Sprintf("%f", n.Value)
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING }
func (s *String) Inspect() string  { return s.Value }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN }
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
