package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Henry-Sarabia/igdb"
)

var key string
var game string

func init() {
	flag.StringVar(&key, "k", "", "Key for use tuse the IGDB API")
	flag.StringVar(&game, "g", "", "Game to be searched")
	flag.Parse()
}

func main() {
	if key == "" {
		fmt.Println("No Key has been provided. Please provide it by doing: \"-k <YOUR API KEY>")
		return
	}
	if game == "" {
		fmt.Println("No Game has been provided. Please provide it by doing: \"-g <\"NAME OF THE GAME\">\"")
		return
	}

	c := igdb.NewClient(key, nil)

	// Composing options set to retrieve all fields
	allOpts := igdb.ComposeOptions(
		igdb.SetFields("*"),
	)

	// Retrieve Game information with given name, and log error if it exists.
	search, err := c.Games.Search(
		game,
		allOpts, // top 5 popular results
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print out and format the results.
	fmt.Printf("Information about \"%s\":\n", game)
	for _, game := range search {
		fmt.Printf("\n")
		fmt.Printf("%s\n", game.Name)

		fmt.Printf("\n")
		fmt.Printf("Summary:\n %s\n", game.Summary)

		genres, err := c.Genres.List(game.Genres, igdb.SetFields("name"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n")
		fmt.Println("Genres:")
		for i := range genres {
			fmt.Printf("%s\n", genres[i].Name)
		}

		platforms, err := c.Platforms.List(game.Platforms, igdb.SetFields("name"))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n")
		fmt.Println("Platforms:")
		for i := range platforms {
			fmt.Printf("%s\n", platforms[i].Name)
		}
	}

	return

}
