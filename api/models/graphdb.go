package models

import (
	"api/config"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type GraphDb struct {
	Driver   neo4j.Driver
	Database string
}

func NewGraphDb(cfg config.Neo4j) *GraphDb {
	driver, err := neo4j.NewDriver(cfg.Uri, neo4j.BasicAuth(cfg.Username, cfg.Password, ""))
	if err != nil {
		log.Fatalf("error connecting to neo4j %s", err)
	}
	return &GraphDb{
		Driver:   driver,
		Database: cfg.Database}
}

func (n *GraphDb) NewReadSession() neo4j.Session {
	return n.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead, DatabaseName: n.Database})
}

func (n *GraphDb) NewWriteSession() neo4j.Session {
	return n.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite, DatabaseName: n.Database})
}
