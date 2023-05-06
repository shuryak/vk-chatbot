package main

import (
	"flag"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/app"
	"github.com/shuryak/vk-chatbot/internal/config"
	"log"
)

func main() {
	var cfg *config.Config
	var err error

	configFilePath := flag.String("config", "", "Path to the config file")
	flag.Parse()
	if *configFilePath == "" {
		cfg, err = config.ParseEnv()
		if err != nil {
			log.Fatalf("config err: %v", err)
		}
		fmt.Println("Config file not provided. Only environment variables are used.")
	} else {
		cfg, err = config.ParseFileAndEnv(*configFilePath)
		if err != nil {
			log.Fatalf("config err: %v", err)
		}
		fmt.Printf("Using config file at: %s\n.", *configFilePath)
	}

	app.Run(cfg)
}
