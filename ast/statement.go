package ast

type Statement struct {
	Type   string              `json:"type"`
	Assign *VariableAssignment `json:"assign"`
	Expr   *Expression         `json:"expr"`
	If     *IfStatement        `json:"if"`
	For    *ForLoop            `json:"for"`
	Return *Expression         `json:"return"`
}
