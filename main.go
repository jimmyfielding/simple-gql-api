package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/simple-gql-api/gql"
	"github.com/simple-gql-api/postgres"
	"github.com/simple-gql-api/server"
)

type Config struct {
	ServerAddr string
	DbAddr     string
	TableName  string
	DbName     string
}

func main() {

	if err := startServer(); err != nil {
		log.Fatal(err)
	}

}

func startServer() error {
	router := chi.NewRouter()
	db, err := postgres.New(
		postgres.FormatConnection("localhost", 5432, "", ""),
	)
	if err != nil {
		return err
	}

	defer db.Close()

	rootQuery := gql.NewRoot(db)
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		return err
	}

	s := server.Server{
		GqlSchema: &schema,
	}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Post("/graphql", s.GraphQL())
	if err = http.ListenAndServe(":4000", router); err != nil {
		return err
	}

	return nil
}
