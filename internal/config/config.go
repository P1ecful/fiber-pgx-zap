package config

type Config struct {
	Service ServiceConfig `yaml:"service"`
	Storage StorageConfig `yaml:"storage"`
}

type ServiceConfig struct{}
