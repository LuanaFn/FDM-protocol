package configs

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	"gopkg.in/yaml.v3"
)

var Config Configuration

type Configuration struct {
	Business struct {
		Endpoints map[string]string `yaml:"endpoints"`
	} `yaml:"business"`
	ServiceHost string
}

func (c *Configuration) Load(environment string) error {
	path := "configs/config.yml"
	if len(environment) > 0 {
		path = fmt.Sprintf("%s_%s", environment, path)
	}
	f, err := os.Open("configs/config.yml")
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Println("error closing config yml: ", err.Error())
		}
	}()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		return err
	}
	return nil
}
