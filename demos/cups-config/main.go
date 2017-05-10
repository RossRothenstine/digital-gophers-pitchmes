package main

import (
	"fmt"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	MicroserviceA = "my-microservice"
)

type Config struct {
	MicroserviceA *Microservice
}

func (c *Config) String() string {
	return fmt.Sprintf("[microserviceA=%v]", c.MicroserviceA)
}

type Microservice struct {
	URI          string `envconfig:"MICROSERVICE_A_URI" required:"true"`
	ClientID     string `envconfig:"MICROSERVICE_A_CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"MICROSERVICE_A_CLIENT_SECRET" required:"true"`
	IssuerURI    string `envconfig:"MICROSERVICE_A_ISSUER_URI" required:"true"`
}

func (ms *Microservice) String() string {
	return fmt.Sprintf("[uri=%s,clientID=%d,clientSecret=%s,issuerURI=%s]", ms.URI, ms.ClientID, ms.ClientSecret, ms.IssuerURI)
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
		if c.MicroserviceA, err = loadFromEnvironment(); err != nil {
			return
		}
	} else {
		if c.MicroserviceA, err = readMicroserviceAFromVcap(cfEnv); err != nil {
			return
		}
	}

	return
}

func loadFromEnvironment() (c *Microservice, err error) {
	c = &Microservice{}
	err = envconfig.Process("app", c)
	return
}

func readMicroserviceAFromVcap(cfEnv *cfenv.App) (c *Microservice, err error) {
	c = &Microservice{}
	var svc *cfenv.Service

	if svc, err = cfEnv.Services.WithName(MicroserviceA); err != nil {
		return
	}

	var ok bool

	if c.URI, ok = svc.Credentials["uri"].(string); !ok {
		err = fmt.Errorf("env: Could not find microserviceA uri in vcap services")
		return
	}
	if c.ClientID, ok = svc.Credentials["clientId"].(string); !ok {
		err = fmt.Errorf("env: Could not find microserviceA clientId in vcap services")
		return
	}
	if c.ClientSecret, ok = svc.Credentials["clientSecret"].(string); !ok {
		err = fmt.Errorf("env: Could not find microserviceA clientSecret in vcap services")
		return
	}
	if c.IssuerURI, ok = svc.Credentials["issuerUri"].(string); !ok {
		err = fmt.Errorf("env: Could not find microserviceA issuerUri in vcap services")
		return
	}
	return
}
