package main

import (
	"fmt"
	"log"

	"github.com/IlyaKislitsin/goDeveloper.Level1/lesson9/config"
	consoleconfig "github.com/IlyaKislitsin/goDeveloper.Level1/lesson9/config/console"
	fileconfig "github.com/IlyaKislitsin/goDeveloper.Level1/lesson9/config/file"
)

func main() {
	// consoleConfig()
	fileConfig()
}

func consoleConfig() {
	cc := consoleconfig.ConsoleConfig{}
	if ccErr := buildConfig(&cc); ccErr != nil {
		log.Fatalf("console config errors: %v", ccErr)
	}
	fmt.Println("ConsoleCofig:")
	printConfig(&cc)
}

func fileConfig() {
	fc := fileconfig.FileConfig{}
	if fcErr := buildConfig(&fc); fcErr != nil {
		log.Fatalf("file config errors: %v", fcErr)
	}
	fmt.Println("FileCofig:")
	printConfig(&fc)
}

func buildConfig(c config.Configurer) []error {
	if setErrors := c.Set(); setErrors != nil {
		return setErrors
	}
	if configErrors := c.Validate(); configErrors != nil {
		return configErrors
	}

	return nil
}

func printConfig(c config.Configurer) {
	c.Print()
}
