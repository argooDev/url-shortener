package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	// Получаем содержимое конфига
	cfg := config.MustLoad()
	fmt.Println(cfg)

}
