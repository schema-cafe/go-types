package ast

type Type struct {
	IsArray   bool       `json:"isArray"`
	IsMap     bool       `json:"isMap"`
	IsPointer bool       `json:"isPointer"`
	BaseType  Identifier `json:"baseType"`
}
