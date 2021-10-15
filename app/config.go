package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HttpServer  HttpServerConfig
	MysqlConfig MysqlConfig
	Swagger     SwaggerConfig
}

type MysqlConfig struct {
	DatabaseName    string `env:"DB_NAME" default:"sys"`
	User            string `env:"DB_USER" default:"root"`
	Password        string `env:"DB_PASS" default:"123456"`
	Host            string `env:"DB_HOST" default:"127.0.0.1"`
	Port            string `env:"DB_PORT" default:"3306"`
	MultiStatements string `env:"MULTI_STATEMENTS" default:"true"`
}

type HttpServerConfig struct {
	Port            int           `env:"HTTP_PORT" default:"8080"`
	ShutdownTimeout time.Duration `env:"HTTP_SHUTDOWN_TIMEOUT" default:"1s"`
	ReadTimeout     time.Duration `env:"HTTP_READ_TIMEOUT" default:"30s"`
	WriteTimeout    time.Duration `env:"HTTP_WRITE_TIMEOUT" default:"10s"`
	AccessSecret    string        `env:"ACCESS_SECRET" default:"asdhjkasjheee"`
}

type SwaggerConfig struct {
	SwaggerHost string `env: "SWAGGER_HOST" default:"localhost:8080/docs/swagger"`
}

func ReadConfigFromEnv() *Config {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal("Error Read env file")
	}

	return &cfg
}

func (mysql MysqlConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?multiStatements=%s",
		mysql.User,
		mysql.Password,
		mysql.Host,
		mysql.DatabaseName,
		mysql.MultiStatements,
	)

}

func (mysql MysqlConfig) URL() string {
	//	var connectionString = fmt.Sprint("root:123456@tcp(localhost:3306)/challengego?multiStatements=true")
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=%s",
		mysql.User, mysql.Password, mysql.Host, mysql.DatabaseName, mysql.MultiStatements)

	return connectString
}

func ReadConfigFromFile(filename string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(filename, &cfg)
	if err != nil {
		log.Fatal("error reading file")
	}

	return &cfg
}

func ReadConfig(filename string) *Config {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("File not found %s", filename)
		return ReadConfigFromEnv()
	}

	return ReadConfigFromFile(filename)
}
