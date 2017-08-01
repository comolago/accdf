package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)
func main() {
var client *elastic.Client
var err error

client, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false))
// Delete tweet with specified ID
res, err := client.Delete().
    Index("accdf").
    Type("benchmarks").
    Id("AV1v9Jq3NtNJqOG7ngsH").
    Do(context.Background())
if err != nil {
    // Handle error
    panic(err)
}
if res.Found {
    fmt.Print("Document deleted from from index\n")
}
}
