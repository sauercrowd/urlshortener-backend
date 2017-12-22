package main

import (
	"log"

	"github.com/sauercrowd/urlshortener-backend/pkg/flags"
	"github.com/sauercrowd/urlshortener-backend/pkg/persistence"
	"github.com/sauercrowd/urlshortener-backend/pkg/web"
)

const basename = "http://example.org"

func main() {
	f := flags.Parse()
	storage, err := persistence.Create(f)
	if err != nil {
		log.Fatal("Could not connect to redis:", err)
	}
	ctx := web.Context{
		Redis:    storage,
		Basename: basename,
	}
	web.Handle(":8080", &ctx)
}
