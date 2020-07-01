package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Database *Database `json:"database"`
	Token    *Token    `json:"token"`
	Server   *Server   `json:"server"`
}

type Database struct {
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Key string `json:"key"`
}

type Server struct {
	Port string `json:"port"`
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
