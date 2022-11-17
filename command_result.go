package types

type CommandResult struct {
	Mutations []Mutation `json:"mutation"`
	Error     error      `json:"error"`
}
