package infrastructure

import (
	"domain"
	"fmt"
	"interfaces"
	"os"
	"path/filepath"
	"testing"
)

const URL = "http://127.0.0.1:9200"
const INDEX = "accdf"
const BENCHMARKFILE = "../../test/xml/benchmarks/001.xml"

func TestAddBenchmark(t *testing.T) {
	var elasticSearch = new(ElasticsearchStore)
	elasticSearch.Index = INDEX
	elasticSearch.URL = URL
	config := new(interfaces.Config)
	config.Repositories.Init()
	if err := config.Repositories.AddRepo(elasticSearch, "Benchmarks"); err != nil {
		errmsg := fmt.Sprintf("config.Repositories.AddRepo(elasticSearch, \"Benchmarks\")\n%s", err)
		t.Error(errmsg)
	}
	// Benchmark
	BenchmarkFile, err := filepath.Abs(BENCHMARKFILE)
	if err != nil {
		t.Error(err)

	}
	file, err := os.Open(BenchmarkFile)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	var bench domain.Benchmark
	err = bench.FromXML(file)
	if err != nil {
		t.Error(err)
	}
	if err := config.Repositories.AddDocument(bench, "Benchmarks"); err != nil {
		t.Error(err)
	}
	// Benchmark  *** END
}

func TestLookup(t *testing.T) {
	var elasticSearch = new(ElasticsearchStore)
	elasticSearch.Index = INDEX
	elasticSearch.URL = URL
	config := new(interfaces.Config)
	config.Repositories.Init()
	if err := config.Repositories.AddRepo(elasticSearch, "Benchmarks"); err != nil {
		errmsg := fmt.Sprintf("config.Repositories.AddRepo(elasticSearch, \"Benchmarks\")\n%s", err)
		t.Error(errmsg)
	}
	// Benchmark
	BenchmarkFile, err := filepath.Abs(BENCHMARKFILE)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Open(BenchmarkFile)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	var bench domain.Benchmark
	err = bench.FromXML(file)
	if err != nil {
		t.Error(err)
	}
	var filter domain.Benchmark
	filter.SetName(bench.GetName())
	if values, err := config.Repositories.Lookup(filter, "Benchmarks"); err != nil {
		fmt.Println(err)
		t.Error(`(*ElasticsearchStore) Lookup(filter, "Benchmarks") = false`)
	} else {
		if values != nil {

			if values[0].GetId() != bench.GetId() {
				errmsg := `the Id of the benchmark found on Elasticsearch should should be `
				errmsg += bench.GetId()
				fmt.Println(errmsg)
				t.Error(errmsg)
			} else {
				fmt.Println("We have found the benchmark inserted in previous step. Id is", bench.GetId())
			}

		} else {
			t.Error("unable to find the benchmark we have inserted for the test on elasticsearch!!!")
		}
	}
}

/*
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
}*/
