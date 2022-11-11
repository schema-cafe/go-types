package types

type Command struct {
	Inputs []Field
	Impl   func(inputs map[string]string) (Mutation, error)
}
