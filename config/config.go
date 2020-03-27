package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

var Setting *Conf

type Conf struct {
	Server *serverModel `yaml:"server"`
	Mongo  *mongo       `yaml:"mongo"`
	Log    *log         `yaml:"log"`
}

type serverModel struct {
	Mode string `yaml:mode`
	Host string `yaml:host`
	Port string `yaml:port`
}

type mongo struct {
	Port string `yaml:port`
}

type log struct {
	FilePath string `yaml:filePath`
}

//LoadConfigInformation load config information for application
func InitConfig() {
	filepath, _ := os.Getwd()
	fileFullPath := path.Join(filepath, "\\config\\dev_config.yaml")
	fmt.Println("fileFullPath: ", fileFullPath)
	configData, err := ioutil.ReadFile(fileFullPath)
	if err != nil {
		fmt.Printf(" config file read failed: %s", err)
		os.Exit(-1)
	}
	var configApplication *Conf

	err = yaml.Unmarshal(configData, &configApplication)
	if err != nil {
		fmt.Printf(" config parse failed: %s", err)
		os.Exit(-1)
	}
	Setting = configApplication
}

func ConfigGetServicePort() string {
	return Setting.Server.Port
}

func ConfigGetLoggingFilePath() string {
	fmt.Println("filepath:", Setting.Log.FilePath)
	return Setting.Log.FilePath
}
