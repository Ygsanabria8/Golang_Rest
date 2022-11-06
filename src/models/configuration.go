package models

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Mongo struct {
		ConnectionString string `yaml:"connectionString"`
		Database         string `yaml:"database"`
		Users            string `yaml:"users"`
		Tweet            string `yaml:"tweet"`
	} `yaml:"mongo"`
	Jwt struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
	Costencrypt int `yaml:"costencrypt"`
}
