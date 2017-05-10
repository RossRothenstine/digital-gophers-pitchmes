package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var (
	configFile = flag.String("config", "config.yml", "YAML Configuration File.")
)

type Config struct {
	Redis *RedisConfig `yaml:"redis"`
}

func (c *Config) String() string {
	return fmt.Sprintf("[redis=%v]", c.Redis)
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

func (rc *RedisConfig) String() string {
	return fmt.Sprintf("[host=%s,port=%d,password=%s]", rc.Host, rc.Port, rc.Password)
}

func main() {
	config, err := initConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", config)
}

func initConfig() (c *Config, err error) {
	c = &Config{}
	var bytes []byte

	if bytes, err = readAll(*configFile); err != nil {
		return
	}

	err = yaml.Unmarshal(bytes, &c)
	return
}

func readAll(file string) (bytes []byte, err error) {
	var f *os.File
	if f, err = os.Open(file); err != nil {
		return
	}
	bytes, err = ioutil.ReadAll(f)
	return
}
