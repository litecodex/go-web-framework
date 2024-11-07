package model

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db-name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
