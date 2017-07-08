package interfaces

// Define a repository handler structure
type Repositories struct {
	dbHandlers map[string]DbHandler
	//benchmarks domain.Benchmarks
	//filter     Filter
}

// Creates a repository handler map
func CreateRepositoriesMap() *Repositories {
	repository := new(Repositories)
	repository.dbHandlers = make(map[string]DbHandler)
	return repository
}

// Interface: Repositories
// Add the supplied DbHandler identified by repo argument to repository
// handler map
func (db *Repositories) AddRepo(dbh DbHandler, repo string) {
	db.dbHandlers[repo] = dbh
}

// Interface: Repositories
// Lookups repository specified by repo argument applying f filter
func (db *Repositories) Lookup(f Filter, repo string) error {
	return db.dbHandlers[repo].Lookup(f) //, repo)
}
