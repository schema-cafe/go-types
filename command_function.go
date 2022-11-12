package types

import "io/fs"

type CommandFunction func(state fs.ReadDirFS, inputs map[string]string) (Mutation, error)
