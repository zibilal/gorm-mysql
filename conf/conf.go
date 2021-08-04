package conf

import (
	"bytes"
	"io"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Host string `yaml:"host"`
	Db struct {
		Name string `yaml:"name"`
		Hostname string `yaml:"host"`
		Port int `yaml:"port"`
		User string `yaml:"user"`
		Password string `yaml:"pswd"`
	} `yaml:"dbengine"`
}

func GetConfiguration(reader io.Reader) (*Configuration, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)

	if err != nil {
		return nil, err
	}

	c := new(Configuration)

	err = yaml.Unmarshal(buf.Bytes(), &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
