package api

import (
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

type Match struct {
	Id                       string `json:"MatchID"`
	StartTimestamp           int    `json:"MatchStartTime"`
	CompetetiveMovement      string `json:"CompetitiveMovement"`
	MapId                    string `json:"MapID"`
	TierProgressBeforeUpdate int    `json:"TierProgressBeforeUpdate"`
	TierProgressAfterUpdate  int    `json:"TierProgressAfterUpdate"`
}

type matchHistoryResponse struct {
	Matches *[]Match `json:"Matches"`
	// Matches interface{}
}

func (match *Match) VerboseMapName() string {
	var mapping = map[string]string{
		"ascent":  "ascent",
		"duality": "bind",
		"triad":   "haven",
		"port":    "icebox",
		"bonsai":  "split",
	}
	shortName := strings.ToLower(strings.Split(match.MapId, "/")[4])
	return strings.Title(mapping[shortName])
}

func (match *Match) StartTime() time.Time {
	return time.Unix(0, int64(match.StartTimestamp)*int64(time.Millisecond))
}

func (match *Match) VerboseTime() string {
	return humanize.Time(match.StartTime())
}

func (match *Match) VerboseEloChange() string {
	sign := ""
	if match.EloChange() > 0 {
		sign = "+"
	}
	elo := sign + strconv.FormatInt(int64(match.EloChange()), 10)
	return elo
}

func (match *Match) EloChange() int {
	return match.TierProgressAfterUpdate - match.TierProgressBeforeUpdate
}

func (client *Client) GetMatchHistory() []Match {
	var matches []Match
	var response = matchHistoryResponse{
		Matches: &matches,
	}
	url := "https://pd." + client.region + ".a.pvp.net/mmr/v1/players/" + client.authTokens.UserId + "/competitiveupdates?startIndex=0&endIndex=20"
	client.Request("GET", url, nil, &response)
	return matches
}
