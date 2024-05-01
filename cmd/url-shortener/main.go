package main

import (
	"fmt"
	"urlshortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
