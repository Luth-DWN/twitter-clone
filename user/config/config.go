package config

import "gopkg.in/yaml.v3"

type Clients struct {
	Postgres `yaml:"postgres"`
	ElasticSearch `yaml:"elasticsearch"`
}

type App struct {
	Clients
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func ReadYAMLFile(path string) (App, error) {
	var app App

	err := yaml.Unmarshal([]byte(path), &app)
	if err != nil {
		return App{}, err
	}

	return app, nil
}
