package types

import "io/fs"

type QueryFunction func(state fs.ReadDirFS, inputs map[string]string) (any, error)
