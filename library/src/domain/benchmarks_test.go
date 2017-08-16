package domain

import (
	"fmt"
	"os"
	"path/filepath"
)

func benchmarks_test() {
	fmt.Printf("\nbenchmark_test\n")
	BenchmarkFile, err := filepath.Abs("../../test/xml/benchmarks/000.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputFile, err := os.Open(BenchmarkFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer inputFile.Close()
	bench := new(Benchmark)
	err = bench.FromXML(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data, _ := bench.ToXML()
	outputFile, err := os.Create("../../tmp/xml/benchmarks/000.xml")
	check(err)
	defer outputFile.Close()
	n, err := outputFile.Write(data)
	check(err)
	fmt.Printf("wrote %d bytes\n", n)
	outputFile.Sync()
}
