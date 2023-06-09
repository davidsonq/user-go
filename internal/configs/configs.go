package configs

import (
	"log"
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
	Ssl  string
}

func newConfig() *configs {
	return &configs{}
}

func GetConfig() *configs {

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
		Ssl:  os.Getenv("SSL"),
	}

	return cfg
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}
}
