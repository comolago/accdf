all: tests

tests:  test-add test-lookup

test-add:
	go test -v -run=TestAddBenchmark

test-lookup:
	go test -v -run=TestLookup

dump: dump-benchmarks

dump-benchmarks:
	curl  -XGET 'http://localhost:9200/accdf/benchmarks/_search?pretty'

index-delete:
	curl -XDELETE 'http://localhost:9200/accdf'

index-create:
	curl  -XPUT 'http://localhost:9200/accdf' -d @index.json
