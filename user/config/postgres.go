package config

import "fmt"

type Postgres struct {
	DBName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func (p *Postgres) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", p.Host, p.Port, p.User, p.DBName, p.Password)
}
