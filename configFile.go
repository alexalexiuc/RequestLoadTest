package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// type yaml struct {
// }

// func (u *yaml) Unmarshal(in []byte, out interface{}) (error) {
// 	return yamlV2.Unmarshal(in, out)
// }

type ConfigCredentials struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

type ConfigUser struct {
	Credentials ConfigCredentials `yaml:"credentials"`
	LoginUrl    string            `yaml:"login_url"`
}

type ConfigFileStruct struct {
	User ConfigUser `yaml:"user"`
}

func readConfigFile(fPath string) ([]byte, LocalError) {
	data, err := ioutil.ReadFile(fPath)
	if err != nil {
		return nil, ConfigFileReadErr.WithError(err)
	}
	return data, nil
}

func LoadConfigFile() (ConfigFileStruct, LocalError) {
	data, lErr := readConfigFile("./config.yaml")
	if lErr != nil {
		return ConfigFileStruct{}, lErr
	}
	output := ConfigFileStruct{}
	err := yaml.Unmarshal(data, &output)
	if err != nil {
		return ConfigFileStruct{}, ConfigFileUnmarshalErr.WithError(err)
	}
	return output, nil
}
