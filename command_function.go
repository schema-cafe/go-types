package types

type CommandFunction func(stateDir string, inputs map[string]string) (Mutation, error)
