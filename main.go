package main

import (
	"github.com/jasonyangmh/sayakaya/config"
	"github.com/jasonyangmh/sayakaya/database"
	"github.com/jasonyangmh/sayakaya/router"
)

func main() {
	cfg := config.Load()

	db := database.Connect(cfg)
	db.Migrate()

	r := router.New(db.Database)
	r.Start()
}
