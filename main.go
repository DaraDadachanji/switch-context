package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

func main() {
	profiles, err := ReadConfigFile()
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) != 1 {
		log.Fatal("expected 1 argument")
	}
	name := os.Args[0]

	if profile, ok := profiles[name]; ok {
		for key, value := range profile {
			os.Setenv(key, value)
		}
	} else {
		fmt.Println("Profile not found")
		fmt.Println("Available profiles are:")
		for key, _ := range profiles {
			fmt.Println(key)
		}
	}

}

type Profile map[string]string
type Profiles map[string]Profile

func ReadConfigFile() (Profiles, error) {
	configFile := filepath.Join(HomeDirectory(), ".scontext/profiles.yaml")

	if fileExists(configFile) {
		data, err := os.ReadFile(configFile)
		if err != nil {
			log.Fatalln("could not read config file: ", err)
		}
		type Config struct {
			Profiles Profiles `yaml:"profiles"`
		}
		var config Config
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			return nil, err
		}
		return config.Profiles, nil
	}
	log.Println("file not found:", configFile)
	return nil, fmt.Errorf("file not found: %s", configFile)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) || info.IsDir() {
		return false
	}
	return true
}

func HomeDirectory() string {
	u, _ := user.Current()
	return u.HomeDir
}
