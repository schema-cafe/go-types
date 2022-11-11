package types

import "io/fs"

type CommandFunction func(state fs.FS, inputs map[string]string) (Mutation, error)
