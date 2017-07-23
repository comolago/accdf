package infrastructure

import (
	"context"
	"domain"

	"fmt"
	"interfaces"
	"reflect"

	"github.com/olivere/elastic"
)

// 	"encoding/xml"

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
func (es *ElasticsearchStore) Open() error {
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
func (es *ElasticsearchStore) Lookup(f interfaces.Filter) ([]domain.Domain, error) {
	/*if s, ok := f.(domain.TestCase); ok {

	}*/
	var query elastic.BoolQuery
	var objTypeIdx int
	if s, ok := f.(domain.Benchmark); ok {
		query = es.lookupBenchmarks(s)
		objTypeIdx = 1
	} else if s, ok := f.(domain.TestCase); ok {

		fmt.Printf("Testcase.Name=%s\n", s.Name)
		query = es.lookupTestcases(s)
		objTypeIdx = 2
	}
	/*if err := es.open(); err != nil {
		return nil, err
	}*/
	res, err := es.connection.Search(es.Index).Type(objectTypes[objTypeIdx]).Query(&query).Do(context.Background())
	if err != nil {
		return nil, interfaces.ErrHandler{5, "func (es *ElasticsearchStore)", "Lookup", ""}
	}
	if objTypeIdx == 1 {
		// Benchmarks
		for _, iT := range res.Each(reflect.TypeOf(&domain.Benchmark{})) {
			es.values = append(es.values, iT.(*domain.Benchmark))
			//fmt.Printf("Elasticsearch Name: %s\n", iT.(*domain.Benchmark).GetName())
		}
	} else if objTypeIdx == 2 {
		// Testcases
		for _, iT := range res.Each(reflect.TypeOf(&domain.TestCase{})) {
			es.values = append(es.values, iT.(*domain.TestCase))
			//fmt.Printf("Elasticsearch Name: %s\n", iT.(*domain.Testcase).GetName())
		}
	}
	if es.values != nil {
		//fmt.Printf("Elasticsearch Name: %s\n", es.values[0].GetName())
		/*data, err := xml.Marshal(es.values[0])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", data)*/
		return es.values, nil
	}
	return nil, nil
}

// creates a benchmark filter
func (es *ElasticsearchStore) lookupTestcases(f interfaces.Filter) elastic.BoolQuery {
	fmt.Printf("QUI\n")
	filter := f.(domain.TestCase)
	query := elastic.NewBoolQuery()

	//fmt.Printf("LEN=%d\n", len(filter.Tests))
	var testsQuery []elastic.BoolQuery
	var nestedTestsQuery []*elastic.NestedQuery
	for i := 0; i < len(filter.Tests); i++ {
		//fmt.Printf("Benchmark.Tests[%d].Name=%s\n", i, filter.Tests[i].Name)
		var boolQuery elastic.BoolQuery
		testsQuery = append(testsQuery, *boolQuery.Filter(
			elastic.NewMatchQuery("Tests>Test.Name", filter.Tests[i].Name),
			elastic.NewMatchQuery("Tests>Test.Label", filter.Tests[i].Label),
		))
		nestedTestsQuery = append(nestedTestsQuery, elastic.NewNestedQuery("Tests>Test", &testsQuery[i]))
		query = query.Filter(nestedTestsQuery[i])
	}

	var mainDocQuery elastic.BoolQuery
	mainDocQuery.Filter(
		elastic.NewMatchQuery("Name", filter.Name),
	)
	query = query.Filter(&mainDocQuery)
	return *query
}

// creates a benchmark filter
func (es *ElasticsearchStore) lookupBenchmarks(f interfaces.Filter) elastic.BoolQuery {
	filter := f.(domain.Benchmark)
	//fmt.Printf("Benchmark.Name=%s\n", filter.Name)
	query := elastic.NewBoolQuery()
	//fmt.Printf("LEN=%d\n", len(filter.Platforms))
	var platformsQuery []elastic.BoolQuery
	var nestedPlatformsQuery []*elastic.NestedQuery
	for i := 0; i < len(filter.Platforms); i++ {
		//fmt.Printf("Benchmark.Platforms[%d].Id=%s\n", i, filter.Platforms[i].Id)
		var boolQuery elastic.BoolQuery
		platformsQuery = append(platformsQuery, *boolQuery.Filter(
			elastic.NewMatchQuery("Platforms>Platform.Id", filter.Platforms[i].Id),
			elastic.NewMatchQuery("Platforms>Platform.Version", filter.Platforms[i].Version),
		))
		nestedPlatformsQuery = append(nestedPlatformsQuery, elastic.NewNestedQuery("Platforms>Platform", &platformsQuery[i]))
		query = query.Filter(nestedPlatformsQuery[i])
	}
	var mainDocQuery elastic.BoolQuery
	mainDocQuery.Filter(
		elastic.NewMatchQuery("Name", filter.Name),
	)
	query = query.Filter(&mainDocQuery)
	return *query
}

// Interface: DbHandler
// add whatsoever Document
func (es *ElasticsearchStore) AddDocument(d interfaces.Document) error {
	var objTypeIdx int
	if _, ok := d.(domain.Benchmark); ok {
		objTypeIdx = 1
	}
	_, err := es.connection.Index().Index(es.Index).Type(objectTypes[objTypeIdx]).Id("").BodyJson(d).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Interface: DbHandler
// delete whatsoever Document
func (es *ElasticsearchStore) DeleteDocumentById(objtype string, id string) error {
	fmt.Printf("Deleting Index=%s objecttype=%s id=%s\n", es.Index, objtype, id)
	_, err := es.connection.Delete().Index(es.Index).Type(objtype).Id(id).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
