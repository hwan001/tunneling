package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
)

type Config struct {
	ServerCertFile string `json:"server.crt"`
	ServerKeyFile  string `json:"server.key"`
}

func loadConfig(configFile string) (*Config, error) {
	var config Config

	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func watchConfigFile(configFile string, config *Config) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Config file modified:", event.Name)
					updatedConfig, err := loadConfig(configFile)
					if err == nil {
						*config = *updatedConfig
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(configFile)
	if err != nil {
		log.Fatal(err)
	}
}
