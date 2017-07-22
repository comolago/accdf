package main

import (
	"domain"
	"fmt"
	"infrastructure"
	"interfaces"
	"os"
)

func main() {
	// Benchmarks Library
	var elasticSearch = new(infrastructure.ElasticsearchStore)
	elasticSearch.Index = "accdf"
	elasticSearch.URL = "http://127.0.0.1:9200"

	// Define a configuration handler, that is a stuct with
	// interfaces
	config := new(interfaces.Config)
	// Add available repositories to the configuration handler
	config.Repositories = interfaces.CreateRepositoriesMap()
	config.Repositories.AddRepo(elasticSearch, "Benchmarks")
	config.Repositories.AddRepo(elasticSearch, "Testcases")

	// define a filter
	/*var filter domain.Benchmark
	filter.Name = "checkconn"
	filter.AddPlatform("rhel", "Red Hat Enterprise Linux", "6.x")
	filter.AddPlatform("rhel", "Red Hat Enterprise Linux", "7.x")*/

	var filter domain.TestCase
	filter.Name = "LDAP Servers"
	filter.AddTest("ldap01.carcano.local", "ldap01.carcano.local")

	// performs a lookup using a filter
	//if err := config.Repositories.Lookup(filter, "Benchmarks"); err != nil {
	if err := config.Repositories.Lookup(filter, "Testcases"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
