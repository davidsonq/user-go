package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type configs struct {
	APIconfigs APIconfigs
	DBconfigs  DBconfigs
}
type APIconfigs struct {
	Port string
}
type DBconfigs struct {
	Port string
	User string
	Host string
	Pass string
	Name string
}

func newConfig() *configs {
	return &configs{}
}

func GetConfig() *configs {
	godotenv.Load(".env")

	cfg := newConfig()

	cfg.APIconfigs = APIconfigs{
		Port: os.Getenv("API_PORT"),
	}

	cfg.DBconfigs = DBconfigs{
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Host: os.Getenv("DB_HOST"),
		Name: os.Getenv("DB"),
		Pass: os.Getenv("DB_PASS"),
	}

	return cfg
}
