package cfg

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Auth string
}

// Read read a cfg file from $HOME directory
func Read() Config {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Can not read files from home directory")
	}

	filePath := home + "/.wbcfg.toml"
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("The config file not found, please add a .wbcfg.toml in $HOME directory")
	}

	doc := string(fileBytes)
	var cfg Config
	err = toml.Unmarshal([]byte(doc), &cfg)
	if err != nil {
		log.Fatal("Parse config file failed, please check format: ", err)
	}

	return cfg
}
