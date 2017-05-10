package main

import (
	"fmt"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	EnvVarRedisHost = "REDIS_HOST"
	EnvVarRedisPort = "REDIS_PORT"

	Redis = "redis"
)

type Config struct {
	Redis *RedisConfig
}

func (c *Config) String() string {
	return fmt.Sprintf("[redis=%v]", c.Redis)
}

type RedisConfig struct {
	Host     string `envconfig:"REDIS_HOST" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD" required:"true"`
	Port     int    `envconfig:"REDIS_PORT" required:"true"`
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
	for {
	}
}

func initConfig() (c *Config, err error) {
	c = &Config{}

	var cfEnv *cfenv.App
	if cfEnv, err = cfenv.Current(); err != nil {
		if c.Redis, err = loadFromEnvironment(); err != nil {
			return
		}
	} else {
		if c.Redis, err = readRedisFromVcap(cfEnv); err != nil {
			return
		}
	}

	return
}

func loadFromEnvironment() (c *RedisConfig, err error) {
	c = &RedisConfig{}
	err = envconfig.Process("app", c)
	return
}

func readRedisFromVcap(cfEnv *cfenv.App) (c *RedisConfig, err error) {
	c = &RedisConfig{}
	var services []cfenv.Service

	if services, err = cfEnv.Services.WithTag(Redis); err != nil {
		return
	}

	svc := services[0]
	var ok bool

	if c.Host, ok = svc.Credentials["host"].(string); !ok {
		err = fmt.Errorf("env: Could not find redis host in vcap services")
		return
	}
	if c.Password, ok = svc.Credentials["password"].(string); !ok {
		err = fmt.Errorf("env: Could not find redis password in vcap services")
		return
	}
	if c.Port, ok = svc.Credentials["port"].(int); !ok {
		err = fmt.Errorf("env: Could not find redis port in vcap services")
		return
	}
	return
}
