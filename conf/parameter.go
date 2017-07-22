package conf

import (

)

type Parameter struct {
	Name 		string
	Value 		string
}
func NewParameter(aName, aValue string) *Parameter {
	return &Parameter{Name : aName, Value : aValue}
}