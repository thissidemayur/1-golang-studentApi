package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct{
	Addr string 
}

type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true" env-default:"production" `
	StoragePath string `yaml:"storage_path"  env:"STORAGE_PATH" env-required:"true" `
	HTTPServer  `yaml:"http_server"`
}

// MustLoad loads the configuration from the specified file or environment variables
func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		// check path in flags
		flags:=flag.String("config","","path to configure value")
		flag.Parse()

		configPath = *flags

		// if still empty throw error
		if configPath == "" {
			log.Fatalln("Config file is not sent")
		}
	}

	// if we get path then check if file exists
	if _,err :=os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exists in path: %s \n",configPath)
	}

	var cfg Config
	// load config
	if err:=cleanenv.ReadConfig(configPath,&cfg); err != nil{
		log.Fatalf("Cannot load config from file: %s , error: %v \n",configPath,err.Error())
	}

	return &cfg

}