package config

import (
	"os"
	"io/ioutil"
    "encoding/json"
)

type Config struct {
	Database *Database 		`json:"database"`
}

type Database struct {
	Database string 		`json:"database"`
	Username string 		`json:"username"`
	Password string 		`json:"password"`
}

func Load(fileName string) (*Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	var cfg Config
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, err
	}
	
	return &cfg, nil
}
