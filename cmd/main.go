package main

import (
	"scoi-web/internal/config"
)

func main() {
	cfg := config.MustLoad()
	_ = cfg
}
