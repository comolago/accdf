package infrastructure

import (
	"context"
	"domain"
	"encoding/xml"
	"fmt"
	"interfaces"
	"reflect"

	"github.com/olivere/elastic"
)

var objectTypes = [...]string{
	1: "benchmarks",
	2: "testcases",
}

// Store struct definition
type ElasticsearchStore struct {
	URL        string
	Index      string
	connection *elastic.Client
	values     []domain.Domain // interface
}

// Open a connection to es.URL and tests the existence od es.Index
func (es *ElasticsearchStore) open() error {
	if es.URL == "" {
		return interfaces.ErrHandler{2, "func (es *ElasticsearchStore)", "open", ""}
	}
	if es.Index == "" {
		return interfaces.ErrHandler{3, "func (es *ElasticsearchStore)", "open", ""}
	}
	var err error
	es.connection, err = elastic.NewClient(elastic.SetURL(es.URL), elastic.SetSniff(false))
	if err != nil {
		return interfaces.ErrHandler{1, "func (es *ElasticsearchStore)", "open", fmt.Sprintf("%s", err)}
	}
	exists, err := es.connection.IndexExists(es.Index).Do(context.Background())
	if err != nil {
		return interfaces.ErrHandler{1, "func (es *ElasticsearchStore)", "open", fmt.Sprintf("%s", err)}
	}
	if !exists {
		return interfaces.ErrHandler{4, "func (es *ElasticsearchStore)", "open", fmt.Sprintf("Supplied index value: %s", es.Index)}
	}
	return nil
}

// Interface: DbHandler
// Performs a lookup applying f filter
func (es *ElasticsearchStore) Lookup(f interfaces.Filter) error {
	if f, ok := f.(domain.TestCase); ok {
		fmt.Printf("Testcase.Name=%s\n", f.Name)
	}
	var query elastic.BoolQuery
	var objTypeIdx int
	if f, ok := f.(domain.Benchmark); ok {
		query = es.lookupBenchmarks(f)
		objTypeIdx = 1
	}

	if err := es.open(); err != nil {
		return err
	}
	res, err := es.connection.Search(es.Index).Type(objectTypes[objTypeIdx]).Query(&query).Do(context.Background())
	if err != nil {
		return interfaces.ErrHandler{5, "func (es *ElasticsearchStore)", "Lookup", ""}
	}
	if objTypeIdx == 1 {
		// Benchmarks
		for _, iT := range res.Each(reflect.TypeOf(&domain.Benchmark{})) {
			es.values = append(es.values, iT.(*domain.Benchmark))
			fmt.Printf("Elasticsearch Name: %s\n", iT.(*domain.Benchmark).GetName())
		}
	}
	if es.values != nil {
		fmt.Printf("Elasticsearch Name: %s\n", es.values[0].GetName())
		data, err := xml.Marshal(es.values[0])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", data)
	}
	return nil
}

// creates a benchmark filter
func (es *ElasticsearchStore) lookupBenchmarks(f interfaces.Filter) elastic.BoolQuery {
	filter := f.(domain.Benchmark)
	fmt.Printf("Benchmark.Name=%s\n", filter.Name)
	query := elastic.NewBoolQuery()
	// if (len(filter.Platforms) > 0 ) {
	fmt.Printf("LEN=%d\n", len(filter.Platforms))

	//platformsQuery *elastic.BoolQuery
	platformsQuery := elastic.NewBoolQuery()
	platformsQuery2 := elastic.NewBoolQuery()
	//for i := 0; i < 2; i++ {
	//fmt.Printf("Benchmark.Platforms[%d].Id=%s\n", i, filter.Platforms[i].Id)
	platformsQuery.Filter(

		elastic.NewMatchQuery("Platforms>Platform.Id", filter.Platforms[0].Id),
		elastic.NewMatchQuery("Platforms>Platform.Version", filter.Platforms[0].Version),
		//elastic.NewMatchQuery("Name", "checkconn"),

	)
	platformsQuery2.Filter(
		elastic.NewMatchQuery("Platforms>Platform.Id", filter.Platforms[1].Id),
		elastic.NewMatchQuery("Platforms>Platform.Version", filter.Platforms[1].Version),
	)
	//}
	nestedPlatformsQuery := elastic.NewNestedQuery("Platforms>Platform", platformsQuery)
	nestedPlatformsQuery2 := elastic.NewNestedQuery("Platforms>Platform", platformsQuery2)

	//}

	query = query.Filter(nestedPlatformsQuery, nestedPlatformsQuery2)

	//filter.Platforms[0].Id="rhel"
	//filter.Platforms[0].Version="7.x"

	return *query
}

// Modified
