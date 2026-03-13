package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Account struct {
    Puuid    string `json:"puuid"`
    GameName string `json:"gameName"`
    TagLine  string `json:"tagLine"`
}

type MatchList struct {
    History []struct {
        MatchID string `json:"matchId"`
    } `json:"history"`
}

func FetchRiot(apiKey, url string, v any) error {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return err
    }

    req.Header.Set("X-Riot-Token", apiKey)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return fmt.Errorf("riot api error %d: %s", resp.StatusCode, string(body))
    }

    return json.NewDecoder(resp.Body).Decode(v)
}


func main() {

	apiKey := os.Getenv("API_KEY")
	accountURL := os.Getenv("ACCOUNT_URL")

    // Account lookup (SEA cluster)
    var acc Account
    if err := FetchRiot(apiKey, accountURL, &acc); err != nil {
        panic(err)
    }
    fmt.Printf("Puuid: %s", acc.Puuid)

    // VALORANT matchlist (AP region)
    matchIdsURL := fmt.Sprintf("https://sea.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?type=ranked&start=%d&count=100", acc.Puuid, 0)
    // matchURL := "https://asia.api.riotgames.com/lol/match/v5/matches/by-puuid/l8Hb-Xt4DswcRxrfyLZGfAuoarzAjTpzWUxokCDQc4DUmdt3ixK0c0tAhJkUk--FWzUnvQCusODXZw/ids?type=normal&start=0&count=20&api_key=RGAPI-d50de9e9-f430-4647-bf2f-70f1ef610416"
	fmt.Println(matchIdsURL)
    var matchIds []string
    if err := FetchRiot(apiKey, matchIdsURL, &matchIds); err != nil {
        panic(err)
    }

	
    getMatchURL := fmt.Sprintf("https://sea.api.riotgames.com/lol/match/v5/matches/%s", matchIds[0])
	var matchData Match
    if err := FetchRiot(apiKey, getMatchURL, &matchData); err != nil {
        panic(err)
    }
	
	fmt.Println(matchData.Metadata.MatchID)
}