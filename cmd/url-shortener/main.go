package main

import (
	"fmt"
	"github.com/mchekalov/urlshortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
