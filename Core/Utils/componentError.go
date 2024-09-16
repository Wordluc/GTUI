package Utils


type CError struct {
	error    []error
	hasError bool
}

func NewError() *CError {
	return &CError{
		error:    make([]error, 0),
		hasError: false,
	}
}

func (e *CError) Error() string {
	var error string
	for i, v := range e.error {
		if i == len(e.error)-1 {
			error += v.Error() + "\n"
			break
		}
		error += v.Error() + "->"
	}
	return error
}

// add an error, return true if there is an error and false if is nil
func (e *CError) Add(err error) bool {
	if err != nil {
		e.hasError = true
		e.error = append(e.error, err)
		return true
	}
	return false
}

func (e *CError) HasError() bool {
	return e.hasError
}

// add an error if the condition is true and return true if the error is not nil
func (e *CError) AddIf(condition bool, error error) bool {
	if condition {
		return e.Add(error)
	}
	return false
}
