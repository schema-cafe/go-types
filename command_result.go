package types

type CommandResult struct {
	Mutation Mutation `json:"mutation"`
	Error    error    `json:"error"`
}
