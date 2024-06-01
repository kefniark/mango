package config

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigAPIServer struct {
	Name string `mapstructure:"name"`
	URL  string `mapstructure:"url"`
}

type ConfigBuildStatic struct {
	Enable bool `mapstructure:"enable"`
}

type ConfigBuildPlatform struct {
	Os   string `mapstructure:"os"`
	Arch string `mapstructure:"arch"`
}

type ConfigBuild struct {
	Static    ConfigBuildStatic     `mapstructure:"static"`
	Platforms []ConfigBuildPlatform `mapstructure:"platforms"`
}

type ConfigAPI struct {
	Servers []ConfigAPIServer `mapstructure:"servers"`
}

type ConfigDatabase struct {
	Engine string `mapstructure:"engine"`
	URL    string `mapstructure:"url"`
}

type ConfigApp struct {
	Title       string         `mapstructure:"title"`
	Description string         `mapstructure:"description"`
	Version     string         `mapstructure:"version"`
	Port        int            `mapstructure:"port"`
	API         ConfigAPI      `mapstructure:"api"`
	Build       ConfigBuild    `mapstructure:"build"`
	Database    ConfigDatabase `mapstructure:"db"`
}

type Config map[string]ConfigApp

func Parse() *Config {
	config := Config{}
	viper.SetConfigName(".mango")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(".mango.yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	// fmt.Println(viper.AllSettings())

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &config
}
