package types

type QueryFunction func(stateDir string, inputs map[string]string) (any, error)
