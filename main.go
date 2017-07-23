package main

import (
	"domain"
	"encoding/xml"
	"fmt"
	"infrastructure"
	"interfaces"
	"os"
)

//
//	"path/filepath"
//"encoding/xml"

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
	var filter domain.Benchmark
	filter.Name = "checkconn"
	filter.AddPlatform("rhel", "Red Hat Enterprise Linux", "6.x")
	filter.AddPlatform("rhel", "Red Hat Enterprise Linux", "7.x")

	//var filter domain.TestCase
	//filter.Name = "LDAP Servers"
	//filter.AddTest("ldap01.carcano.local", "ldap01.carcano.local")

	// performs a lookup using a filter
	if values, err := config.Repositories.Lookup(filter, "Benchmarks"); err != nil {
		//if err := config.Repositories.Lookup(filter, "Testcases"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		_, err := xml.Marshal(values[0])
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%s\n", data)
		fmt.Printf("%s\n", values[0].GetId())
	}

	/*
		BenchmarkFile, err := filepath.Abs("benchmark.xml")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		file, err := os.Open(BenchmarkFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer file.Close()
		var bench domain.Benchmark
		err = bench.FromXML(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("\n\n%s\n", bench)
		//data, err := bench.ToXML()
		//fmt.Printf("\n\n%s\n", data)
		if err := config.Repositories.AddDocument(bench, "Benchmarks"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	*/
	/*if err := config.Repositories.DeleteDocumentById("benchmarks", "AV1vvZfZJNdVlB7WYRR-", "Benchmarks"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}*/

	fmt.Printf("\n\nFinito\n")
}
