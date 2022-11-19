package ui

type Element struct {
	Type     string            `json:"type"`
	Attrs    map[string]string `json:"attrs"`
	Children []Element         `json:"children"`
}
