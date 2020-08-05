package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tommycuang/cubeprox/entity"
	"github.com/tommycuang/cubeprox/lib/helper"
)

type ErrorConfig struct {
	Prefix string
	Route  string
	Errs   []error
}

func main() {
	wd, _ := os.Getwd()

	configFilesName := helper.GetListOfFile(wd + "/config")
	configs := helper.GetConfigs(wd+"/config/", configFilesName)

	r := gin.Default()
	GenerateRoutes(r, configs, wd+"/fixtures")

	r.Run(":2323")
}

func GenerateRoutes(engine *gin.Engine, configs []entity.Config, jsonRoot string) {
	var errs []ErrorConfig
	for _, config := range configs {
		if cfgErrs := config.Validate(); len(cfgErrs) != 0 {
			errs = append(errs, ErrorConfig{config.Prefix, "", cfgErrs})
			continue
		}

		for _, route := range config.Routes {
			prefix := "/" + config.Prefix
			jsonPath := route.GetJsonPath(jsonRoot + prefix)

			if routesErrs := route.Validate(jsonPath); len(routesErrs) != 0 {
				errs = append(errs, ErrorConfig{config.Prefix, route.Route, routesErrs})
				continue
			}

			url := route.GetURL(prefix)
			jsonData := route.GetJSON(jsonPath)
			AssignRoute(engine, route.Method, url, jsonData, route.ResponseCode)
		}
	}
	printErrors(errs)
}

func AssignRoute(engine *gin.Engine, method, route, jsonData string, responseCode int) {
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonData), &result)

	engine.Handle(method, route, func(c *gin.Context) {
		c.JSON(responseCode, result)
	})
}

func printErrors(errs []ErrorConfig) {
	fmt.Println("====================")
	for _, cfgError := range errs {
		fmt.Printf("Error Config for Prefix: %v\n", cfgError.Prefix)
		if cfgError.Route != "" {
			fmt.Printf("Error for Route: %v\n", cfgError.Route)
		}

		for idx, err := range cfgError.Errs {
			fmt.Printf("\t%d. %v\n", idx+1, err.Error())
		}
		fmt.Println()
	}
}
