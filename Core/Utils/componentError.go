package Utils

import (
	"fmt"
)

type CError struct {
	error[]     error
	hasError    bool
}
func (e CError) HasError() bool {
	return e.hasError
}

func (e CError) Error() string {
	var error string
	for i, v := range e.error {
		if i == len(e.error)-1 {
			error += v.Error() + "\n"
			break
		}
		error += v.Error() +"->"
	}
	return fmt.Sprint(e.error)
}
//add an error, return true if there is an error and false if is nil
func (e CError) Add(error error)bool {
	if error == nil {
		return false
	}
	e.error = append(e.error, error)
	e.hasError = true
	return true
}
//add an error is the condition is true
func (e CError) AddIf(condition bool, error error)bool {
	if condition {
		return e.Add(error)
	}
	return false
}
