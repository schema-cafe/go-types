package types

type API struct {
	Commands map[string]Command
	Queries  map[string]Query
}
