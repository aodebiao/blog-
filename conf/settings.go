package conf

import (
	"fmt"
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
	fmt.Print("settings",Setting)
}


var Setting Settings
type  Settings struct {
	Db DB `yaml:"db"`
}

type DB struct {
	Host string `yaml:"host"`
	Type string `yaml:"type"`
	Port string `yaml:"port"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Name string `yaml:"name"`
}