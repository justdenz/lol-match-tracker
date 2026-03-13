package Models

type Match struct {
	Metadata Metadata `json:"metadata"`
	Info     Info     `json:"info"`
}

type Metadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type Info struct {
	GameCreation       int64         `json:"gameCreation"`
	GameDuration       int64         `json:"gameDuration"`
	GameEndTimestamp   int64         `json:"gameEndTimestamp"`
	GameID             int64         `json:"gameId"`
	GameMode           string        `json:"gameMode"`
	GameName           string        `json:"gameName"`
	GameStartTimestamp int64         `json:"gameStartTimestamp"`
	GameType           string        `json:"gameType"`
	GameVersion        string        `json:"gameVersion"`
	MapID              int           `json:"mapId"`
	Participants       []Participant `json:"participants"`
	QueueID            int           `json:"queueId"`
	Teams              []Team        `json:"teams"`
}

type Participant struct {
	Assists              int    `json:"assists"`
	BaronKills           int    `json:"baronKills"`
	BountyLevel          int    `json:"bountyLevel"`
	ChampExperience      int    `json:"champExperience"`
	ChampLevel           int    `json:"champLevel"`
	ChampionID           int    `json:"championId"`
	ChampionName         string `json:"championName"`
	Deaths               int    `json:"deaths"`
	GoldEarned           int    `json:"goldEarned"`
	GoldSpent            int    `json:"goldSpent"`
	IndividualPosition   string `json:"individualPosition"`
	Kills                int    `json:"kills"`
	Lane                 string `json:"lane"`
	MagicDamageDealt     int    `json:"magicDamageDealt"`
	MagicDamageTaken     int    `json:"magicDamageTaken"`
	NeutralMinionsKilled int    `json:"neutralMinionsKilled"`
	PUUID                string `json:"puuid"`
	PhysicalDamageDealt  int    `json:"physicalDamageDealt"`
	PhysicalDamageTaken  int    `json:"physicalDamageTaken"`
	Role                 string `json:"role"`
	SummonerID           string `json:"summonerId"`
	SummonerName         string `json:"summonerName"`
	TeamID               int    `json:"teamId"`
	TotalDamageDealt     int    `json:"totalDamageDealt"`
	TotalDamageTaken     int    `json:"totalDamageTaken"`
	TotalMinionsKilled   int    `json:"totalMinionsKilled"`
	TrueDamageDealt      int    `json:"trueDamageDealt"`
	TrueDamageTaken      int    `json:"trueDamageTaken"`
	Win                  bool   `json:"win"`
}

type Team struct {
	TeamID     int        `json:"teamId"`
	Win        bool       `json:"win"`
	Objectives Objectives `json:"objectives"`
}

type Objectives struct {
	Baron    Objective `json:"baron"`
	Champion Objective `json:"champion"`
	Dragon   Objective `json:"dragon"`
	Tower    Objective `json:"tower"`
}

type Objective struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}