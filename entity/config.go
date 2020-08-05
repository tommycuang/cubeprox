package entity

import (
	"errors"
)

var (
	errConfigPrefix = errors.New("configuration prefix cannot be empty")
	errConfigRoutes = errors.New("configuration routes cannot be empty")
)

type Config struct {
	Prefix string  `yaml:"prefix"`
	Routes []Route `yaml:"routes"`
}

func (cf *Config) Validate() []error {
	var errs []error

	if cf.Prefix == "" {
		errs = append(errs, errConfigPrefix)
	}
	if len(cf.Routes) == 0 {
		errs = append(errs, errConfigRoutes)
	}

	return errs
}
