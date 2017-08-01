package interfaces

import (
	"domain"
)

type Filter interface{}
type Document interface{}

// DbHandler should be implemented using the connection library
// of the choosed backend. Foe example a DbHandler implementation
// could be an elasticsearch package, MongoDb or any other kind
// of backend
type DbHandler interface {
	// Lookup into the backend
	// NoSQL: lookups applying the supplied filter
	// SQL: execute the SQL statement supplied as filter
	Open() error
	Lookup(f Filter) ([]domain.Domain, error)
	AddDocument(d Document) error
	DeleteDocumentById(objtype string, id string) error
}

// Methods used to access data are tightly coupled with the chosen backend
// Here we define the interface with the goal of make them
// loosely coupled
type repositories interface {
	Lookup(f Filter, repo string) ([]domain.Domain, error)
	AddDocument(d Document, repo string) error
	DeleteDocumentById(objtype string, id string, repo string) error
	AddRepo(dbh DbHandler, repo string) error
}

type Config struct {
	Repositories repositories
}
