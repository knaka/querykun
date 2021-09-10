package main

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve/v2"
)

func main() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("/tmp/example.bleve", mapping)
	if err != nil {
		fmt.Println("cp2:", err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "text",
	}

	// index some data
	if err := index.Index("id", data); err != nil {
		log.Fatal(err.Error())
	}

	// search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println("cp0:", err)
		return
	}
	fmt.Println("cp1:", searchResults)
}
