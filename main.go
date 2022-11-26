package main

import (
	"fmt"
	"log"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/joho/godotenv"
	"github.com/razonyang/hugo-theme-bootstrap-algolia/internal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
		err = nil
	}

	client := search.NewClient(internal.MustGetEnv("ALGOLIA_APP_ID"), internal.MustGetEnv("ALGOLIA_API_KEY"))

	index := client.InitIndex(internal.MustGetEnv("ALGOLIA_INDEX_NAME"))

	indexer := internal.NewFileIndexer(internal.GetEnvDefault("ALGOLIA_INDEX_FILE", "public/index.json"))
	records, err := indexer.FetchRecords()
	if err != nil {
		log.Fatal(err.Error())
	}

	resSave, err := index.ReplaceAllObjects(records)
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	}
	resSave.Wait()
}
