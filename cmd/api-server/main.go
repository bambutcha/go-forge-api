package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	apiserver "github.com/bambutcha/http-rest-api/internal/app/api-server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/api-server.toml", "config file path")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}

