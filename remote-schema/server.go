package main

import (
	"log"
	"net/http"
	"os"
	"remote-schema/graph/generated"
	"remote-schema/graph/resolver"
	"remote-schema/pkg/config"
	"remote-schema/pkg/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cfg, err := config.Environ()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.InitWithDSN(cfg.Database.URL)
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		DB: conn,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
