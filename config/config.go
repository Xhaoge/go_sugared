package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

var setting *Conf

type Conf struct {
	Server *ServerModel `yaml:"server"`
	Mongo  *Mongo       `yaml:"mongo"`
	Path   *Path        `yaml:"path"`
}

type ServerModel struct {
	Mode string `yaml:"mode"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Mongo struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pwd  string `yaml:"password"`
	Port string `yaml:"port"`
}

type Path struct {
	LogPath    string `yaml:"logPath"`
	StaticPath string `yaml:"staticPath"`
}

//LoadConfigInformation load config information for application
func init() {
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
	setting = configApplication
}

func ConfigGetServicePort() string {
	return setting.Server.Port
}

func ConfigGetLoggingPath() string {
	return setting.Path.LogPath
}

func ConfigGetStaticPath() string {
	return setting.Path.StaticPath
}

func ConfigGetMongo() Mongo {
	return Mongo{
		Host: setting.Mongo.Host,
		User: setting.Mongo.User,
		Pwd:  setting.Mongo.Pwd,
		Port: setting.Mongo.Port,
	}
}
