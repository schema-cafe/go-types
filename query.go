package types

type Query struct {
	Inputs []Field
	Impl   func(inputs map[string]string) (any, error)
}
