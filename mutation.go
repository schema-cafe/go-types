package types

type Mutation struct {
	Sets    map[string][]byte
	Deletes []string
}
