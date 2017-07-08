package interfaces

type Filter interface{}

// DbHandler should be implemented using the connection library
// of the choosed backend. Foe example a DbHandler implementation
// could be an elasticsearch package, MongoDb or any other kind
// of backend
type DbHandler interface {
	// Lookup into the backend
	// NoSQL: lookups applying the supplied filter
	// SQL: execute the SQL statement supplied as filter
	Lookup(f Filter) error
}

// Methods used to access data are tightly coupled with the chosen backend
// Here we define the interface with the goal of make them
// loosely coupled
type repositories interface {
	Lookup(f Filter, repo string) error
	AddRepo(dbh DbHandler, repo string)
}

type Config struct {
	Repositories repositories
}
