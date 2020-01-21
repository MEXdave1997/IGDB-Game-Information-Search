/*
Game Information search Powered by
the Internet Gaming Database (IGDB)

This package uses a given game name to querey the IGDB and
present important information about the game to the user.

Currently, it presents the following information:
	- Name: The given name of the Game
	- Summary: A Short Description of the Game
	- Genres: A Formatted List of Genres associated with the Game
	- Platforms: A List of Platforms the Game was released on.
	- Release Dates: A List of Released dates for the Game
					 along with Where it was released
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/Henry-Sarabia/igdb"
)

// NOTE: These package level variables are only req.
// for the flags and messages in main()
var key string
var game string

func init() {
	flag.StringVar(&key, "k", "", "Key for use tuse the IGDB API")
	flag.StringVar(&game, "g", "", "Game to be searched")
	flag.Parse()
}

func main() {
	// Provide helpful messages here for when parameters are missing.
	if key == "" {
		fmt.Println("No Key has been provided. Please provide it by doing: \"-k <YOUR API KEY>")
		return
	}
	if game == "" {
		fmt.Println("No Game has been provided. Please provide it by doing: \"-g <\"NAME OF THE GAME\">\"")
		return
	}

	// TODO: Figure out how to implement a way to only grab an exact match, if it exists in the database.
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
		// NOTE: Maybe remove all the "\n" Printf() calls for Println() calls instead?
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

		// NOTE: Grabbing the data for this request was tricky.
		// I needed to iterate through __BOTH__ the platforms __AND__ releases endpoint
		// in order to get the desired results of view release date via the platform endpoint
		release, err := c.ReleaseDates.List(game.ReleaseDates, igdb.SetFields("human", "platform", "region"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n")

		// NOTE: Region Release endpoint still needs more work.
		// TODO: "Prettify" the region string for longer region names (i.e. "NorthAmerica" -> "North America")
		fmt.Println("Release Dates:")
		for i := range platforms {
			for j := range release {
				fmt.Printf("%s(%s): %s\n", platforms[i].Name, strings.TrimLeft(release[j].Region.String(), "Region"), release[j].Human)
			}
		}

	}

	return

}
