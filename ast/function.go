package ast

type Function struct {
	InputType  *Type       `json:"inputType"`
	OutputType *Type       `json:"outputType"`
	Body       []Statement `json:"body"`
}
