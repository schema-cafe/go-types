package types

type Mutation struct {
	Sets    map[string]string
	Deletes []string
}
