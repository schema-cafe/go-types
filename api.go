package types

type API struct {
	Commands map[string]CommandFunction
	Queries  map[string]QueryFunction
}
