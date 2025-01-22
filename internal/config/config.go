package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// Читает файл с конфига, создаст и заполнит объект Config
func MustLoad() *Config {

	// Не видит файл с конфигом и роняет приложение
	// // Файл с конфигом будет считываться из переменной окружения
	// configPath := os.Getenv("CONFIG_PATH")
	// // Если файл с конфигом не найден, то приложение упадет с фаталом
	// if configPath == "" {
	// 	log.Fatal("CONFIG_PATH is not set")
	// }

	// Хардкод, но что поделать
	configPath := "./config/local.yaml"

	// Проверим: существует ли такой файл
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	// Объявляем объект конфига
	var cfg Config

	// Считываем файл по пути, который указали выше и проверяем на ошибки
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	// Возвращаем ссылку на конфиг
	return &cfg
}
