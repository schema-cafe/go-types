package types

type Schema struct {
	Description string  `json:"description"`
	Fields      []Field `json:"fields"`
}
