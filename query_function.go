package types

type QueryFunction struct {
	Inputs []Field
	Impl   func(inputs map[string]string) (any, error)
}
