package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tommycuang/cubeprox"
	"github.com/tommycuang/cubeprox/lib/helper"
)

func main() {
	wd, _ := os.Getwd()

	configFilesName := helper.GetListOfFile(wd + "/config")
	configs := helper.GetConfigs(wd+"/config/", configFilesName)

	r := gin.Default()
	GenerateRoutes(r, configs, wd+"/fixtures/")

	r.Run(":2323")
}

func GenerateRoutes(engine *gin.Engine, configs []cubeprox.Config, jsonRoot string) {
	for _, config := range configs {
		for _, route := range config.Routes {
			url := "/" + config.Prefix + "/" + route.Route
			jsonPath := jsonRoot + "/" + config.Prefix + "/" + route.Json
			AssignRoute(engine, route.Method, url, jsonPath, route.ResponseCode)
		}
	}
}

func AssignRoute(engine *gin.Engine, method, route, jsonPath string, responseCode int) {
	jsonData := helper.GetJSON(jsonPath)
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonData), &result)

	engine.Handle(method, route, func(c *gin.Context) {
		c.JSON(responseCode, result)
	})
}
