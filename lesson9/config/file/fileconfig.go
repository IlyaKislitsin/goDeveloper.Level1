package fileconfig

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	toml "github.com/pelletier/go-toml"
	yaml "gopkg.in/yaml.v3"
)

// FileConfig type
type FileConfig struct {
	Name        string   `json:"name" yaml:"name" toml:"name"`
	Description string   `json:"description" yaml:"description" toml:"description"`
	Keywords    []string `json:"keywords" yaml:"keywords" toml:"keywords"`
	Image       string   `json:"image" yaml:"image" toml:"image"`
	MountDir    string   `json:"mount_dir" yaml:"mount_dir" toml:"mount_dir"`
	Website     string   `json:"website" yaml:"website" toml:"website"`
	Repository  string   `json:"repository" yaml:"repository" toml:"repository"`
}

// Set FileConfig variables from file
func (fc *FileConfig) Set() []error {
	var errorList []error

	configFilename := flag.String("config", "conf.json", "Файл конфигурации(json, yaml, toml)")
	flag.Parse()

	if _, statErr := os.Stat(*configFilename); statErr != nil {
		errorList = append(errorList, errors.New("get file info error"))
		return errorList
	}

	file, fileErr := os.Open(*configFilename)
	if fileErr != nil {
		errorList = append(errorList, errors.New("open file error"))
		return errorList
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("close file error: %v", err)
		}
	}()

	switch strings.ToLower(path.Ext(*configFilename)) {
	case ".json":
		decodeErr := json.NewDecoder(file).Decode(fc)
		if decodeErr != nil {
			errorList = append(errorList, decodeErr)
			return errorList
		}
	case ".yaml", ".yml":
		decodeErr := yaml.NewDecoder(file).Decode(fc)
		if decodeErr != nil {
			errorList = append(errorList, decodeErr)
			return errorList
		}
	case ".toml":
		decodeErr := toml.NewDecoder(file).Decode(fc)
		if decodeErr != nil {
			errorList = append(errorList, decodeErr)
			return errorList
		}
	default:
		errorList = append(errorList, errors.New("unknown config file"))
		return errorList
	}

	return nil
}

// Validate FileConfig variables
func (fc *FileConfig) Validate() []error {
	var errorList []error

	if err := validateURL(fc.Website, true); err != nil {
		errorList = append(errorList, fmt.Errorf("website validate error: %w", err))
	}

	if err := validateURL(fc.Repository, true); err != nil {
		errorList = append(errorList, fmt.Errorf("repository validate error: %w", err))
	}

	return errorList
}

// Print FileConfig variables
func (fc FileConfig) Print() {
	fmt.Printf("name: %s\n", fc.Name)
	fmt.Printf("description: %s\n", fc.Description)
	fmt.Printf("keywords: %v\n", fc.Keywords)
	fmt.Printf("image: %s\n", fc.Image)
	fmt.Printf("mountDir: %s\n", fc.MountDir)
	fmt.Printf("website: %s\n", fc.Website)
	fmt.Printf("repository: %s\n", fc.Repository)
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
