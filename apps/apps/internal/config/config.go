package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DataBase struct {
		MaxIdleConns int `yaml:"max_idle_conns"`
		MaxOpenConns int `yaml:"max_open_conns"`
	}
}

func LoadConfig() (*Config, error) {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, fmt.Errorf("設定ファイルを開けませんでした: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("設定ファイルのパースに失敗しました: %w", err)
	}

	return &cfg, err
}
