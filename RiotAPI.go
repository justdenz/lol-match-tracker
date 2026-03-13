package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	region_url             = "https://%s.api.riotgames.com"
	account_url_by_riot_id = "/riot/account/v1/accounts/by-riot-id/%s/%s"
	matches_url_by_puuid      = "/lol/match/v5/matches/by-puuid/%s/ids?type=%s&start=%d&count=%d"
	matches_by_match_id           = "/lol/match/v5/matches/%s"
)

func GetData(url string, apiKey string, v any) error {
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

func FetchAccount(apiKey, riotID string, riotTag string, v any) error {
	urlParam := fmt.Sprintf(account_url_by_riot_id, riotID, riotTag)
	url := fmt.Sprintf(region_url, "asia") + urlParam

	return GetData(url, apiKey, v)
}

func FetchMatches(apiKey, puuid, matchType string, startRange int, endRange int, v any) error {
	urlParam := fmt.Sprintf(matches_url_by_puuid, puuid, matchType, startRange, endRange)
	url := fmt.Sprintf(region_url, "sea") + urlParam
	return GetData(url, apiKey, v)
}

func FetchMatchByID(apiKey, matchID string, v any) error {
	urlParam := fmt.Sprintf(matches_by_match_id, matchID)
	url := fmt.Sprintf(region_url, "sea") + urlParam
	return GetData(url, apiKey, v)
}