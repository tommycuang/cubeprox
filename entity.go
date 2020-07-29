package cubeprox

type Config struct {
	Prefix string `yaml:"prefix"`
	Routes []struct {
		Route        string `yaml:"route"`
		Method       string `yaml:"method"`
		Json         string `yaml:"json"`
		ResponseCode int    `yaml:"response_code"`
	} `yaml:"routes"`
}
