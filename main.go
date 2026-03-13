package main

import (
	"fmt"
	"lol-match-tracker/Models"
	"os"
)

var match Models.Match
var acc Models.Account
var matchIds []string


func main() {

	apiKey := os.Getenv("API_KEY")

    // Account lookup (SEA cluster)
    riotId := "Munch"
    riotTag := "aivan"
    FetchAccount(apiKey, riotId, riotTag, &acc)
    FetchMatches(apiKey, acc.Puuid, "ranked", 0, 100, &matchIds)
    FetchMatchByID(apiKey, matchIds[0], &match)
    fmt.Printf("Match: %+v\n", match)
}