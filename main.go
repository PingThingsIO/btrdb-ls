package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	btrdb "gopkg.in/btrdb.v4"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Server string
	Prefix string
}

func printRow(longest int, first string, second string) {
	width := fmt.Sprintf("%v", longest+1)
	fmt.Printf("%-"+width+"v%v\n", first, second)
}

func printCollectionsList(server *btrdb.BTrDB, collections []string) {
    if len(collections) == 0 {
        fmt.Println("No collections found!")
        return
    }

	longest := 0
	for _, col := range collections {
		if len(col) > longest {
			longest = len(col)
		}
	}
	printRow(longest, "Collection name", "Stream count")
	for _, collection := range collections {
		streams, err := server.LookupStreams(context.Background(), collection, false, nil, nil)
		streamsCount := ""
		if err != nil {
			streamsCount = fmt.Sprintf("Error getting streams for this collection %v", err)
		} else {
			streamsCount = fmt.Sprintf("%v", len(streams))
		}
		printRow(longest, collection, streamsCount)
	}
}

func printCollectionDetails(server *btrdb.BTrDB, collection string) {
	streams, err := server.LookupStreams(context.Background(), collection, false, nil, nil)
	if err != nil {
		fmt.Printf("Error finding streams for %v: %v\n", collection, err)
		return
	}
	fmt.Println("Collection: " + collection + ":")
	fmt.Println("Streams:")
	for _, stream := range streams {
		fmt.Println(" * UUID: " + stream.UUID().String())

		fmt.Println(" * Tags: ")
		tags, err := stream.Tags(context.Background())
		if err != nil {
			fmt.Printf("     - Error getting tags: %v\n", err)
		} else if len(tags) == 0 {
			fmt.Println("     - None")
		} else {
			for k, v := range tags {
				fmt.Printf("     - %v: %v\n", k, v)
			}
		}

		fmt.Println(" * Annontations: ")
		ann, _, err := stream.Annotations(context.Background())
		if err != nil {
			fmt.Printf("     - Error getting annotations: %v\n", err)
		} else if len(ann) == 0 {
			fmt.Println("     - None")
		} else {
			for k, v := range ann {
				fmt.Printf("     - %v: %v\n", k, v)
			}
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: btrdb-ls <config yaml>\n")
		os.Exit(1)
	}
	configFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Could not read config file %q: %v\n", os.Args[1], err)
		os.Exit(1)
	}

	config := Config{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Printf("Could not parse config file: %v\n", err)
		os.Exit(1)
	}

	server, err := btrdb.Connect(context.Background(), config.Server)
	if err != nil {
		fmt.Printf("Could not connect to server: %v\n", err)
		os.Exit(1)
	}

	collections, err := server.ListCollections(context.Background(), config.Prefix)
	if err != nil {
		fmt.Printf("Error listing collectiosn: %v\n", err)
		os.Exit(1)
	}

	if len(collections) == 1 {
		printCollectionDetails(server, collections[0])
	} else {
		printCollectionsList(server, collections)
	}
}
