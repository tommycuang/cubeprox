package helper

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/tommycuang/cubeprox/entity"
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

func GetConfigs(absPath string, configFilesName []string) []entity.Config {
	var configs []entity.Config
	for _, fileName := range configFilesName {
		var config entity.Config
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
