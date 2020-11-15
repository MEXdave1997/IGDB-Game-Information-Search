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

	"net/http"
	// "os"
	"encoding/json"

	"github.com/Henry-Sarabia/igdb/v2"
)

// NOTE: These package-level variables are only req.
// for the flags and messages in main()
var clientID string = "YOUR_CLIENT_ID"
var clientSecret string = "YOUR_CLIENT_SECRET"
var game string

func init() {
	// flag.StringVar(&key, "k", "", "Key for use tuse the IGDB API")
	flag.StringVar(&game, "g", "", "Game to be searched")
	flag.Parse()
}

func main() {

	accessToken, err := getAppToken(clientID, clientSecret)
	if err != nil {
		panic(err)
	}

	// Provide helpful messages here for when parameters are missing.
	if game == "" {
		fmt.Println(`No Game has been provided.
                Please provide it by providing the "-g" flag
                followed by the name of the game you intend to search.`)
		return
	}

	// TODO: Figure out how to implement a way to only grab an exact match, if it exists in the database.
	c := igdb.NewClient(clientID, accessToken.AccessToken, nil)

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

// AccessTokenResponse Struct
// This type stores the AccessToken required to autorize to the API server.
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// This function will validate your clientID and clientSecret on the IGDB Authserver, and store it in an AccessTokenResponse struct
func getAppToken(clientID, clientSecret string) (*AccessTokenResponse, error) {
	appToken := &AccessTokenResponse{}

	resp, err := http.Post(fmt.Sprintf("https://id.twitch.tv/oauth2/token?client_id=%s&client_secret=%s&grant_type=client_credentials", clientID, clientSecret), "text/plain-text", nil)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(appToken)
	if err != nil {
		return nil, err
	}

	return appToken, nil
}
