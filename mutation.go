package types

type Mutation struct {
	Path     string `json:"path"`
	Contents string `json:"contents"`
}
