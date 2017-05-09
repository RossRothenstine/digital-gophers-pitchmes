package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EnvVarRedisHost = "REDIS_HOST"
	EnvVarRedisPort = "REDIS_PORT"
)

type Config struct {
	Redis *RedisConfig
}

func (c *Config) String() string {
	return fmt.Sprintf("[redis=%v]", c.Redis)
}

type RedisConfig struct {
	Host string
	Port int
}

func (rc *RedisConfig) String() string {
	return fmt.Sprintf("[host=%s,port=%d]", rc.Host, rc.Port)
}

func main() {
	config, err := initConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", config)
}

func initConfig() (c *Config, err error) {
	var ok bool
	c = &Config{
		Redis: &RedisConfig{},
	}

	if c.Redis.Host, ok = os.LookupEnv(EnvVarRedisHost); !ok {
		err = formatConfigError(EnvVarRedisHost)
		return
	}

	if c.Redis.Port, ok = lookupEnvInt(EnvVarRedisPort); !ok {
		err = formatConfigError(EnvVarRedisPort)
		return
	}

	return
}

func lookupEnvInt(envvar string) (i int, ok bool) {
	var value string
	var err error

	if value, ok = os.LookupEnv(envvar); !ok {
		return
	}

	if i, err = strconv.Atoi(value); err != nil {
		panic(err)
	}
	return
}

func formatConfigError(envvar string) (err error) {
	return fmt.Errorf("%s is not defined in the environment.", envvar)
}
