package interfaces

import (
	"domain"
)

// Define a repository handler structure
type Repositories struct {
	dbHandlers map[string]DbHandler
	//benchmarks domain.Benchmarks
	//filter     Filter
}

// Creates a repository handler map
/*func CreateRepositoriesMap() *Repositories {
	repository := new(Repositories)
	repository.dbHandlers = make(map[string]DbHandler)
	return repository
}*/

// Creates a repository handler map
func (db *Repositories) Init() {
	db.dbHandlers = make(map[string]DbHandler)
}

// Interface: Repositories
// Add the supplied DbHandler identified by repo argument to repository
// handler map
func (db *Repositories) AddRepo(dbh DbHandler, repo string) error {
	db.dbHandlers[repo] = dbh
	return db.dbHandlers[repo].Open()
}

// Interface: Repositories
// Lookups repository specified by repo argument applying f filter
func (db *Repositories) Lookup(f Filter, repo string) ([]domain.Domain, error) {
	return db.dbHandlers[repo].Lookup(f)
}

func (db *Repositories) AddDocument(d Document, repo string) error {
	return db.dbHandlers[repo].AddDocument(d)
}

func (db *Repositories) DeleteDocumentById(objtype string, id string, repo string) error {
	return db.dbHandlers[repo].DeleteDocumentById(objtype, id)
}
