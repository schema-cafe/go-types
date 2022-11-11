package types

type CommandFunction struct {
	Inputs []Field
	Impl   func(inputs map[string]string) (Mutation, error)
}
