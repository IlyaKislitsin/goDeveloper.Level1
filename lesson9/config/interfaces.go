package config

// Configurer interface
type Configurer interface {
	Validate() []error
	Set() []error
	Print()
}
