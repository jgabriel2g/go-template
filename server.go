package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go-template/src/graph/generated"
	graph "go-template/src/graph/resolvers"
	"go-template/src/infrastructure"
	"go-template/src/persistence/repositories"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func init() {
	godotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// connection with neo4j
	ctx := context.Background()
	neo4jSession := infrastructure.NewNeo4jSession(ctx)
	defer neo4jSession.Close(ctx)

	repository := repositories.Repository{
		Driver: neo4jSession,
	}

	resolver := &graph.Resolver{
		Repository: repository,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:" + port,
		},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
