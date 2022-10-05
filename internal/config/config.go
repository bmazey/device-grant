package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Config provides a template for marshalling .yml configuration files
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	OAuth struct {
		RSABits  int    `yaml:"rsa_bits"`
		Audience string `yaml:"audience"`
		Issuer   string `yaml:"issuer"`
		TokenTTL string `yaml:"jwt_ttl"`
		JWKS     string `yaml:"jwks"`
	} `yaml:"oauth"`
	DeviceGrant struct {
		Registration string `yaml:"registration"`
		UserCode     struct {
			TTL    string `yaml:"ttl"`
			Length int    `yaml:"length"`
		} `yaml:"user_code"`
	} `yaml:"device_grant"`
}

// NewConfig takes a .yml filename from the same /config directory, and returns a populated configuration
func NewConfig(s string) Config {
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
