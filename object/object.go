package object

import "fmt"

type ObjectType string

const (
	ObjectInteger = "Integer"
	ObjectBoolean = "Boolean"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

/* --- Integer Object ------------------------------------------------------- */

type IntegerObject struct {
	Value int64
}

func (obj *IntegerObject) Type() ObjectType {
	return ObjectInteger
}

func (obj *IntegerObject) Inspect() string {
	return fmt.Sprintf("%d", obj.Value)
}

/* --- Boolean Object ------------------------------------------------------- */

type BooleanObject struct {
	Value bool
}

func (obj *BooleanObject) Type() ObjectType {
	return ObjectBoolean
}

func (obj *BooleanObject) Inspect() string {
	if obj.Value {
		return "true"
	} else {
		return "false"
	}
}
