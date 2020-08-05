package entity

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

var (
	errRoutesRouteEmpty        = errors.New("configuration routes route cannot be empty")
	errRoutesMethodNotValid    = errors.New("configuration routes method not valid")
	errRoutesJSONEmpty         = errors.New("configuration routes json cannot be empty")
	errRoutesResponseCodeEmpty = errors.New("configuration routes response code must be 100 - 999")
	errRoutesJSONEmptyFile     = errors.New("configuration routes json file cannot be empty")
)

type Route struct {
	Route        string `yaml:"route"`
	Method       string `yaml:"method"`
	Json         string `yaml:"json"`
	ResponseCode int    `yaml:"response_code"`
}

func (rt *Route) GetURL(prefix string) string {
	return prefix + "/" + rt.Route
}

func (rt *Route) GetJsonPath(jsonRoot string) string {
	return jsonRoot + "/" + rt.Json
}

func (rt *Route) GetJSON(absPath string) string {
	file, _ := os.Open(absPath)
	defer file.Close()
	jsonFile, _ := ioutil.ReadAll(file)
	return string(jsonFile)
}

func (rt *Route) Validate(jsonPath string) []error {
	var errs []error

	if rt.Route == "" {
		errs = append(errs, errRoutesRouteEmpty)
	}
	if !rt.ValidMethod() {
		errs = append(errs, errRoutesMethodNotValid)
	}
	if rt.Json == "" {
		errs = append(errs, errRoutesJSONEmpty)
	}
	if !rt.ValidResponseCode() {
		errs = append(errs, errRoutesResponseCodeEmpty)
	}
	if rt.GetJSON(jsonPath) == "" {
		errs = append(errs, errRoutesJSONEmptyFile)
	}

	return errs
}

func (rt *Route) ValidMethod() bool {
	availMethods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	for _, method := range availMethods {
		if method == strings.ToUpper(rt.Method) {
			return true
		}
	}
	return false
}

func (rt *Route) ValidResponseCode() bool {
	if rt.ResponseCode >= 100 && rt.ResponseCode <= 999 {
		return true
	}
	return false
}
