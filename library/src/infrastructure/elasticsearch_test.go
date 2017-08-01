package infrastructure

import (
	"domain"
	"fmt"
	"interfaces"
	"testing"
)

func TestLookup(t *testing.T) {
	var elasticSearch = new(ElasticsearchStore)
	elasticSearch.Index = "accdf"
	elasticSearch.URL = "http://127.0.0.1:9200"
	config := new(interfaces.Config)
	// Add available repositories to the configuration handler
	config.Repositories = interfaces.CreateRepositoriesMap()
	config.Repositories.AddRepo(elasticSearch, "Benchmarks")

	var filter domain.Benchmark
	filter.Name = "checkconn"
	filter.AddPlatform("rhel", "Red Hat Enterprise Linux", "6.x")
	filter.AddPlatform("rhel", "Red Hat Enterprise Linux", "7.x")
	if values, err := config.Repositories.Lookup(filter, "Benchmarks"); err != nil {
		fmt.Println(err)
		t.Error(`(*ElasticsearchStore) Lookup(filter, "Benchmarks") = false`)
	} else {
		if values[0].GetId() != "000000000" {
			t.Error(`(*ElasticsearchStore) Lookup(filter, "Benchmarks") = false: returned Id should be 000000000`)
		}
	}
}
