package types

// Error types
// types/Errors.go
type ValidationError struct {
	Message string
}


func (e *ValidationError) Error() string {
	return " Validation Error: " + e.Message
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return " Not Found Error: " + e.Message
}

type InvalidInputError struct {
	Message string
}

func (e *InvalidInputError) Error() string {
	return " Invalid Input Error: " + e.Message
}
