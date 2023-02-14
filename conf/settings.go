package conf

import (
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"runtime"
)

func init()  {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename)
	file, err := os.ReadFile(dir +string(os.PathSeparator)+ "application.yml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(file,&Setting)
}


var Setting Settings
type  Settings struct {
	Db DB `yaml:"db"`
	Log Log `yaml:"log"`
}

type DB struct {
	Host string `yaml:"host"`
	Type string `yaml:"type"`
	Port string `yaml:"port"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Name string `yaml:"name"`
}

type Log struct {
	Level int8 `yaml:"level"`
	Path string `yaml:"path"`
}