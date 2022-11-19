package ui

type Element struct {
	Type     string            `json:"type"`
	Attrs    map[string]string `json:"Attrs"`
	Children []Element         `json:"children"`
}
