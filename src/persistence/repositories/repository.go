package repositories

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

type Repository struct {
	Driver neo4j.SessionWithContext
}
