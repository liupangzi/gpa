package cmd

import (
	"context"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
)

func main() {
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(es.Info())

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(``),
		es.Search.WithBody(strings.NewReader(`
{
	"query":
		{
			"match":
				{
					"title": "test"
} }}
`)),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
}
