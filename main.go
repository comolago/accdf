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
	var benchmarkLibrary = new(infrastructure.ElasticsearchStore)
	benchmarkLibrary.Index = "accdf"
	benchmarkLibrary.URL = "http://127.0.0.1:9200"

	// Define a configuration handler, that is a stuct with
	// interfaces
	config := new(interfaces.Config)
	// Add available repositories to the configuration handler
	config.Repositories = interfaces.CreateRepositoriesMap()
	config.Repositories.AddRepo(benchmarkLibrary, "Benchmarks")

	// define a filter
	var filter domain.Benchmark
	filter.Name = "checkconn"
	filter.AddPlatform("rhel", "Red Hat Enterprise Linux", "6.x")
	filter.AddPlatform("rhel", "Red Hat Enterprise Linux", "7.x")

	// performs a lookup using a filter
	if err := config.Repositories.Lookup(filter, "Benchmarks"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
