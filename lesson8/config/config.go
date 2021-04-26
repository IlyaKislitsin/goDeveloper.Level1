package config

import (
	"flag"
	"fmt"
)

// Config type
type Config struct {
	FloatPrecision *int64
	Unsigned       *bool
	Greeting       *string
}

// SetFlags set Config variables
func (c *Config) SetFlags() {
	c.FloatPrecision = flag.Int64("precision", 2, "Колличество цифр после запятой")
	c.Unsigned = flag.Bool("unsigned", false, "Разрешить только положительные числа")
	c.Greeting = flag.String("greeting", "", "Приветственное слово")
}

// Print Config variables
func (c *Config) Print() {
	fmt.Printf("precision: %d\n", *c.FloatPrecision)
	fmt.Printf("unsigned: %v\n", *c.Unsigned)
	fmt.Printf("greeting: %s\n", *c.Greeting)
}

// New creates and populates Config struct
func New() (Config, error) {
	config := Config{}
	config.SetFlags()

	flag.Parse()
	return config, nil
}
