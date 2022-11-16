package types

type QueryResult struct {
	Result any   `json:"result"`
	Error  error `json:"error"`
}
