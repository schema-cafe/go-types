package ast

type IfStatement struct {
	Condition *Expression `json:"condition"`
	Body      []Statement `json:"body"`
}
