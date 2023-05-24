package infrastructure

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"os"
)

type config struct {
	Host     string
	User     string
	Password string
}

func NewNeo4jSession(ctx context.Context) neo4j.SessionWithContext {
	host := os.Getenv("NEO4J_HOST")
	db := os.Getenv("NEO4J_DATABASE")

	if db != "" {
		host = fmt.Sprintf("%s?database=%s", host, db)
	}

	data := config{
		Host:     host,
		User:     os.Getenv("NEO4J_USERNAME"),
		Password: os.Getenv("NEO4J_PASSWORD"),
	}
	driver, err := neo4j.NewDriverWithContext(data.Host, neo4j.BasicAuth(data.User, data.Password, ""))
	if err != nil {
		panic("CANT CONNECT TO NEO4J")
	}
	fmt.Println("Connect to Neo4J...")

	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	return session
}
