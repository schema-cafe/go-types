package types

import "github.com/schema-cafe/go-types/form"

type Schema struct {
	Description string       `json:"description"`
	Fields      []form.Field `json:"fields"`
}
