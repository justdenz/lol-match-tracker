package Models

type MatchList struct {
	History []struct {
		MatchID string `json:"matchId"`
	} `json:"history"`
}
