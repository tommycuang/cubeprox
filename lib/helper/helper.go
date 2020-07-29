package helper

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/tommycuang/cubeprox"
	"gopkg.in/yaml.v2"
)

func GetListOfFile(absPath string) []string {
	var configFilesName []string

	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		configFilesName = append(configFilesName, f.Name())
	}

	return configFilesName
}

func GetConfigs(absPath string, configFilesName []string) []cubeprox.Config {
	var configs []cubeprox.Config
	for _, fileName := range configFilesName {
		var config cubeprox.Config
		file, _ := os.Open(absPath + fileName)
		defer file.Close()
		configList, _ := ioutil.ReadAll(file)
		err := yaml.Unmarshal(configList, &config)
		if err != nil {
			panic(err)
		}
		configs = append(configs, config)
	}

	return configs
}

func GetJSON(absPath string) string {
	file, _ := os.Open(absPath)
	defer file.Close()
	jsonFile, _ := ioutil.ReadAll(file)
	return string(jsonFile)
}
