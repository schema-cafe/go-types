package ast

type VariableAssignment struct {
	ID    *Identifier `json:"id"`
	Value *Expression `json:"value"`
}
