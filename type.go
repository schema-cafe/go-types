package types

type Type struct {
	IsArray   bool
	IsMap     bool
	IsPointer bool
	BaseType  Identifier
}
