package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler" // ✅ FIXED
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountUrl string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL string `envconfig:"CATALOG_SERVICE_URL"`
	OrderURL   string `envconfig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	s, err := NewGraphQLServer(cfg.AccountUrl, cfg.CatalogURL, cfg.OrderURL)
	if err != nil {
		log.Fatal(err)
	}

	// ✅ Use new handler.Server
	srv := handler.NewDefaultServer(s.ToExecutableSchema())

	http.Handle("/graphql", srv)
	http.Handle("/playground", playground.Handler("murali", "/graphql"))

	log.Println("GraphQL server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
