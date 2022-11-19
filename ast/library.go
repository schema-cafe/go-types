package ast

type Library struct {
	Types    map[string]*Type       `json:"types"`
	Funcs    map[string]*Function   `json:"funcs"`
	Vars     map[string]*Expression `json:"vars"`
	Children map[string]*Library    `json:"children"`
}
