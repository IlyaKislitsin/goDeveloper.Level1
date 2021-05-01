package consoleconfig

import (
	"flag"
	"fmt"
	"net/url"
)

// ConsoleConfig type
type ConsoleConfig struct {
	Name        *string
	Description *string
	Keywords    *string
	Image       *string
	MountDir    *string
	Website     *string
	Repository  *string
}

// Set ConsoleConfig variables from console flags
func (cc *ConsoleConfig) Set() []error {
	cc.Name = flag.String("name", "", "")
	cc.Description = flag.String("desc", "", "")
	cc.Keywords = flag.String("keys", "", "")
	cc.Image = flag.String("img", "", "")
	cc.MountDir = flag.String("mountDir", "", "")
	cc.Website = flag.String("website", "", "")
	cc.Repository = flag.String("repo", "", "")

	flag.Parse()
	return nil
}

// Validate ConsoleConfig variables
func (cc *ConsoleConfig) Validate() []error {
	var errorList []error

	if err := validateURL(*cc.Website, true); err != nil {
		errorList = append(errorList, fmt.Errorf("website validate error: %w", err))
	}

	if err := validateURL(*cc.Repository, true); err != nil {
		errorList = append(errorList, fmt.Errorf("repository validate error: %w", err))
	}

	return errorList
}

// Print ConsoleConfig variables
func (cc ConsoleConfig) Print() {
	fmt.Printf("name: %v\n", *cc.Name)
	fmt.Printf("description: %v\n", *cc.Description)
	fmt.Printf("keywords: %v\n", *cc.Keywords)
	fmt.Printf("image: %v\n", *cc.Image)
	fmt.Printf("mountDir: %v\n", *cc.MountDir)
	fmt.Printf("website: %v\n", *cc.Website)
	fmt.Printf("repository: %v\n", *cc.Repository)
}

func validateURL(strURL string, absolute bool) error {
	if absolute {
		if _, err := url.ParseRequestURI(strURL); err != nil {
			return err
		}
	} else {
		if _, err := url.Parse(strURL); err != nil {
			return err
		}
	}

	return nil
}
