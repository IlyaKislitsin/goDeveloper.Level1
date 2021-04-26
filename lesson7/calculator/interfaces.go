package calculator

// Calculator interface for expressions
type Calculator interface {
	Validate() []error
	Sum() (interface{}, error)
	Subtract() (interface{}, error)
	Multiply() (interface{}, error)
	Divide() (interface{}, error)
	Calculate() (interface{}, error)
}
