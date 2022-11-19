package form

import "github.com/schema-cafe/go-types/ast"

type Field struct {
	Name string   `json:"name"`
	Type ast.Type `json:"type"`
}
